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
	"github.com/ruslanBik4/httpgo/models/db/qb"
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
func StreamCheckFields(qw422016 *qt422016.Writer, fields map[string]*qb.QBField) {
	//line views/templates/json/anyForm.qtpl:16
	for COLUMN_NAME, qbField := range fields {
		//line views/templates/json/anyForm.qtpl:20
		field := qbField.GetSchema()
		titleFull, titleLabel, placeholder, pattern, dataJson := field.GetColumnTitles()
		typeInput := field.TypeInput()

		//line views/templates/json/anyForm.qtpl:23
		qw422016.N().S(`"`)
		//line views/templates/json/anyForm.qtpl:24
		qw422016.E().S(field.COLUMN_NAME)
		//line views/templates/json/anyForm.qtpl:24
		qw422016.N().S(`": {`)
		//line views/templates/json/anyForm.qtpl:28
		if typeInput == "checkbox" {
			//line views/templates/json/anyForm.qtpl:28
			qw422016.N().S(`"value": 1,`)
			//line views/templates/json/anyForm.qtpl:28
		}
		//line views/templates/json/anyForm.qtpl:30
		if (COLUMN_NAME == "id") || field.IsHidden {
			//line views/templates/json/anyForm.qtpl:30
			qw422016.N().S(`"type": "hidden" },`)
			//line views/templates/json/anyForm.qtpl:32
			continue
			//line views/templates/json/anyForm.qtpl:33
		}
		//line views/templates/json/anyForm.qtpl:35
		if field.IS_NULLABLE == "NO" {
			//line views/templates/json/anyForm.qtpl:35
			qw422016.N().S(`"required": true,`)
			//line views/templates/json/anyForm.qtpl:35
		}
		//line views/templates/json/anyForm.qtpl:36
		if titleFull > "" {
			//line views/templates/json/anyForm.qtpl:36
			qw422016.N().S(`"title": "`)
			//line views/templates/json/anyForm.qtpl:36
			qw422016.E().J(titleFull)
			//line views/templates/json/anyForm.qtpl:36
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:36
		}
		//line views/templates/json/anyForm.qtpl:37
		if titleLabel > "" {
			//line views/templates/json/anyForm.qtpl:37
			qw422016.N().S(`"label": "`)
			//line views/templates/json/anyForm.qtpl:37
			qw422016.E().J(titleLabel)
			//line views/templates/json/anyForm.qtpl:37
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:37
		}
		//line views/templates/json/anyForm.qtpl:38
		if placeholder > "" {
			//line views/templates/json/anyForm.qtpl:38
			qw422016.N().S(`"placeholder": "`)
			//line views/templates/json/anyForm.qtpl:38
			qw422016.E().J(placeholder)
			//line views/templates/json/anyForm.qtpl:38
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:38
		}
		//line views/templates/json/anyForm.qtpl:39
		if pattern > "" {
			//line views/templates/json/anyForm.qtpl:39
			qw422016.N().S(`"pattern": "`)
			//line views/templates/json/anyForm.qtpl:39
			qw422016.N().J(pattern)
			//line views/templates/json/anyForm.qtpl:39
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:39
		}
		//line views/templates/json/anyForm.qtpl:40
		if dataJson > "" {
			//line views/templates/json/anyForm.qtpl:40
			qw422016.N().S(dataJson)
			//line views/templates/json/anyForm.qtpl:40
			qw422016.N().S(`,`)
			//line views/templates/json/anyForm.qtpl:40
		}
		//line views/templates/json/anyForm.qtpl:41
		if field.CSSClass > "" {
			//line views/templates/json/anyForm.qtpl:41
			qw422016.N().S(`"CSSClass": "`)
			//line views/templates/json/anyForm.qtpl:41
			qw422016.E().J(field.CSSClass)
			//line views/templates/json/anyForm.qtpl:41
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:41
		}
		//line views/templates/json/anyForm.qtpl:42
		if field.CSSStyle > "" {
			//line views/templates/json/anyForm.qtpl:42
			qw422016.N().S(`"CSSStyle": "`)
			//line views/templates/json/anyForm.qtpl:42
			qw422016.E().J(field.CSSStyle)
			//line views/templates/json/anyForm.qtpl:42
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:42
		}
		//line views/templates/json/anyForm.qtpl:43
		if field.Figure > "" {
			//line views/templates/json/anyForm.qtpl:43
			qw422016.N().S(`"Figure": "`)
			//line views/templates/json/anyForm.qtpl:43
			qw422016.E().J(field.Figure)
			//line views/templates/json/anyForm.qtpl:43
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:43
		}
		//line views/templates/json/anyForm.qtpl:44
		if field.CHARACTER_MAXIMUM_LENGTH > 0 {
			//line views/templates/json/anyForm.qtpl:44
			qw422016.N().S(`"maxLength":`)
			//line views/templates/json/anyForm.qtpl:44
			qw422016.N().D(field.CHARACTER_MAXIMUM_LENGTH)
			//line views/templates/json/anyForm.qtpl:44
			qw422016.N().S(`,`)
			//line views/templates/json/anyForm.qtpl:44
		}
		//line views/templates/json/anyForm.qtpl:45
		if strings.Contains(field.COLUMN_TYPE, "unsigned") {
			//line views/templates/json/anyForm.qtpl:45
			qw422016.N().S(`"min":0,`)
			//line views/templates/json/anyForm.qtpl:45
		}
		//line views/templates/json/anyForm.qtpl:46
		if field.MinDate > "" {
			//line views/templates/json/anyForm.qtpl:46
			qw422016.N().S(`"minDate": "`)
			//line views/templates/json/anyForm.qtpl:46
			qw422016.N().J(field.MinDate)
			//line views/templates/json/anyForm.qtpl:46
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:46
		}
		//line views/templates/json/anyForm.qtpl:47
		if field.MaxDate > "" {
			//line views/templates/json/anyForm.qtpl:47
			qw422016.N().S(`"maxDate": "`)
			//line views/templates/json/anyForm.qtpl:47
			qw422016.N().J(field.MaxDate)
			//line views/templates/json/anyForm.qtpl:47
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:47
		}
		//line views/templates/json/anyForm.qtpl:48
		if len(field.Events) > 0 {
			//line views/templates/json/anyForm.qtpl:48
			qw422016.N().S(`"Events": {`)
			//line views/templates/json/anyForm.qtpl:50
			for name, funcName := range field.Events {
				//line views/templates/json/anyForm.qtpl:50
				qw422016.N().S(`"`)
				//line views/templates/json/anyForm.qtpl:51
				qw422016.E().S(name)
				//line views/templates/json/anyForm.qtpl:51
				qw422016.N().S(`": "`)
				//line views/templates/json/anyForm.qtpl:51
				qw422016.N().J(funcName)
				//line views/templates/json/anyForm.qtpl:51
				qw422016.N().S(`",`)
				//line views/templates/json/anyForm.qtpl:52
			}
			//line views/templates/json/anyForm.qtpl:52
			qw422016.N().S(`"count":`)
			//line views/templates/json/anyForm.qtpl:53
			qw422016.N().D(len(field.Events))
			//line views/templates/json/anyForm.qtpl:53
			qw422016.N().S(`},`)
			//line views/templates/json/anyForm.qtpl:55
		}
		//line views/templates/json/anyForm.qtpl:56
		if field.COLUMN_DEFAULT > "" {
			//line views/templates/json/anyForm.qtpl:56
			qw422016.N().S(`"default":`)
			//line views/templates/json/anyForm.qtpl:58
			switch field.DATA_TYPE {
			//line views/templates/json/anyForm.qtpl:59
			case "tinyint":
				//line views/templates/json/anyForm.qtpl:60
				if field.COLUMN_DEFAULT == "1" {
					//line views/templates/json/anyForm.qtpl:60
					qw422016.N().S(`table`)
					//line views/templates/json/anyForm.qtpl:62
				} else {
					//line views/templates/json/anyForm.qtpl:62
					qw422016.N().S(`false`)
					//line views/templates/json/anyForm.qtpl:64
				}
				//line views/templates/json/anyForm.qtpl:64
				qw422016.N().S(`}`)
			//line views/templates/json/anyForm.qtpl:65
			case "int", "uint", "int64":
				//line views/templates/json/anyForm.qtpl:66
				qw422016.N().S(field.COLUMN_DEFAULT)
			//line views/templates/json/anyForm.qtpl:67
			case "float", "double":
				//line views/templates/json/anyForm.qtpl:68
				qw422016.N().S(field.COLUMN_DEFAULT)
			//line views/templates/json/anyForm.qtpl:69
			default:
				//line views/templates/json/anyForm.qtpl:69
				qw422016.N().S(`"`)
				//line views/templates/json/anyForm.qtpl:70
				qw422016.N().J(field.COLUMN_DEFAULT)
				//line views/templates/json/anyForm.qtpl:70
				qw422016.N().S(`"`)
				//line views/templates/json/anyForm.qtpl:71
			}
			//line views/templates/json/anyForm.qtpl:71
			qw422016.N().S(`,`)
			//line views/templates/json/anyForm.qtpl:73
		}
		//line views/templates/json/anyForm.qtpl:74
		if field.CHARACTER_SET_NAME > "" {
			//line views/templates/json/anyForm.qtpl:74
			qw422016.N().S(`"charSet": "`)
			//line views/templates/json/anyForm.qtpl:74
			qw422016.E().S(field.CHARACTER_SET_NAME)
			//line views/templates/json/anyForm.qtpl:74
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:74
		}
		//line views/templates/json/anyForm.qtpl:75
		if len(qbField.SelectValues) > 0 {
			//line views/templates/json/anyForm.qtpl:75
			qw422016.N().S(`"list": {`)
			//line views/templates/json/anyForm.qtpl:77
			for key, val := range qbField.SelectValues {
				//line views/templates/json/anyForm.qtpl:77
				qw422016.N().S(`"`)
				//line views/templates/json/anyForm.qtpl:78
				qw422016.N().D(key)
				//line views/templates/json/anyForm.qtpl:78
				qw422016.N().S(`":"`)
				//line views/templates/json/anyForm.qtpl:78
				qw422016.N().J(val)
				//line views/templates/json/anyForm.qtpl:78
				qw422016.N().S(`",`)
				//line views/templates/json/anyForm.qtpl:79
			}
			//line views/templates/json/anyForm.qtpl:79
			qw422016.N().S(`"count":`)
			//line views/templates/json/anyForm.qtpl:80
			qw422016.N().D(len(qbField.SelectValues))
			//line views/templates/json/anyForm.qtpl:80
			qw422016.N().S(`},`)
			//line views/templates/json/anyForm.qtpl:86
		} else if len(field.EnumValues) > 0 {
			//line views/templates/json/anyForm.qtpl:86
			qw422016.N().S(`"list": [`)
			//line views/templates/json/anyForm.qtpl:87
			for i, val := range field.EnumValues {
				//line views/templates/json/anyForm.qtpl:87
				if i > 0 {
					//line views/templates/json/anyForm.qtpl:87
					qw422016.N().S(`,`)
					//line views/templates/json/anyForm.qtpl:87
				}
				//line views/templates/json/anyForm.qtpl:87
				qw422016.N().S(`"`)
				//line views/templates/json/anyForm.qtpl:87
				qw422016.N().J(val)
				//line views/templates/json/anyForm.qtpl:87
				qw422016.N().S(`"`)
				//line views/templates/json/anyForm.qtpl:87
			}
			//line views/templates/json/anyForm.qtpl:87
			qw422016.N().S(`],`)
			//line views/templates/json/anyForm.qtpl:88
		}
		//line views/templates/json/anyForm.qtpl:90
		if field.TABLEID {
			//line views/templates/json/anyForm.qtpl:90
			qw422016.N().S(`"list": {`)
			//line views/templates/json/anyForm.qtpl:91
			StreamCheckFields(qw422016, qbField.ChildQB.GetFields().Fields)
			//line views/templates/json/anyForm.qtpl:91
			qw422016.N().S(`},`)
			//line views/templates/json/anyForm.qtpl:92
		}
		//line views/templates/json/anyForm.qtpl:92
		qw422016.N().S(`"type": "`)
		//line views/templates/json/anyForm.qtpl:94
		qw422016.N().J(typeInput)
		//line views/templates/json/anyForm.qtpl:94
		qw422016.N().S(`"},`)
		//line views/templates/json/anyForm.qtpl:96
	}
	//line views/templates/json/anyForm.qtpl:96
	qw422016.N().S(`"count":`)
	//line views/templates/json/anyForm.qtpl:97
	qw422016.N().D(len(fields))
//line views/templates/json/anyForm.qtpl:99
}

