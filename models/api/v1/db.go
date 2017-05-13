// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"
	"github.com/ruslanBik4/httpgo/models/db"
	"github.com/ruslanBik4/httpgo/views"
	"github.com/ruslanBik4/httpgo/models/admin"
	"github.com/ruslanBik4/httpgo/views/templates/json"
	"strings"
	"github.com/ruslanBik4/httpgo/models/services"
	"github.com/ruslanBik4/httpgo/models/db/schema"
	viewsSystem "github.com/ruslanBik4/httpgo/views/templates/system"
)

func HandleFieldsJSON(w http.ResponseWriter, r *http.Request) {

	tableName := r.FormValue("table")

	if tableName == "" {
		views.RenderBadRequest(w)
		return
	}

	fields := admin.GetFields(tableName)

	addJSON := make(map[string]string, 0)
	if r.FormValue("id") > "" {
		id := r.FormValue("id")
		arrJSON, err := db.SelectToMultidimension("select * from " + tableName + " where id=?", id)
		if err != nil {
			panic(err)
		}

		// значение приходит в виде строки. Для агрегатных полей нужно формировать вложеность
		for sKey, sValue := range arrJSON {

			for key, value := range sValue {

				if strings.HasPrefix(key, "setid_") || strings.HasPrefix(key, "nodeid_") || strings.HasPrefix(key, "tableid_") {

					switch vv := value.(type) {
					case []map[string]interface{} :
						arrJSON[sKey][key] = convertToMultiDimension(vv)
					}
				}
			}
		}

		addJSON["data"] = json.WriteSliceJSON(arrJSON)
	}

	views.RenderJSONAnyForm(w, &fields, new (json.FormStructure), addJSON)
}

func convertToMultiDimension(array [] map[string]interface{}) json.MapMultiDimension {

	var mapToDem = json.MapMultiDimension{}
	for _,val := range array {
		mapToDem = append(mapToDem, val)
	}
	return mapToDem
}
func HandleSchema(w http.ResponseWriter, r *http.Request) {
	tableName := r.FormValue("table")
	if table, err := services.Get("schema", tableName); err != nil {
		views.RenderInternalError(w, err)
	} else {
		w.Write([]byte(viewsSystem.ShowSchema(table.(*schema.FieldsTable) )))
	}

}
func HandleRowJSON(w http.ResponseWriter, r *http.Request) {
	tableName := r.FormValue("table")
	id := r.FormValue("id")

	if (tableName > "") && (id > "") {
		arrJSON, err := db.SelectToMultidimension("select * from " + tableName + " where id=?", id)
		if err != nil {
			views.RenderInternalError(w, err)
		}
		if len(arrJSON) > 0 {
			views.RenderAnyJSON(w, arrJSON[0])
			return
		}
	} else {
		views.RenderBadRequest(w)
	}
}
