// Copyright 2020 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"fmt"
	"reflect"
	"regexp"
	"runtime"

	"github.com/valyala/fasthttp"
)

type AuthBearer struct {
	tokens Tokens
}

func NewAuthBearer(tokens Tokens) *AuthBearer {
	if tokens == nil {
		tokens = &mapTokens{
			expiresIn: tokenExpires,
			tokens:    make(map[string]*mapToken, 0),
		}
	}

	return &AuthBearer{tokens}
}

func (a *AuthBearer) NewToken(userData TokenData) (string, error) {
	return a.tokens.NewToken(userData)
}

func (a *AuthBearer) GetToken(ctx *fasthttp.RequestCtx) TokenData {
	bearer := a.getBearer(ctx)
	if bearer == "" {
		return nil
	}

	return a.tokens.GetToken(bearer)
}

var regBearer = regexp.MustCompile(`Bearer\s+(\S+)`)

func (a *AuthBearer) getBearer(ctx *fasthttp.RequestCtx) string {
	b := regBearer.FindSubmatch(ctx.Request.Header.Peek("Authorization"))
	if len(b) == 0 {
		return ""
	}

	return string(b[1])
}

func (a *AuthBearer) String() string {

	return `implement auth for Bearer standart: 
	 user: ` + getStringOfFnc(reflect.ValueOf(a.Auth).Pointer()) + `
	 admin: ` + getStringOfFnc(reflect.ValueOf(a.AdminAuth).Pointer())
}

func (a *AuthBearer) Auth(ctx *fasthttp.RequestCtx) bool {

	token := a.GetToken(ctx)
	if token == nil {
		return false
	}

	ctx.SetUserValue(UserValueToken, token)

	return true
}

func (a *AuthBearer) AdminAuth(ctx *fasthttp.RequestCtx) bool {

	return a.Auth(ctx) && (ctx.UserValue(UserValueToken).(TokenData).IsAdmin())
}

func getStringOfFnc(pc uintptr) string {
	fnc := runtime.FuncForPC(pc)
	fName, line := fnc.FileLine(0)

	return fmt.Sprintf("%s:%d %s()", fName, line, fnc.Name())
}