//line views/templates/json/anyForm.qtpl:99
func WriteCheckFields(qq422016 qtio422016.Writer, fields map[string]*qb.QBField) {
	//line views/templates/json/anyForm.qtpl:99
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/json/anyForm.qtpl:99
	StreamCheckFields(qw422016, fields)
	//line views/templates/json/anyForm.qtpl:99
	qt422016.ReleaseWriter(qw422016)
//line views/templates/json/anyForm.qtpl:99
}

//line views/templates/json/anyForm.qtpl:99
func CheckFields(fields map[string]*qb.QBField) string {
	//line views/templates/json/anyForm.qtpl:99
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/json/anyForm.qtpl:99
	WriteCheckFields(qb422016, fields)
	//line views/templates/json/anyForm.qtpl:99
	qs422016 := string(qb422016.B)
	//line views/templates/json/anyForm.qtpl:99
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/json/anyForm.qtpl:99
	return qs422016
//line views/templates/json/anyForm.qtpl:99
}

//line views/templates/json/anyForm.qtpl:102
func (thisForm *FormStructure) StreamJSONAnyForm(qw422016 *qt422016.Writer, table qb.QBTable, arrJSON MultiDimension) {
	//line views/templates/json/anyForm.qtpl:104
	thisForm.setFormDefaults(table.GetSchema())

	//line views/templates/json/anyForm.qtpl:105
	qw422016.N().S(`{"fields": {`)
	//line views/templates/json/anyForm.qtpl:108
	StreamCheckFields(qw422016, table.Fields)
	//line views/templates/json/anyForm.qtpl:108
	qw422016.N().S(`},"form": {`)
	//line views/templates/json/anyForm.qtpl:112
	if thisForm.ClassCSS != "" {
		//line views/templates/json/anyForm.qtpl:112
		qw422016.N().S(`"class" : "`)
		//line views/templates/json/anyForm.qtpl:112
		qw422016.N().J(thisForm.ClassCSS)
		//line views/templates/json/anyForm.qtpl:112
		qw422016.N().S(`",`)
		//line views/templates/json/anyForm.qtpl:112
	}
	//line views/templates/json/anyForm.qtpl:112
	qw422016.N().S(`"action" : "`)
	//line views/templates/json/anyForm.qtpl:113
	qw422016.N().J(thisForm.Action)
	//line views/templates/json/anyForm.qtpl:113
	qw422016.N().S(`","id" : "`)
	//line views/templates/json/anyForm.qtpl:114
	qw422016.N().J(thisForm.IdCSS)
	//line views/templates/json/anyForm.qtpl:114
	qw422016.N().S(`","name": "`)
	//line views/templates/json/anyForm.qtpl:115
	qw422016.N().J(thisForm.Name)
	//line views/templates/json/anyForm.qtpl:115
	qw422016.N().S(`"`)
	//line views/templates/json/anyForm.qtpl:116
	for name, event := range thisForm.Events {
		//line views/templates/json/anyForm.qtpl:116
		qw422016.N().S(`,"`)
		//line views/templates/json/anyForm.qtpl:117
		qw422016.E().S(name)
		//line views/templates/json/anyForm.qtpl:117
		qw422016.N().S(`": "`)
		//line views/templates/json/anyForm.qtpl:117
		qw422016.N().J(event)
		//line views/templates/json/anyForm.qtpl:117
		qw422016.N().S(`"`)
		//line views/templates/json/anyForm.qtpl:118
	}
	//line views/templates/json/anyForm.qtpl:118
	qw422016.N().S(`}`)
	//line views/templates/json/anyForm.qtpl:121
	for key, value := range arrJSON {
		//line views/templates/json/anyForm.qtpl:121
		qw422016.N().S(`, "`)
		//line views/templates/json/anyForm.qtpl:122
		qw422016.E().S(key)
		//line views/templates/json/anyForm.qtpl:122
		qw422016.N().S(`":`)
		//line views/templates/json/anyForm.qtpl:122
		StreamElement(qw422016, value)
		//line views/templates/json/anyForm.qtpl:123
	}
	//line views/templates/json/anyForm.qtpl:123
	qw422016.N().S(`}`)
//line views/templates/json/anyForm.qtpl:125
}

//line views/templates/json/anyForm.qtpl:125
func (thisForm *FormStructure) WriteJSONAnyForm(qq422016 qtio422016.Writer, table qb.QBTable, arrJSON MultiDimension) {
	//line views/templates/json/anyForm.qtpl:125
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/json/anyForm.qtpl:125
	thisForm.StreamJSONAnyForm(qw422016, table, arrJSON)
	//line views/templates/json/anyForm.qtpl:125
	qt422016.ReleaseWriter(qw422016)
//line views/templates/json/anyForm.qtpl:125
}

//line views/templates/json/anyForm.qtpl:125
func (thisForm *FormStructure) JSONAnyForm(table qb.QBTable, arrJSON MultiDimension) string {
	//line views/templates/json/anyForm.qtpl:125
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/json/anyForm.qtpl:125
	thisForm.WriteJSONAnyForm(qb422016, table, arrJSON)
	//line views/templates/json/anyForm.qtpl:125
	qs422016 := string(qb422016.B)
	//line views/templates/json/anyForm.qtpl:125
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/json/anyForm.qtpl:125
	return qs422016
//line views/templates/json/anyForm.qtpl:125
}
