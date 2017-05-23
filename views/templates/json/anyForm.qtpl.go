// This file is automatically generated by qtc from "anyForm.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line views/templates/json/anyForm.qtpl:1
package json

//line views/templates/json/anyForm.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

//line views/templates/json/anyForm.qtpl:5
import (
	"github.com/ruslanBik4/httpgo/models/db/schema"
	"strings"
)

// Формируем JSON, который затем будет использован в форме на клиенте
// dataJson и содержимое AddJson вставляются КАК ЕСТЬ, ПОТОМУ ЧТО БЫЛИ ОБРАБОТАНЫ РАНЕЕ!!!

//line views/templates/json/anyForm.qtpl:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/json/anyForm.qtpl:13
func StreamCheckFields(qw422016 *qt422016.Writer, ns schema.FieldsTable) {
	//line views/templates/json/anyForm.qtpl:15
	for _, field := range ns.Rows {
		//line views/templates/json/anyForm.qtpl:19
		titleFull, titleLabel, placeholder, pattern, dataJson := field.GetColumnTitles()
		typeInput := field.TypeInput()

		//line views/templates/json/anyForm.qtpl:21
		qw422016.N().S(`"`)
		//line views/templates/json/anyForm.qtpl:22
		qw422016.E().S(field.COLUMN_NAME)
		//line views/templates/json/anyForm.qtpl:22
		qw422016.N().S(`": {`)
		//line views/templates/json/anyForm.qtpl:23
		if field.Value > "" {
			//line views/templates/json/anyForm.qtpl:23
			qw422016.N().S(`"value": "`)
			//line views/templates/json/anyForm.qtpl:23
			qw422016.N().J(field.Value)
			//line views/templates/json/anyForm.qtpl:23
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:23
		}
		//line views/templates/json/anyForm.qtpl:25
		if (field.COLUMN_NAME == "id") || field.IsHidden {
			//line views/templates/json/anyForm.qtpl:25
			qw422016.N().S(`"type": "hidden" },`)
			//line views/templates/json/anyForm.qtpl:27
			continue
			//line views/templates/json/anyForm.qtpl:28
		}
		//line views/templates/json/anyForm.qtpl:30
		if field.IS_NULLABLE == "NO" {
			//line views/templates/json/anyForm.qtpl:30
			qw422016.N().S(`"required": true,`)
			//line views/templates/json/anyForm.qtpl:30
		}
		//line views/templates/json/anyForm.qtpl:31
		if titleFull > "" {
			//line views/templates/json/anyForm.qtpl:31
			qw422016.N().S(`"title": "`)
			//line views/templates/json/anyForm.qtpl:31
			qw422016.E().J(titleFull)
			//line views/templates/json/anyForm.qtpl:31
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:31
		}
		//line views/templates/json/anyForm.qtpl:32
		if titleLabel > "" {
			//line views/templates/json/anyForm.qtpl:32
			qw422016.N().S(`"label": "`)
			//line views/templates/json/anyForm.qtpl:32
			qw422016.E().J(titleLabel)
			//line views/templates/json/anyForm.qtpl:32
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:32
		}
		//line views/templates/json/anyForm.qtpl:33
		if placeholder > "" {
			//line views/templates/json/anyForm.qtpl:33
			qw422016.N().S(`"placeholder": "`)
			//line views/templates/json/anyForm.qtpl:33
			qw422016.E().J(placeholder)
			//line views/templates/json/anyForm.qtpl:33
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:33
		}
		//line views/templates/json/anyForm.qtpl:34
		if pattern > "" {
			//line views/templates/json/anyForm.qtpl:34
			qw422016.N().S(`"pattern": "`)
			//line views/templates/json/anyForm.qtpl:34
			qw422016.N().J(pattern)
			//line views/templates/json/anyForm.qtpl:34
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:34
		}
		//line views/templates/json/anyForm.qtpl:35
		if dataJson > "" {
			//line views/templates/json/anyForm.qtpl:35
			qw422016.N().S(dataJson)
			//line views/templates/json/anyForm.qtpl:35
			qw422016.N().S(`,`)
			//line views/templates/json/anyForm.qtpl:35
		}
		//line views/templates/json/anyForm.qtpl:36
		if field.CSSClass > "" {
			//line views/templates/json/anyForm.qtpl:36
			qw422016.N().S(`"CSSClass": "`)
			//line views/templates/json/anyForm.qtpl:36
			qw422016.E().J(field.CSSClass)
			//line views/templates/json/anyForm.qtpl:36
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:36
		}
		//line views/templates/json/anyForm.qtpl:37
		if field.CSSStyle > "" {
			//line views/templates/json/anyForm.qtpl:37
			qw422016.N().S(`"CSSStyle": "`)
			//line views/templates/json/anyForm.qtpl:37
			qw422016.E().J(field.CSSStyle)
			//line views/templates/json/anyForm.qtpl:37
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:37
		}
		//line views/templates/json/anyForm.qtpl:38
		if field.Figure > "" {
			//line views/templates/json/anyForm.qtpl:38
			qw422016.N().S(`"Figure": "`)
			//line views/templates/json/anyForm.qtpl:38
			qw422016.E().J(field.Figure)
			//line views/templates/json/anyForm.qtpl:38
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:38
		}
		//line views/templates/json/anyForm.qtpl:39
		if field.CHARACTER_MAXIMUM_LENGTH > 0 {
			//line views/templates/json/anyForm.qtpl:39
			qw422016.N().S(`"maxLength":`)
			//line views/templates/json/anyForm.qtpl:39
			qw422016.N().D(field.CHARACTER_MAXIMUM_LENGTH)
			//line views/templates/json/anyForm.qtpl:39
			qw422016.N().S(`,`)
			//line views/templates/json/anyForm.qtpl:39
		}
		//line views/templates/json/anyForm.qtpl:40
		if strings.Contains(field.COLUMN_TYPE, "unsigned") {
			//line views/templates/json/anyForm.qtpl:40
			qw422016.N().S(`"min":0,`)
			//line views/templates/json/anyForm.qtpl:40
		}
		//line views/templates/json/anyForm.qtpl:41
		if field.MinDate > "" {
			//line views/templates/json/anyForm.qtpl:41
			qw422016.N().S(`"minDate": "`)
			//line views/templates/json/anyForm.qtpl:41
			qw422016.N().J(field.MinDate)
			//line views/templates/json/anyForm.qtpl:41
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:41
		}
		//line views/templates/json/anyForm.qtpl:42
		if field.MaxDate > "" {
			//line views/templates/json/anyForm.qtpl:42
			qw422016.N().S(`"maxDate": "`)
			//line views/templates/json/anyForm.qtpl:42
			qw422016.N().J(field.MaxDate)
			//line views/templates/json/anyForm.qtpl:42
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:42
		}
		//line views/templates/json/anyForm.qtpl:43
		if len(field.Events) > 0 {
			//line views/templates/json/anyForm.qtpl:43
			qw422016.N().S(`"Events": {`)
			//line views/templates/json/anyForm.qtpl:45
			for name, funcName := range field.Events {
				//line views/templates/json/anyForm.qtpl:45
				qw422016.N().S(`"`)
				//line views/templates/json/anyForm.qtpl:46
				qw422016.E().S(name)
				//line views/templates/json/anyForm.qtpl:46
				qw422016.N().S(`": "`)
				//line views/templates/json/anyForm.qtpl:46
				qw422016.N().J(funcName)
				//line views/templates/json/anyForm.qtpl:46
				qw422016.N().S(`",`)
				//line views/templates/json/anyForm.qtpl:47
			}
			//line views/templates/json/anyForm.qtpl:47
			qw422016.N().S(`"count":`)
			//line views/templates/json/anyForm.qtpl:48
			qw422016.N().D(len(field.Events))
			//line views/templates/json/anyForm.qtpl:48
			qw422016.N().S(`},`)
			//line views/templates/json/anyForm.qtpl:50
		}
		//line views/templates/json/anyForm.qtpl:51
		if field.COLUMN_DEFAULT > "" {
			//line views/templates/json/anyForm.qtpl:51
			qw422016.N().S(`"default": "`)
			//line views/templates/json/anyForm.qtpl:51
			qw422016.N().J(field.COLUMN_DEFAULT)
			//line views/templates/json/anyForm.qtpl:51
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:51
		}
		//line views/templates/json/anyForm.qtpl:52
		if field.CHARACTER_SET_NAME > "" {
			//line views/templates/json/anyForm.qtpl:52
			qw422016.N().S(`"charSet": "`)
			//line views/templates/json/anyForm.qtpl:52
			qw422016.E().S(field.CHARACTER_SET_NAME)
			//line views/templates/json/anyForm.qtpl:52
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:52
		}
		//line views/templates/json/anyForm.qtpl:53
		if len(field.EnumValues) > 0 {
			//line views/templates/json/anyForm.qtpl:53
			qw422016.N().S(`"list": [`)
			//line views/templates/json/anyForm.qtpl:54
			for i, val := range field.EnumValues {
				//line views/templates/json/anyForm.qtpl:54
				if i > 0 {
					//line views/templates/json/anyForm.qtpl:54
					qw422016.N().S(`,`)
					//line views/templates/json/anyForm.qtpl:54
				}
				//line views/templates/json/anyForm.qtpl:54
				qw422016.N().S(`"`)
				//line views/templates/json/anyForm.qtpl:54
				qw422016.N().J(val)
				//line views/templates/json/anyForm.qtpl:54
				qw422016.N().S(`"`)
				//line views/templates/json/anyForm.qtpl:54
			}
			//line views/templates/json/anyForm.qtpl:54
			qw422016.N().S(`],`)
			//line views/templates/json/anyForm.qtpl:55
		} else if len(field.SelectValues) > 0 {
			//line views/templates/json/anyForm.qtpl:55
			qw422016.N().S(`"list": {`)
			//line views/templates/json/anyForm.qtpl:57
			for key, val := range field.SelectValues {
				//line views/templates/json/anyForm.qtpl:57
				qw422016.N().S(`"`)
				//line views/templates/json/anyForm.qtpl:58
				qw422016.N().D(key)
				//line views/templates/json/anyForm.qtpl:58
				qw422016.N().S(`":"`)
				//line views/templates/json/anyForm.qtpl:58
				qw422016.N().J(val)
				//line views/templates/json/anyForm.qtpl:58
				qw422016.N().S(`",`)
				//line views/templates/json/anyForm.qtpl:59
			}
			//line views/templates/json/anyForm.qtpl:59
			qw422016.N().S(`"count":`)
			//line views/templates/json/anyForm.qtpl:60
			qw422016.N().D(len(field.SelectValues))
			//line views/templates/json/anyForm.qtpl:60
			qw422016.N().S(`},`)
			//line views/templates/json/anyForm.qtpl:62
		}
		//line views/templates/json/anyForm.qtpl:64
		if field.TABLEID {
			//line views/templates/json/anyForm.qtpl:67
			tableProps := strings.TrimPrefix(field.COLUMN_NAME, "tableid_")
			fields := schema.GetFieldsTable(tableProps)

			//line views/templates/json/anyForm.qtpl:70
			qw422016.N().S(`"list": {`)
			//line views/templates/json/anyForm.qtpl:71
			StreamCheckFields(qw422016, fields)
			//line views/templates/json/anyForm.qtpl:71
			qw422016.N().S(`},`)
			//line views/templates/json/anyForm.qtpl:72
		}
		//line views/templates/json/anyForm.qtpl:72
		qw422016.N().S(`"type": "`)
		//line views/templates/json/anyForm.qtpl:74
		qw422016.N().J(typeInput)
		//line views/templates/json/anyForm.qtpl:74
		qw422016.N().S(`"},`)
		//line views/templates/json/anyForm.qtpl:76
	}
	//line views/templates/json/anyForm.qtpl:76
	qw422016.N().S(`"count":`)
	//line views/templates/json/anyForm.qtpl:77
	qw422016.N().D(len(ns.Rows))
//line views/templates/json/anyForm.qtpl:79
}

