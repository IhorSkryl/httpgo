// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"
	_ "github.com/ruslanBik4/httpgo/views"
)

func HandleVideos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
