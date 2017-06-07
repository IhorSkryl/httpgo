// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"
	"github.com/ruslanBik4/httpgo/models/db/schema"
	"github.com/ruslanBik4/httpgo/views"
	"github.com/ruslanBik4/httpgo/views/templates/json"
	_ "strings"
	"github.com/ruslanBik4/httpgo/models/services"
	viewsSystem "github.com/ruslanBik4/httpgo/views/templates/system"
	"github.com/ruslanBik4/httpgo/models/db/qb"
	"github.com/ruslanBik4/httpgo/models/db"
	"strings"
	"fmt"
	"github.com/ruslanBik4/httpgo/models/logs"
)
// prepare JSON with fields type from structere DB and + 1 row with data if issue parameter "id"
func HandleFieldsJSON(w http.ResponseWriter, r *http.Request) {

	tableName := r.FormValue("table")

	if tableName == "" {
		views.RenderBadRequest(w)
		return
	}

	defer func() {
		err1 := recover()
		switch err := err1.(type) {
		case schema.ErrNotFoundTable:
			views.RenderInternalError(w, err)
		case nil:
		default:
			panic(err)
		}
	}()

	fields := schema.GetFieldsTable(tableName)
	for idx, field := range fields.Rows {

		if field.SETID || field.NODEID || field.IdForeign {

			sqlCommand := field.SQLforFORMList
			for _, enumVal := range field.EnumValues {
				if i := strings.Index(enumVal, ":"); i > 0 {
					// мы добавим условие созначением пол текущей записи, если это поле найдено и в нем установлено значение
					if paramValue := r.FormValue(enumVal[i+1:]); (paramValue > "")  {
						enumVal = enumVal[:i] + fmt.Sprintf("%s", paramValue)
						if strings.Contains(sqlCommand, "WHERE") {
							sqlCommand += " OR " + enumVal
						} else {
							sqlCommand += " WHERE " + enumVal

						}
					} else {
						continue
					}
				}

			}
			//TODO: add where condition
			logs.DebugLog(sqlCommand)
			rows, err := db.DoSelect(sqlCommand)
			if err != nil {
				logs.ErrorLog(err, field.SQLforFORMList)
			} else {

				defer rows.Close()
				for rows.Next() {
					var key int
					var title string
					if err := rows.Scan(&key, &title); err != nil {
						logs.ErrorLog(err)
					}

					fields.Rows[idx].SelectValues[key] = title
				}
			}

		}
	}

	addJSON := make(map[string]string, 0)
	if id := r.FormValue("id"); id > "" {
		// получаем данные для суррогатных полей
		qBuilder := qb.Create("id=?", "", "")
		qBuilder.AddTable("a", tableName)
		qBuilder.AddArg(id)
		arrJSON, err := qBuilder.SelectToMultidimension()
		if err != nil {
			views.RenderInternalError(w, err)
			return
		}

		addJSON["data"] = json.WriteSliceJSON(arrJSON)
	}

	views.RenderJSONAnyForm(w, fields, new (json.FormStructure), addJSON)
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
		qBuilder := qb.Create("id=?", "", "")
		qBuilder.AddTable("a", tableName)
		qBuilder.AddArg(id)
		arrJSON, err := qBuilder.SelectToMultidimension()
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
func HandleAllRowsJSON(w http.ResponseWriter, r *http.Request) {
	tableName := r.FormValue("table")

	if (tableName > "") {
		qBuilder := qb.Create("", "", "")
		qBuilder.AddTable("a", tableName)
		arrJSON, err := qBuilder.SelectToMultidimension()
		if err != nil {
			views.RenderInternalError(w, err)
			return
		}
		if len(arrJSON) > 0 {
			views.RenderAnyJSON(w, arrJSON[0])
			return
		}
	} else {
		views.RenderBadRequest(w)
	}
}
