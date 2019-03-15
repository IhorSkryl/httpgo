// Copyright 2018 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"net"
	"os"
	"path"
	"path/filepath"
	"strings"

	httpgo "github.com/dvbogdan/analytics/httpsgo"
	"github.com/valyala/fasthttp"

	"github.com/ruslanBik4/httpgo/apis"
	"github.com/ruslanBik4/httpgo/models/logs"
)

var (
	routes = apis.APIRoutes{
		"/": &apis.APIRoute{
			Desc: "default endpoint",
			Fnc:  HandleIndex,
		},
	}

	fPort    = flag.String("port", ":8080", "host address to listen on")
	fSystem  = flag.String("path", "./", "path to system files")
	fCfgPath = flag.String("config path", "cfg", "path to cfg files")
	fWeb     = flag.String("web", "front", "path to web files")

	httpServer *httpgo.Httpgo
)

func init() {
	flag.Parse()

	listener, err := net.Listen("tcp", *fPort)
	if err != nil {
		// port is occupied - work serve unpossable
		logs.Fatal(err)
	}

	ctxApis := apis.NewCtxApis(0)
	ctxApis.AddValue("version", "0.2")
	apis := apis.NewApis(ctxApis, routes, func(ctx *fasthttp.RequestCtx) bool {
		return false
	})
	if err != nil {
		// not work without correct config
		logs.Fatal(err)
	}

	cfg, err := httpgo.NewCfgHttp(path.Join(*fSystem, *fCfgPath, "httpgo.yml"))
	if err != nil {
		// not work without correct config
		logs.Fatal(err)
	}
	httpServer = httpgo.NewHttpgo(cfg, listener, apis)

}

func main() {
	err := httpServer.StartServer(false, "", "")

	if err != nil {
		logs.ErrorLog(err)
	} else {
		logs.StatusLog("Server https correct shutdown")
	}

}

func HandleIndex(ctx *fasthttp.RequestCtx) (interface{}, error) {
	filename := strings.TrimLeft(string(ctx.Request.URI().Path()), "/")
	if filename == "" {
		fasthttp.ServeFile(ctx, filepath.Join(*fWeb, "index.html"))
		return nil, nil
	}

	ext := filepath.Ext(filename)
	if ext == "css" {

	}
	fullName := filepath.Join(*fWeb, filename)
	_, err := os.Stat(fullName)
	if os.IsNotExist(err) {
		fasthttp.ServeFile(ctx, filepath.Join(*fWeb, "index.html"))
	} else {
		fasthttp.ServeFile(ctx, fullName)
	}

	return nil, nil
}