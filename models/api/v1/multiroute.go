// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"github.com/ruslanBik4/httpgo/views"
	"io"
	"net/http"
)

// HandleMultiRouteJSON prepare JSON with fields type from structere DB and + 1 row with data if issue parameter "id"
// @/api/multiroute/?route[]={routes list}
func HandleMultiRouteJSON(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(_2K)

	routes, ok := r.Form["route"]

	if !ok {
		views.RenderNotParamsInPOST(w, "route")
		return
	}

	r.Form.Del("route")
	url := "http://" + r.Host

	comma := `{"`
	for _, val := range routes {
		resp, err := http.PostForm(url+val, r.Form)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(comma + val + `":`))
		io.Copy(w, resp.Body)
		comma = `,"`
	}
	w.Write([]byte("}"))

}
