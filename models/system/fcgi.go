// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// prepare & run request into php-fpm

package system

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/textproto"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"bitbucket.org/PinIdea/fcgi_client"
	"github.com/pkg/errors"
	. "github.com/valyala/fasthttp"

	"github.com/ruslanBik4/httpgo/apis"
	"github.com/ruslanBik4/httpgo/logs"
)

const internalRewriteFieldName = "travel"

var (
	headerNameReplacer = strings.NewReplacer(" ", "_", "-", "_")
)

// FCGI structure for request in PHP-FPM
type FCGI struct {
	Sock string
	Env  func(ctx *RequestCtx) map[string]string
}

func (c *FCGI) defaultEnv(ctx *RequestCtx) map[string]string {
	return map[string]string{
		"REQUEST_METHOD":  string(ctx.Method()),
		"SCRIPT_FILENAME": string(ctx.URI().Path()),
		"SCRIPT_NAME":     string(ctx.URI().Path()),
		"QUERY_STRING":    string(ctx.URI().QueryString()),
	}
}

// Do run request
func (c *FCGI) Do(ctx *RequestCtx) error {
	const typeSckt = "unix" // or "unixgram" or "unixpacket"

	fcgi, err := fcgiclient.Dial(typeSckt, c.Sock)
	if err != nil {
		return err
	}
	env := c.Env
	if env == nil {
		env = c.defaultEnv
	}
	params := env(ctx)

	var req io.Reader

	switch string(ctx.Method()) {
	case "GET":
		req = nil
	case "POST":
		b := ctx.PostBody()
		req = bytes.NewReader(b)
	}

	r, err := fcgi.Do(params, req)
	if err != nil {
		return errors.Wrap(err, "")
	}
	rb := bufio.NewReader(r)
	tp := textproto.NewReader(rb)

	// Parse the response headers.
	mimeHeader, err := tp.ReadMIMEHeader()
	if err != nil {
		logs.ErrorLog(err, mimeHeader)
		err = apis.WriteJSON(ctx, params)
		// ctx.Response.AppendBodyString("<html><body>")
		// for key, val := range params {
		// 	str := fmt.Sprintf("<p>%s = %s </p>", key, val)
		// 	ctx.Response.AppendBodyString(str)
		// }
		// ctx.Response.AppendBodyString("</body></html>")

		return err
	}

	for key, val := range mimeHeader {
		switch key {
		case "Content-Length":
			contentLength, _ := strconv.Atoi(val[0])
			ctx.Response.Header.SetContentLength(contentLength)
		case "Content-Type":
			ctx.Response.Header.SetContentType(val[0])
		case "Status":
			c, _ := strconv.Atoi(val[0])
			ctx.Response.SetStatusCode(c)
		default:
			ctx.Response.Header.Add(key, strings.Join(val, ";"))
		}
	}

	ctx.Response.SetBodyStream(ioutil.NopCloser(rb), -3)

	return err
}

// ServeHTTP get request response & render to output
func (c *FCGI) ServeHTTP(ctx *RequestCtx) (interface{}, error) {
	err := c.Do(ctx)
	if err != nil {
		return nil, err
	}

	// status, isStatus := resp.Header[]
	// location, isURL := ctx.Response.Header["Location"]
	// if isStatus && (status[0] == "302 Found") && isURL {
	// 	ctx.Redirect(location[0], StatusTemporaryRedirect)
	// 	return nil, nil
	// }

	return nil, nil
}

// NewFPM create new FCGI
func NewFPM(sock string) *FCGI {
	return &FCGI{Sock: sock}
}

var pathinfoRe = regexp.MustCompile(`^(.+\.php)(/?.+)$`)

