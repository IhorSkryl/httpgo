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
	"github.com/ruslanBik4/httpgo/views/templates/forms"
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
func (thisForm *FormStructure) StreamJSONAnyForm(qw422016 *qt422016.Writer, ns *forms.FieldsTable, AddJson map[string]string) {
	//line views/templates/json/anyForm.qtpl:15
	thisForm.setFormDefaults(ns)
	if _, ok := ns.SaveFormEvents["errorFunction"]; !ok {
		ns.SaveFormEvents["errorFunction"] = ""
	}

	//line views/templates/json/anyForm.qtpl:19
	qw422016.N().S(`{"fields": {`)
	//line views/templates/json/anyForm.qtpl:22
	for _, field := range ns.Rows {
		//line views/templates/json/anyForm.qtpl:24
		titleFull, titleLabel, placeholder, pattern, dataJson := field.GetColumnTitles()
		typeInput := field.TypeInput()

		//line views/templates/json/anyForm.qtpl:26
		qw422016.N().S(`"`)
		//line views/templates/json/anyForm.qtpl:27
		qw422016.E().S(field.COLUMN_NAME)
		//line views/templates/json/anyForm.qtpl:27
		qw422016.N().S(`": {`)
		//line views/templates/json/anyForm.qtpl:28
		if field.IS_NULLABLE == "NO" {
			//line views/templates/json/anyForm.qtpl:28
			qw422016.N().S(`"required": true,`)
			//line views/templates/json/anyForm.qtpl:28
		}
		//line views/templates/json/anyForm.qtpl:29
		if titleFull > "" {
			//line views/templates/json/anyForm.qtpl:29
			qw422016.N().S(`"title": "`)
			//line views/templates/json/anyForm.qtpl:29
			qw422016.E().J(titleFull)
			//line views/templates/json/anyForm.qtpl:29
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:29
		}
		//line views/templates/json/anyForm.qtpl:30
		if titleLabel > "" {
			//line views/templates/json/anyForm.qtpl:30
			qw422016.N().S(`"label": "`)
			//line views/templates/json/anyForm.qtpl:30
			qw422016.E().J(titleLabel)
			//line views/templates/json/anyForm.qtpl:30
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:30
		}
		//line views/templates/json/anyForm.qtpl:31
		if placeholder > "" {
			//line views/templates/json/anyForm.qtpl:31
			qw422016.N().S(`"placeholder": "`)
			//line views/templates/json/anyForm.qtpl:31
			qw422016.E().S(placeholder)
			//line views/templates/json/anyForm.qtpl:31
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:31
		}
		//line views/templates/json/anyForm.qtpl:32
		if pattern > "" {
			//line views/templates/json/anyForm.qtpl:32
			qw422016.N().S(`"pattern": "`)
			//line views/templates/json/anyForm.qtpl:32
			qw422016.E().J(pattern)
			//line views/templates/json/anyForm.qtpl:32
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:32
		}
		//line views/templates/json/anyForm.qtpl:33
		if dataJson > "" {
			//line views/templates/json/anyForm.qtpl:33
			qw422016.N().S(dataJson)
			//line views/templates/json/anyForm.qtpl:33
			qw422016.N().S(`,`)
			//line views/templates/json/anyForm.qtpl:33
		}
		//line views/templates/json/anyForm.qtpl:34
		if field.Value > "" {
			//line views/templates/json/anyForm.qtpl:34
			qw422016.N().S(`"value": "`)
			//line views/templates/json/anyForm.qtpl:34
			qw422016.E().S(field.Value)
			//line views/templates/json/anyForm.qtpl:34
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:34
		}
		//line views/templates/json/anyForm.qtpl:35
		if field.CSSClass > "" {
			//line views/templates/json/anyForm.qtpl:35
			qw422016.N().S(`"CSSClass": "`)
			//line views/templates/json/anyForm.qtpl:35
			qw422016.E().S(field.CSSClass)
			//line views/templates/json/anyForm.qtpl:35
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:35
		}
		//line views/templates/json/anyForm.qtpl:36
		if field.CSSStyle > "" {
			//line views/templates/json/anyForm.qtpl:36
			qw422016.N().S(`"CSSStyle": "`)
			//line views/templates/json/anyForm.qtpl:36
			qw422016.E().S(field.CSSStyle)
			//line views/templates/json/anyForm.qtpl:36
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:36
		}
		//line views/templates/json/anyForm.qtpl:37
		if field.Figure > "" {
			//line views/templates/json/anyForm.qtpl:37
			qw422016.N().S(`"Figure": "`)
			//line views/templates/json/anyForm.qtpl:37
			qw422016.E().S(field.Figure)
			//line views/templates/json/anyForm.qtpl:37
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:37
		}
		//line views/templates/json/anyForm.qtpl:38
		if field.CHARACTER_MAXIMUM_LENGTH > 0 {
			//line views/templates/json/anyForm.qtpl:38
			qw422016.N().S(`"maxlength":`)
			//line views/templates/json/anyForm.qtpl:38
			qw422016.N().D(field.CHARACTER_MAXIMUM_LENGTH)
			//line views/templates/json/anyForm.qtpl:38
			qw422016.N().S(`,`)
			//line views/templates/json/anyForm.qtpl:38
		}
		//line views/templates/json/anyForm.qtpl:39
		if strings.Contains(field.COLUMN_TYPE, "unsigned") {
			//line views/templates/json/anyForm.qtpl:39
			qw422016.N().S(`"min":0,`)
			//line views/templates/json/anyForm.qtpl:39
		}
		//line views/templates/json/anyForm.qtpl:40
		if field.MinDate > "" {
			//line views/templates/json/anyForm.qtpl:40
			qw422016.N().S(`"minDate": "`)
			//line views/templates/json/anyForm.qtpl:40
			qw422016.E().S(field.MinDate)
			//line views/templates/json/anyForm.qtpl:40
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:40
		}
		//line views/templates/json/anyForm.qtpl:41
		if field.MaxDate > "" {
			//line views/templates/json/anyForm.qtpl:41
			qw422016.N().S(`"maxDate": "`)
			//line views/templates/json/anyForm.qtpl:41
			qw422016.E().S(field.MaxDate)
			//line views/templates/json/anyForm.qtpl:41
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:41
		}
		//line views/templates/json/anyForm.qtpl:42
		if len(field.Events) > 0 {
			//line views/templates/json/anyForm.qtpl:42
			qw422016.N().S(`"Events": {`)
			//line views/templates/json/anyForm.qtpl:44
			for name, funcName := range field.Events {
				//line views/templates/json/anyForm.qtpl:44
				qw422016.N().S(`"`)
				//line views/templates/json/anyForm.qtpl:45
				qw422016.E().S(name)
				//line views/templates/json/anyForm.qtpl:45
				qw422016.N().S(`": "`)
				//line views/templates/json/anyForm.qtpl:45
				qw422016.E().S(funcName)
				//line views/templates/json/anyForm.qtpl:45
				qw422016.N().S(`",`)
				//line views/templates/json/anyForm.qtpl:46
			}
			//line views/templates/json/anyForm.qtpl:46
			qw422016.N().S(`"count":`)
			//line views/templates/json/anyForm.qtpl:47
			qw422016.N().D(len(field.Events))
			//line views/templates/json/anyForm.qtpl:47
			qw422016.N().S(`},`)
			//line views/templates/json/anyForm.qtpl:49
		}
		//line views/templates/json/anyForm.qtpl:50
		if len(field.EnumValues) > 0 {
			//line views/templates/json/anyForm.qtpl:50
			qw422016.N().S(`"EnumValues": [`)
			//line views/templates/json/anyForm.qtpl:51
			for i, val := range field.EnumValues {
				//line views/templates/json/anyForm.qtpl:51
				if i > 0 {
					//line views/templates/json/anyForm.qtpl:51
					qw422016.N().S(`,`)
					//line views/templates/json/anyForm.qtpl:51
				}
				//line views/templates/json/anyForm.qtpl:51
				qw422016.N().S(`"`)
				//line views/templates/json/anyForm.qtpl:51
				qw422016.E().S(val)
				//line views/templates/json/anyForm.qtpl:51
				qw422016.N().S(`"`)
				//line views/templates/json/anyForm.qtpl:51
			}
			//line views/templates/json/anyForm.qtpl:51
			qw422016.N().S(`],`)
			//line views/templates/json/anyForm.qtpl:52
		}
		//line views/templates/json/anyForm.qtpl:53
		if field.COLUMN_DEFAULT > "" {
			//line views/templates/json/anyForm.qtpl:53
			qw422016.N().S(`"COLUMN_DEFAULT": "`)
			//line views/templates/json/anyForm.qtpl:53
			qw422016.E().S(field.COLUMN_DEFAULT)
			//line views/templates/json/anyForm.qtpl:53
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:53
		}
		//line views/templates/json/anyForm.qtpl:54
		if field.CHARACTER_SET_NAME > "" {
			//line views/templates/json/anyForm.qtpl:54
			qw422016.N().S(`"CHARACTER_SET_NAME": "`)
			//line views/templates/json/anyForm.qtpl:54
			qw422016.E().S(field.CHARACTER_SET_NAME)
			//line views/templates/json/anyForm.qtpl:54
			qw422016.N().S(`",`)
			//line views/templates/json/anyForm.qtpl:54
		}
		//line views/templates/json/anyForm.qtpl:54
		qw422016.N().S(`"type": "`)
		//line views/templates/json/anyForm.qtpl:55
		qw422016.E().S(typeInput)
		//line views/templates/json/anyForm.qtpl:55
		qw422016.N().S(`"},`)
		//line views/templates/json/anyForm.qtpl:57
	}
	//line views/templates/json/anyForm.qtpl:57
	qw422016.N().S(`"count":`)
	//line views/templates/json/anyForm.qtpl:58
	qw422016.N().D(len(ns.Rows))
	//line views/templates/json/anyForm.qtpl:58
	qw422016.N().S(`}, "form": {`)
	//line views/templates/json/anyForm.qtpl:60
	if thisForm.ClassCSS != "" {
		//line views/templates/json/anyForm.qtpl:60
		qw422016.N().S(`"class" : "`)
		//line views/templates/json/anyForm.qtpl:60
		qw422016.N().S(thisForm.ClassCSS)
		//line views/templates/json/anyForm.qtpl:60
		qw422016.N().S(`",`)
		//line views/templates/json/anyForm.qtpl:60
	}
	//line views/templates/json/anyForm.qtpl:60
	qw422016.N().S(`"action" : "`)
	//line views/templates/json/anyForm.qtpl:61
	qw422016.N().S(thisForm.Action)
	//line views/templates/json/anyForm.qtpl:61
	qw422016.N().S(`","id" : "`)
	//line views/templates/json/anyForm.qtpl:62
	qw422016.N().S(thisForm.IdCSS)
	//line views/templates/json/anyForm.qtpl:62
	qw422016.N().S(`","name": "`)
	//line views/templates/json/anyForm.qtpl:63
	qw422016.N().S(thisForm.Name)
	//line views/templates/json/anyForm.qtpl:63
	qw422016.N().S(`"`)
	//line views/templates/json/anyForm.qtpl:64
	for name, event := range thisForm.Events {
		//line views/templates/json/anyForm.qtpl:64
		qw422016.N().S(`,"`)
		//line views/templates/json/anyForm.qtpl:65
		qw422016.E().S(name)
		//line views/templates/json/anyForm.qtpl:65
		qw422016.N().S(`": "`)
		//line views/templates/json/anyForm.qtpl:65
		qw422016.N().S(event)
		//line views/templates/json/anyForm.qtpl:65
		qw422016.N().S(`"`)
		//line views/templates/json/anyForm.qtpl:66
	}
	//line views/templates/json/anyForm.qtpl:66
	qw422016.N().S(`}`)
	//line views/templates/json/anyForm.qtpl:68
	for name, val := range AddJson {
		//line views/templates/json/anyForm.qtpl:68
		qw422016.N().S(`,"`)
		//line views/templates/json/anyForm.qtpl:69
		qw422016.E().S(name)
		//line views/templates/json/anyForm.qtpl:69
		qw422016.N().S(`":`)
		//line views/templates/json/anyForm.qtpl:69
		qw422016.N().S(val)
		//line views/templates/json/anyForm.qtpl:70
	}
	//line views/templates/json/anyForm.qtpl:70
	qw422016.N().S(`}`)
//line views/templates/json/anyForm.qtpl:73
}

//line views/templates/json/anyForm.qtpl:73
func (thisForm *FormStructure) WriteJSONAnyForm(qq422016 qtio422016.Writer, ns *forms.FieldsTable, AddJson map[string]string) {
	//line views/templates/json/anyForm.qtpl:73
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/json/anyForm.qtpl:73
	thisForm.StreamJSONAnyForm(qw422016, ns, AddJson)
	//line views/templates/json/anyForm.qtpl:73
	qt422016.ReleaseWriter(qw422016)
//line views/templates/json/anyForm.qtpl:73
}

//line views/templates/json/anyForm.qtpl:73
func (thisForm *FormStructure) JSONAnyForm(ns *forms.FieldsTable, AddJson map[string]string) string {
	//line views/templates/json/anyForm.qtpl:73
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/json/anyForm.qtpl:73
	thisForm.WriteJSONAnyForm(qb422016, ns, AddJson)
	//line views/templates/json/anyForm.qtpl:73
	qs422016 := string(qb422016.B)
	//line views/templates/json/anyForm.qtpl:73
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/json/anyForm.qtpl:73
	return qs422016
//line views/templates/json/anyForm.qtpl:73
}
