// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httpGo

import (
	"io/ioutil"

	"github.com/valyala/fasthttp"
	"gopkg.in/yaml.v2"

	"github.com/ruslanBik4/httpgo/logs"
)

// CfgHttp has some options for Acceptor work
type CfgHttp struct {
	// list tokens to check requests
	KillSignal int              `yaml:"KillSignal"`
	Server     *fasthttp.Server `yaml:"Server"`
}

// NewCfgHttp create CfgHttp from config file
func NewCfgHttp(filename string) (cfgGlobal *CfgHttp, err error) {

	var buf []byte

	buf, err = ioutil.ReadFile(filename)
	if err != nil {
		logs.ErrorLog(err)
		return nil, err
	}

	err = yaml.Unmarshal(buf, &cfgGlobal)

	return
}