// NewPHP create new FCGI for PHP scripts
func NewPHP(root string, priScript, port, sock string) *FCGI {
	return &FCGI{
		Sock: sock,
		Env: func(ctx *RequestCtx) map[string]string {

			ip, port := ctx.RemoteAddr().String(), ""
			if idx := strings.LastIndex(ip, ":"); idx > -1 {
				port = ip[idx+1:]
				ip = ip[:idx]
			}
			pathInfo, docURI, scriptName := "", string(ctx.RequestURI()), string(ctx.Request.URI().Path())

			if matches := pathinfoRe.FindStringSubmatch(scriptName); len(matches) > 0 {
				scriptName, pathInfo = matches[1], matches[2]
			}

			if idx := strings.Index(docURI, pathInfo); idx > -1 {
				docURI = docURI[len(pathInfo):]
			}
			if strings.HasSuffix(scriptName, "/") {
				scriptName = path.Join(scriptName, priScript)
			}
			// Some variables are unused but cleared explicitly to prevent
			// the parent environment from interfering.
			env := map[string]string{

				// Variables defined in CGI 1.1 spec
				"AUTH_TYPE":         "", // Not used
				"CONTENT_LENGTH":    strconv.Itoa(ctx.Request.Header.ContentLength()),
				"CONTENT_TYPE":      string(ctx.Request.Header.ContentType()),
				"GATEWAY_INTERFACE": "CGI/1.1",
				"PATH_INFO":         pathInfo,
				"QUERY_STRING":      string(ctx.URI().QueryString()),
				"REMOTE_ADDR":       ip,
				"REMOTE_HOST":       ip, // For speed, remote host lookups disabled
				"REMOTE_PORT":       port,
				"REMOTE_IDENT":      "", // Not used
				"REMOTE_USER":       "", // Not used
				"REQUEST_URI":       ctx.URI().String(),
				"REQUEST_METHOD":    string(ctx.Method()),
				"REDIRECT_STATUS":   "200",
				"REQUEST_SCHEME":    string(ctx.Request.URI().Scheme()),
				"SERVER_NAME":       string(ctx.Host()),
				"SERVER_PORT":       "80",
				"SERVER_PROTOCOL":   "HTTP/1.1",
				"SERVER_SOFTWARE":   "httpGo 0.01",

				// Other variables
				"DOCUMENT_ROOT":   root,
				"DOCUMENT_INDEX":  priScript,
				"DOCUMENT_URI":    docURI,
				"HTTP_HOST":       string(ctx.Host()), // added here, since not always part of headers
				"SCRIPT_FILENAME": path.Join(root, scriptName),
				"SCRIPT_NAME":     scriptName,
			}
			// compliance with the CGI specification that PATH_TRANSLATED
			// should only exist if PATH_INFO is defined.
			// Info: https://www.ietf.org/rfc/rfc3875 Page 14
			if env["PATH_INFO"] != "" {
				env["PATH_TRANSLATED"] = path.Join(root, pathInfo) // Info: http://www.oreilly.com/openbook/cgi/ch02_04.html
			}

			// Some web apps rely on knowing HTTPS or not
			if ctx.IsTLS() {
				env["HTTPS"] = "on"
			}

			// Add all HTTP headers (except Caddy-Rewrite-Original-URI ) to env variables
			ctx.Request.Header.VisitAll(func(key, value []byte) {
				field, val := string(key), string(value)
				// /observe
				if strings.ToLower(field) == strings.ToLower(internalRewriteFieldName) {
					return
				}
				header := strings.ToUpper(field)
				if header == "CONTENT_TYPE" || header == "CONTENT_LENGTH" {
					return
				}
				header = headerNameReplacer.Replace(header)
				env["HTTP_"+header] = val
			})

			return env
		},
	}
}

// WriteError не уверен, что это должно быть здесь - должен быть какой общий механизм для выдачи такого
func WriteError(ctx *RequestCtx, err error) bool {
	if err == nil {
		return false
	}

	if os.IsNotExist(err) {
		ctx.SetStatusCode(http.StatusNotFound)
		return true
	}

	ctx.SetStatusCode(http.StatusInternalServerError)
	logs.ErrorLog(errors.Wrap(err, "SetStatusCode"))

	return true
}
