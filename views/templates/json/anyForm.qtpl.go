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
	"github.com/ruslanBik4/httpgo/models/db"
	"github.com/ruslanBik4/httpgo/views/templates/forms"
	"strings"
)

// Формируем JSON, который затем будет использован в форме на клиенте
// dataJson и содержимое AddJson вставляются КАК ЕСТЬ, ПОТОМУ ЧТО БЫЛИ ОБРАБОТАНЫ РАНЕЕ!!!

//line views/templates/json/anyForm.qtpl:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/json/anyForm.qtpl:14
func StreamCheckFields(qw422016 *qt422016.Writer, ns *forms.FieldsTable) {
	//line views/templates/json/anyForm.qtpl:16
	for _, field := range ns.Rows {
		//line views/templates/json/anyForm.qtpl:20
		titleFull, titleLabel, placeholder, pattern, dataJson := field.GetColumnTitles()
		typeInput := field.TypeInput()

		//line views/templates/json/anyForm.qtpl:22
		qw422016.N().S(`"`)
		//line views/templates/json/anyForm.qtpl:23
		qw422016.E().S(field.COLUMN_NAME)
		//line views/templates/json/anyForm.qtpl:23
		qw422016.N().S(`": {`)
		//line views/templates/json/anyForm.qtpl:24
		if (field.COLUMN_NAME == "id") || field.IsHidden {
			//line views/templates/json/anyForm.qtpl:24
			qw422016.N().S(`"type": "hidden" },`)
			//line views/templates/json/anyForm.qtpl:26
			continue
			//line views/templates/json/anyForm.qtpl:27
		}
		//line views/templates/json/anyForm.qtpl:29
		if field.IS_NULLABLE == "NO" {
			//line views/templates/json/anyForm.qtpl:29
			qw422016.N().S(`"required": true,`)
			//line views/templates/json/anyForm.qtpl:29
		}
		//line views/templates/json/anyForm.qtpl:30
		if titleFull > "" {
			//line views/templates/json/anyForm.qtpl:30
			qw422016.N().S(`"title": "`)
			//line views/templates/json/anyForm.qtpl:30
			qw422016.E().J(titleFull)
			//line views/templates/json/anyForm.qtpl:30
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:30
		}
		//line views/templates/json/anyForm.qtpl:31
		if titleLabel > "" {
			//line views/templates/json/anyForm.qtpl:31
			qw422016.N().S(`"label": "`)
			//line views/templates/json/anyForm.qtpl:31
			qw422016.E().J(titleLabel)
			//line views/templates/json/anyForm.qtpl:31
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:31
		}
		//line views/templates/json/anyForm.qtpl:32
		if placeholder > "" {
			//line views/templates/json/anyForm.qtpl:32
			qw422016.N().S(`"placeholder": "`)
			//line views/templates/json/anyForm.qtpl:32
			qw422016.E().S(placeholder)
			//line views/templates/json/anyForm.qtpl:32
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:32
		}
		//line views/templates/json/anyForm.qtpl:33
		if pattern > "" {
			//line views/templates/json/anyForm.qtpl:33
			qw422016.N().S(`"pattern": "`)
			//line views/templates/json/anyForm.qtpl:33
			qw422016.E().J(pattern)
			//line views/templates/json/anyForm.qtpl:33
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:33
		}
		//line views/templates/json/anyForm.qtpl:34
		if dataJson > "" {
			//line views/templates/json/anyForm.qtpl:34
			qw422016.N().S(dataJson)
			//line views/templates/json/anyForm.qtpl:34
			qw422016.N().S(`,`)
			//line views/templates/json/anyForm.qtpl:34
		}
		//line views/templates/json/anyForm.qtpl:35
		if field.Value > "" {
			//line views/templates/json/anyForm.qtpl:35
			qw422016.N().S(`"value": "`)
			//line views/templates/json/anyForm.qtpl:35
			qw422016.E().S(field.Value)
			//line views/templates/json/anyForm.qtpl:35
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:35
		}
		//line views/templates/json/anyForm.qtpl:36
		if field.CSSClass > "" {
			//line views/templates/json/anyForm.qtpl:36
			qw422016.N().S(`"CSSClass": "`)
			//line views/templates/json/anyForm.qtpl:36
			qw422016.E().S(field.CSSClass)
			//line views/templates/json/anyForm.qtpl:36
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:36
		}
		//line views/templates/json/anyForm.qtpl:37
		if field.CSSStyle > "" {
			//line views/templates/json/anyForm.qtpl:37
			qw422016.N().S(`"CSSStyle": "`)
			//line views/templates/json/anyForm.qtpl:37
			qw422016.E().S(field.CSSStyle)
			//line views/templates/json/anyForm.qtpl:37
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:37
		}
		//line views/templates/json/anyForm.qtpl:38
		if field.Figure > "" {
			//line views/templates/json/anyForm.qtpl:38
			qw422016.N().S(`"Figure": "`)
			//line views/templates/json/anyForm.qtpl:38
			qw422016.E().S(field.Figure)
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
			qw422016.E().S(field.MinDate)
			//line views/templates/json/anyForm.qtpl:41
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:41
		}
		//line views/templates/json/anyForm.qtpl:42
		if field.MaxDate > "" {
			//line views/templates/json/anyForm.qtpl:42
			qw422016.N().S(`"maxDate": "`)
			//line views/templates/json/anyForm.qtpl:42
			qw422016.E().S(field.MaxDate)
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
				qw422016.E().S(funcName)
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
		if len(field.EnumValues) > 0 {
			//line views/templates/json/anyForm.qtpl:51
			qw422016.N().S(`"list": [`)
			//line views/templates/json/anyForm.qtpl:52
			for i, val := range field.EnumValues {
				//line views/templates/json/anyForm.qtpl:52
				if i > 0 {
					//line views/templates/json/anyForm.qtpl:52
					qw422016.N().S(`,`)
					//line views/templates/json/anyForm.qtpl:52
				}
				//line views/templates/json/anyForm.qtpl:52
				qw422016.N().S(`"`)
				//line views/templates/json/anyForm.qtpl:52
				qw422016.E().S(val)
				//line views/templates/json/anyForm.qtpl:52
				qw422016.N().S(`"`)
				//line views/templates/json/anyForm.qtpl:52
			}
			//line views/templates/json/anyForm.qtpl:52
			qw422016.N().S(`],`)
			//line views/templates/json/anyForm.qtpl:53
		}
		//line views/templates/json/anyForm.qtpl:54
		if field.COLUMN_DEFAULT > "" {
			//line views/templates/json/anyForm.qtpl:54
			qw422016.N().S(`"default": "`)
			//line views/templates/json/anyForm.qtpl:54
			qw422016.E().S(field.COLUMN_DEFAULT)
			//line views/templates/json/anyForm.qtpl:54
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:54
		}
		//line views/templates/json/anyForm.qtpl:55
		if field.CHARACTER_SET_NAME > "" {
			//line views/templates/json/anyForm.qtpl:55
			qw422016.N().S(`"charSet": "`)
			//line views/templates/json/anyForm.qtpl:55
			qw422016.E().S(field.CHARACTER_SET_NAME)
			//line views/templates/json/anyForm.qtpl:55
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:55
		}
		//line views/templates/json/anyForm.qtpl:55
		qw422016.N().S(`"type": "`)
		//line views/templates/json/anyForm.qtpl:56
		qw422016.E().S(typeInput)
		//line views/templates/json/anyForm.qtpl:56
		qw422016.N().S(`"`)
		//line views/templates/json/anyForm.qtpl:58
		if strings.HasPrefix(field.COLUMN_NAME, "id_") {
			//line views/templates/json/anyForm.qtpl:58
			qw422016.N().S(`,`)
			//line views/templates/json/anyForm.qtpl:61
			field.GetOptionsJson(field.COLUMN_NAME[3:])

			//line views/templates/json/anyForm.qtpl:62
			qw422016.N().S(`"list":{`)
			//line views/templates/json/anyForm.qtpl:63
			qw422016.N().S(field.Html)
			//line views/templates/json/anyForm.qtpl:63
			qw422016.N().S(`}`)
			//line views/templates/json/anyForm.qtpl:65
		} else if strings.HasPrefix(field.COLUMN_NAME, "setid_") || strings.HasPrefix(field.COLUMN_NAME, "nodeid_") {
			//line views/templates/json/anyForm.qtpl:65
			qw422016.N().S(`,`)
			//line views/templates/json/anyForm.qtpl:68
			field.GetOptionsJson(field.COLUMN_NAME[6:])

			//line views/templates/json/anyForm.qtpl:69
			qw422016.N().S(`"list":{`)
			//line views/templates/json/anyForm.qtpl:70
			qw422016.N().S(field.Html)
			//line views/templates/json/anyForm.qtpl:70
			qw422016.N().S(`}`)
			//line views/templates/json/anyForm.qtpl:72
		} else if strings.HasPrefix(field.COLUMN_NAME, "tableid_") {
			//line views/templates/json/anyForm.qtpl:72
			qw422016.N().S(`,`)
			//line views/templates/json/anyForm.qtpl:76
			var fields forms.FieldsTable
			var nx db.FieldsTable

			fields.Name = field.COLUMN_NAME[8:]
			nx.Options.GetTableProp(fields.Name)
			nx.GetColumnsProp(fields.Name)

			fields.PutDataFrom(nx)

			for childKey, childFields := range fields.Rows {
				if strings.HasPrefix(childFields.COLUMN_NAME, "id_") && childFields.COLUMN_NAME[3:] == field.Table.Name {

					fields.Rows = append(fields.Rows[:childKey], fields.Rows[childKey+1:]...)
					break
				}
			}

			//line views/templates/json/anyForm.qtpl:92
			qw422016.N().S(`"list": {`)
			//line views/templates/json/anyForm.qtpl:93
			StreamCheckFields(qw422016, &fields)
			//line views/templates/json/anyForm.qtpl:93
			qw422016.N().S(`}`)
			//line views/templates/json/anyForm.qtpl:94
		}
		//line views/templates/json/anyForm.qtpl:94
		qw422016.N().S(`},`)
		//line views/templates/json/anyForm.qtpl:97
	}
	//line views/templates/json/anyForm.qtpl:97
	qw422016.N().S(`"count":`)
	//line views/templates/json/anyForm.qtpl:98
	qw422016.N().D(len(ns.Rows))
//line views/templates/json/anyForm.qtpl:100
}

//line views/templates/json/anyForm.qtpl:100
func WriteCheckFields(qq422016 qtio422016.Writer, ns *forms.FieldsTable) {
	//line views/templates/json/anyForm.qtpl:100
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/json/anyForm.qtpl:100
	StreamCheckFields(qw422016, ns)
	//line views/templates/json/anyForm.qtpl:100
	qt422016.ReleaseWriter(qw422016)
//line views/templates/json/anyForm.qtpl:100
}

//line views/templates/json/anyForm.qtpl:100
func CheckFields(ns *forms.FieldsTable) string {
	//line views/templates/json/anyForm.qtpl:100
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/json/anyForm.qtpl:100
	WriteCheckFields(qb422016, ns)
	//line views/templates/json/anyForm.qtpl:100
	qs422016 := string(qb422016.B)
	//line views/templates/json/anyForm.qtpl:100
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/json/anyForm.qtpl:100
	return qs422016
//line views/templates/json/anyForm.qtpl:100
}

//line views/templates/json/anyForm.qtpl:103
func (thisForm *FormStructure) StreamJSONAnyForm(qw422016 *qt422016.Writer, ns *forms.FieldsTable, AddJson map[string]string) {
	//line views/templates/json/anyForm.qtpl:105
	thisForm.setFormDefaults(ns)

	//line views/templates/json/anyForm.qtpl:106
	qw422016.N().S(`{"fields": {`)
	//line views/templates/json/anyForm.qtpl:109
	StreamCheckFields(qw422016, ns)
	//line views/templates/json/anyForm.qtpl:109
	qw422016.N().S(`},"form": {`)
	//line views/templates/json/anyForm.qtpl:113
	if thisForm.ClassCSS != "" {
		//line views/templates/json/anyForm.qtpl:113
		qw422016.N().S(`"class" : "`)
		//line views/templates/json/anyForm.qtpl:113
		qw422016.N().S(thisForm.ClassCSS)
		//line views/templates/json/anyForm.qtpl:113
		qw422016.N().S(`",`)
		//line views/templates/json/anyForm.qtpl:113
	}
	//line views/templates/json/anyForm.qtpl:113
	qw422016.N().S(`"action" : "`)
	//line views/templates/json/anyForm.qtpl:114
	qw422016.N().S(thisForm.Action)
	//line views/templates/json/anyForm.qtpl:114
	qw422016.N().S(`","id" : "`)
	//line views/templates/json/anyForm.qtpl:115
	qw422016.N().S(thisForm.IdCSS)
	//line views/templates/json/anyForm.qtpl:115
	qw422016.N().S(`","name": "`)
	//line views/templates/json/anyForm.qtpl:116
	qw422016.N().S(thisForm.Name)
	//line views/templates/json/anyForm.qtpl:116
	qw422016.N().S(`"`)
	//line views/templates/json/anyForm.qtpl:117
	for name, event := range thisForm.Events {
		//line views/templates/json/anyForm.qtpl:117
		qw422016.N().S(`,"`)
		//line views/templates/json/anyForm.qtpl:118
		qw422016.E().S(name)
		//line views/templates/json/anyForm.qtpl:118
		qw422016.N().S(`": "`)
		//line views/templates/json/anyForm.qtpl:118
		qw422016.N().S(event)
		//line views/templates/json/anyForm.qtpl:118
		qw422016.N().S(`"`)
		//line views/templates/json/anyForm.qtpl:119
	}
	//line views/templates/json/anyForm.qtpl:119
	qw422016.N().S(`}`)
	//line views/templates/json/anyForm.qtpl:121
	for name, val := range AddJson {
		//line views/templates/json/anyForm.qtpl:121
		qw422016.N().S(`,"`)
		//line views/templates/json/anyForm.qtpl:122
		qw422016.E().S(name)
		//line views/templates/json/anyForm.qtpl:122
		qw422016.N().S(`":`)
		//line views/templates/json/anyForm.qtpl:122
		qw422016.N().S(val)
		//line views/templates/json/anyForm.qtpl:123
	}
	//line views/templates/json/anyForm.qtpl:123
	qw422016.N().S(`}`)
//line views/templates/json/anyForm.qtpl:126
}

//line views/templates/json/anyForm.qtpl:126
func (thisForm *FormStructure) WriteJSONAnyForm(qq422016 qtio422016.Writer, ns *forms.FieldsTable, AddJson map[string]string) {
	//line views/templates/json/anyForm.qtpl:126
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/json/anyForm.qtpl:126
	thisForm.StreamJSONAnyForm(qw422016, ns, AddJson)
	//line views/templates/json/anyForm.qtpl:126
	qt422016.ReleaseWriter(qw422016)
//line views/templates/json/anyForm.qtpl:126
}

//line views/templates/json/anyForm.qtpl:126
func (thisForm *FormStructure) JSONAnyForm(ns *forms.FieldsTable, AddJson map[string]string) string {
	//line views/templates/json/anyForm.qtpl:126
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/json/anyForm.qtpl:126
	thisForm.WriteJSONAnyForm(qb422016, ns, AddJson)
	//line views/templates/json/anyForm.qtpl:126
	qs422016 := string(qb422016.B)
	//line views/templates/json/anyForm.qtpl:126
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/json/anyForm.qtpl:126
	return qs422016
//line views/templates/json/anyForm.qtpl:126
}
