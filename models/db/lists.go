// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package db

import (
	"github.com/ruslanBik4/httpgo/models/server"
	"github.com/ruslanBik4/httpgo/models/logs"
)
//InitLists() - инициализация получения информации по справочникам
//@aurhor Ruslan Bikchentaev
//@version 1.10 2017-05-24 Sergey Litvinov
func InitLists() {
	go func() {
		var tables RecordsTables
		where :=`TABLE_SCHEMA="` + server.GetServerConfig().DBName() + `" AND (RIGHT(table_name, 5) =  "_list")`;
		err := tables.GetSelectTablesProp(where )

		if err != nil {
			logs.ErrorLog(err, where)
		}

	}()
}
