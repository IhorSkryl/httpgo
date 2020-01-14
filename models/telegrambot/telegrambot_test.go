// Copyright 2018 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package telegrambot

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/valyala/fasthttp"

	"github.com/ruslanBik4/httpgo/logs"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestISWrite(t *testing.T) {
	b := &TelegramBot{}

	assert.Implements(t, (*io.Writer)(nil), b)
}

func TestSome(t *testing.T) {
	as := assert.New(t)
	as.Containsf("true false dunno", "true", "dunn")
}

func TestNewTelegramBotFromEnv(t *testing.T) {
	os.Setenv("TBTOKEN", "bottoken")
	os.Setenv("TBCHATID", "chatid")
	tb, err := NewTelegramBotFromEnv()
	as := assert.New(t)

	as.Nil(err, "%v", err)
	as.Equal("bottoken", tb.Token, "Token from env wrong")
	as.Equal("chatid", tb.ChatID, "ChatID from env wrong")

	err = tb.SendMessage("some mess", false)
	as.Nil(err, "%v", err)

	_, err = tb.Write([]byte("some mess"))
	as.Nil(err, "%v", err)

	tb = &TelegramBot{}

	_, err = tb.Write([]byte("some mess"))
	if as.NotNil(err) {
		as.EqualError(err, "TelegramBot.Token empty")
	}
}

//server from fasthttp example
func FastHTTPServer() {
	ln, err := net.Listen("tcp", ":21476")
	if err != nil {
		log.Fatalf("error in net.Listen: %s", err)
	}
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		fmt.Println(ctx, "Requested path is")
	}
	if err := fasthttp.Serve(ln, requestHandler); err != nil {
		log.Fatalf("error in Serve: %s", err)
	}
}

func TestErrorLogTelegramWrite(t *testing.T) {
	os.Setenv("TBTOKEN", "bottoken")
	os.Setenv("TBCHATID", "chatid")
	tb, err := NewTelegramBotFromEnv()
	as := assert.New(t)
	fmt.Println("tb")

	as.Nil(err, "%v", err)
	as.Equal("bottoken", tb.Token, "Token from env wrong")
	as.Equal("chatid", tb.ChatID, "ChatID from env wrong")

	tb.RequestURL = "http://localhost:21476/"

	newError := errors.New("NewERROR")
	newErrorWraped := errors.Wrap(newError, "Wraped")

	parametrsCheck := []string{"chat_id", "text"}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	// ===== Simple net.http/ListenAndServe server with specific handler to read out telgrambot request =====
	http.HandleFunc("/bottoken/sendMessage",
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("HandleFunc r.URL", r.URL)

			for i, paramName := range parametrsCheck {
				params, ok := r.URL.Query()[paramName]

				if !ok || len(params[0]) < 1 {
					log.Println("Url Param", paramName, "is missing")
					t.Fail()
					wg.Done()
					return
				}

				if i == 0 {
					as.Equal(string(params[0]), tb.ChatID, "ChatID in request is wrong")
				} else {
					if strings.Contains(string(params[0]), strings.Replace(newError.Error(), " ", "%20", -1)) ||
						strings.Contains(string(params[0]), strings.Replace(newErrorWraped.Error(), " ", "%20", -1)) {
						wg.Done()
						return
					}
				}
			}
		})
	go http.ListenAndServe(":21476", nil)
	// =========================================

	//// another server version, but hadler hasn't written for the errors and waitgroup
	//go FastHTTPServer()

	logs.SetWriters(tb, logs.FgErr)

	logs.ErrorLog(newError)
	logs.ErrorLog(newErrorWraped)

	wg.Wait()

}