//line views/templates/json/anyForm.qtpl:79
func WriteCheckFields(qq422016 qtio422016.Writer, ns schema.FieldsTable) {
	//line views/templates/json/anyForm.qtpl:79
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/json/anyForm.qtpl:79
	StreamCheckFields(qw422016, ns)
	//line views/templates/json/anyForm.qtpl:79
	qt422016.ReleaseWriter(qw422016)
//line views/templates/json/anyForm.qtpl:79
}

//line views/templates/json/anyForm.qtpl:79
func CheckFields(ns schema.FieldsTable) string {
	//line views/templates/json/anyForm.qtpl:79
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/json/anyForm.qtpl:79
	WriteCheckFields(qb422016, ns)
	//line views/templates/json/anyForm.qtpl:79
	qs422016 := string(qb422016.B)
	//line views/templates/json/anyForm.qtpl:79
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/json/anyForm.qtpl:79
	return qs422016
//line views/templates/json/anyForm.qtpl:79
}

//line views/templates/json/anyForm.qtpl:82
func (thisForm *FormStructure) StreamJSONAnyForm(qw422016 *qt422016.Writer, ns schema.FieldsTable, AddJson map[string]string) {
	//line views/templates/json/anyForm.qtpl:84
	thisForm.setFormDefaults(&ns)

	//line views/templates/json/anyForm.qtpl:85
	qw422016.N().S(`{"fields": {`)
	//line views/templates/json/anyForm.qtpl:88
	StreamCheckFields(qw422016, ns)
	//line views/templates/json/anyForm.qtpl:88
	qw422016.N().S(`},"form": {`)
	//line views/templates/json/anyForm.qtpl:92
	if thisForm.ClassCSS != "" {
		//line views/templates/json/anyForm.qtpl:92
		qw422016.N().S(`"class" : "`)
		//line views/templates/json/anyForm.qtpl:92
		qw422016.N().J(thisForm.ClassCSS)
		//line views/templates/json/anyForm.qtpl:92
		qw422016.N().S(`",`)
		//line views/templates/json/anyForm.qtpl:92
	}
	//line views/templates/json/anyForm.qtpl:92
	qw422016.N().S(`"action" : "`)
	//line views/templates/json/anyForm.qtpl:93
	qw422016.N().J(thisForm.Action)
	//line views/templates/json/anyForm.qtpl:93
	qw422016.N().S(`","id" : "`)
	//line views/templates/json/anyForm.qtpl:94
	qw422016.N().J(thisForm.IdCSS)
	//line views/templates/json/anyForm.qtpl:94
	qw422016.N().S(`","name": "`)
	//line views/templates/json/anyForm.qtpl:95
	qw422016.N().J(thisForm.Name)
	//line views/templates/json/anyForm.qtpl:95
	qw422016.N().S(`"`)
	//line views/templates/json/anyForm.qtpl:96
	for name, event := range thisForm.Events {
		//line views/templates/json/anyForm.qtpl:96
		qw422016.N().S(`,"`)
		//line views/templates/json/anyForm.qtpl:97
		qw422016.E().S(name)
		//line views/templates/json/anyForm.qtpl:97
		qw422016.N().S(`": "`)
		//line views/templates/json/anyForm.qtpl:97
		qw422016.N().J(event)
		//line views/templates/json/anyForm.qtpl:97
		qw422016.N().S(`"`)
		//line views/templates/json/anyForm.qtpl:98
	}
	//line views/templates/json/anyForm.qtpl:98
	qw422016.N().S(`}`)
	//line views/templates/json/anyForm.qtpl:100
	for name, val := range AddJson {
		//line views/templates/json/anyForm.qtpl:100
		qw422016.N().S(`,"`)
		//line views/templates/json/anyForm.qtpl:101
		qw422016.E().S(name)
		//line views/templates/json/anyForm.qtpl:101
		qw422016.N().S(`":`)
		//line views/templates/json/anyForm.qtpl:101
		qw422016.N().S(val)
		//line views/templates/json/anyForm.qtpl:102
	}
	//line views/templates/json/anyForm.qtpl:102
	qw422016.N().S(`}`)
//line views/templates/json/anyForm.qtpl:105
}

//line views/templates/json/anyForm.qtpl:105
func (thisForm *FormStructure) WriteJSONAnyForm(qq422016 qtio422016.Writer, ns schema.FieldsTable, AddJson map[string]string) {
	//line views/templates/json/anyForm.qtpl:105
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/json/anyForm.qtpl:105
	thisForm.StreamJSONAnyForm(qw422016, ns, AddJson)
	//line views/templates/json/anyForm.qtpl:105
	qt422016.ReleaseWriter(qw422016)
//line views/templates/json/anyForm.qtpl:105
}

//line views/templates/json/anyForm.qtpl:105
func (thisForm *FormStructure) JSONAnyForm(ns schema.FieldsTable, AddJson map[string]string) string {
	//line views/templates/json/anyForm.qtpl:105
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/json/anyForm.qtpl:105
	thisForm.WriteJSONAnyForm(qb422016, ns, AddJson)
	//line views/templates/json/anyForm.qtpl:105
	qs422016 := string(qb422016.B)
	//line views/templates/json/anyForm.qtpl:105
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/json/anyForm.qtpl:105
	return qs422016
//line views/templates/json/anyForm.qtpl:105
}
