// Code generated by qtc from "json.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// All the text outside function templates is treated as comments,
// i.e. it is just ignored by quicktemplate compiler (`qtc`). It is for humans.
//

//line views/templates/forms/json.qtpl:4
package forms

//line views/templates/forms/json.qtpl:4
import (
	"github.com/ruslanBik4/httpgo/dbEngine"
)

// json for front forms.

//line views/templates/forms/json.qtpl:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/templates/forms/json.qtpl:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/forms/json.qtpl:12
func StreamFormJSON(qw422016 *qt422016.Writer, title, action, method string, columns []dbEngine.Column, values ...interface{}) {
//line views/templates/forms/json.qtpl:12
	qw422016.N().S(`{"title" : '`)
//line views/templates/forms/json.qtpl:14
	qw422016.N().S(title)
//line views/templates/forms/json.qtpl:14
	qw422016.N().S(`',"action": '`)
//line views/templates/forms/json.qtpl:15
	qw422016.N().S(action)
//line views/templates/forms/json.qtpl:15
	qw422016.N().S(`',"method": '`)
//line views/templates/forms/json.qtpl:16
	qw422016.N().S(method)
//line views/templates/forms/json.qtpl:16
	qw422016.N().S(`',"fields": [`)
//line views/templates/forms/json.qtpl:18
	for i, col := range columns {
//line views/templates/forms/json.qtpl:20
		if i > 0 {
//line views/templates/forms/json.qtpl:20
			qw422016.N().S(`,`)
//line views/templates/forms/json.qtpl:22
		}
//line views/templates/forms/json.qtpl:22
		qw422016.N().S(`{"name": "`)
//line views/templates/forms/json.qtpl:24
		qw422016.E().S(col.Name())
//line views/templates/forms/json.qtpl:24
		qw422016.N().S(`","required":`)
//line views/templates/forms/json.qtpl:25
		qw422016.E().V(col.Required())
//line views/templates/forms/json.qtpl:25
		qw422016.N().S(`,"type":`)
//line views/templates/forms/json.qtpl:26
		StreamConvertType(qw422016, col.Type())
//line views/templates/forms/json.qtpl:27
		if col.CharacterMaximumLength() > 0 {
//line views/templates/forms/json.qtpl:27
			qw422016.N().S(`, "maxLength":`)
//line views/templates/forms/json.qtpl:28
			qw422016.N().D(col.CharacterMaximumLength())
//line views/templates/forms/json.qtpl:29
		}
//line views/templates/forms/json.qtpl:30
		if i < len(values) && values[i] != nil {
//line views/templates/forms/json.qtpl:30
			qw422016.N().S(`, "value": "`)
//line views/templates/forms/json.qtpl:31
			qw422016.E().V(values[i])
//line views/templates/forms/json.qtpl:31
			qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:32
		}
//line views/templates/forms/json.qtpl:34
		if col.Comment() > "" {
//line views/templates/forms/json.qtpl:34
			qw422016.N().S(`, "title": "`)
//line views/templates/forms/json.qtpl:35
			qw422016.E().S(col.Comment())
//line views/templates/forms/json.qtpl:35
			qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:36
		}
//line views/templates/forms/json.qtpl:36
		qw422016.N().S(`}`)
//line views/templates/forms/json.qtpl:38
	}
//line views/templates/forms/json.qtpl:38
	qw422016.N().S(`]}`)
//line views/templates/forms/json.qtpl:41
}

//line views/templates/forms/json.qtpl:41
func WriteFormJSON(qq422016 qtio422016.Writer, title, action, method string, columns []dbEngine.Column, values ...interface{}) {
//line views/templates/forms/json.qtpl:41
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/forms/json.qtpl:41
	StreamFormJSON(qw422016, title, action, method, columns, values...)
//line views/templates/forms/json.qtpl:41
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/json.qtpl:41
}

//line views/templates/forms/json.qtpl:41
func FormJSON(title, action, method string, columns []dbEngine.Column, values ...interface{}) string {
//line views/templates/forms/json.qtpl:41
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/forms/json.qtpl:41
	WriteFormJSON(qb422016, title, action, method, columns, values...)
//line views/templates/forms/json.qtpl:41
	qs422016 := string(qb422016.B)
//line views/templates/forms/json.qtpl:41
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/forms/json.qtpl:41
	return qs422016
//line views/templates/forms/json.qtpl:41
}

//line views/templates/forms/json.qtpl:43
func StreamFormHTML(qw422016 *qt422016.Writer, title, action, method string, columns []dbEngine.Column, values ...interface{}) {
//line views/templates/forms/json.qtpl:43
	qw422016.N().S(`<form id="`)
//line views/templates/forms/json.qtpl:44
	qw422016.E().S(title)
//line views/templates/forms/json.qtpl:44
	qw422016.N().S(`form" name="`)
//line views/templates/forms/json.qtpl:44
	qw422016.E().S(title)
//line views/templates/forms/json.qtpl:44
	qw422016.N().S(`" role='form' class="form-horizontal row-fluid" target="content"action="`)
//line views/templates/forms/json.qtpl:45
	qw422016.E().S(action)
//line views/templates/forms/json.qtpl:45
	qw422016.N().S(`" method="`)
//line views/templates/forms/json.qtpl:45
	qw422016.N().S(method)
//line views/templates/forms/json.qtpl:45
	qw422016.N().S(`"onsubmit="return saveForm(this, afterSaveAnyForm);"  caption="`)
//line views/templates/forms/json.qtpl:46
	qw422016.E().S(title)
//line views/templates/forms/json.qtpl:46
	qw422016.N().S(`" >`)
//line views/templates/forms/json.qtpl:48
	for i, col := range columns {
//line views/templates/forms/json.qtpl:48
		qw422016.N().S(`<label class="input-label" for="`)
//line views/templates/forms/json.qtpl:49
		qw422016.E().S(col.Name())
//line views/templates/forms/json.qtpl:49
		qw422016.N().S(`">`)
//line views/templates/forms/json.qtpl:49
		qw422016.E().S(col.Name())
//line views/templates/forms/json.qtpl:49
		qw422016.E().S(col.Comment())
//line views/templates/forms/json.qtpl:49
		qw422016.N().S(`</label><input type="`)
//line views/templates/forms/json.qtpl:50
		StreamConvertType(qw422016, col.Type())
//line views/templates/forms/json.qtpl:50
		qw422016.N().S(`" name="`)
//line views/templates/forms/json.qtpl:50
		qw422016.E().S(col.Name())
//line views/templates/forms/json.qtpl:50
		qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:51
		if col.Required() {
//line views/templates/forms/json.qtpl:51
			qw422016.N().S(`required`)
//line views/templates/forms/json.qtpl:51
		}
//line views/templates/forms/json.qtpl:52
		if col.CharacterMaximumLength() > 0 {
//line views/templates/forms/json.qtpl:52
			qw422016.N().S(`max = "`)
//line views/templates/forms/json.qtpl:53
			qw422016.N().D(col.CharacterMaximumLength())
//line views/templates/forms/json.qtpl:53
			qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:54
		}
//line views/templates/forms/json.qtpl:55
		if i < len(values) && values[i] != nil {
//line views/templates/forms/json.qtpl:55
			qw422016.N().S(`value = "`)
//line views/templates/forms/json.qtpl:56
			qw422016.E().V(values[i])
//line views/templates/forms/json.qtpl:56
			qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:57
		}
//line views/templates/forms/json.qtpl:57
		qw422016.N().S(`>`)
//line views/templates/forms/json.qtpl:59
	}
//line views/templates/forms/json.qtpl:59
	qw422016.N().S(`<div class="form-actions"><button class="main-btn" type="submit">Сохранить</button></div></form>`)
//line views/templates/forms/json.qtpl:65
}

//line views/templates/forms/json.qtpl:65
func WriteFormHTML(qq422016 qtio422016.Writer, title, action, method string, columns []dbEngine.Column, values ...interface{}) {
//line views/templates/forms/json.qtpl:65
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/forms/json.qtpl:65
	StreamFormHTML(qw422016, title, action, method, columns, values...)
//line views/templates/forms/json.qtpl:65
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/json.qtpl:65
}

//line views/templates/forms/json.qtpl:65
func FormHTML(title, action, method string, columns []dbEngine.Column, values ...interface{}) string {
//line views/templates/forms/json.qtpl:65
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/forms/json.qtpl:65
	WriteFormHTML(qb422016, title, action, method, columns, values...)
//line views/templates/forms/json.qtpl:65
	qs422016 := string(qb422016.B)
//line views/templates/forms/json.qtpl:65
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/forms/json.qtpl:65
	return qs422016
//line views/templates/forms/json.qtpl:65
}

//line views/templates/forms/json.qtpl:67
func StreamConvertType(qw422016 *qt422016.Writer, t string) {
//line views/templates/forms/json.qtpl:68
	switch t {
//line views/templates/forms/json.qtpl:69
	case "int", "int4", "_int4", "int8", "_int8", "float4", "_float4", "float8", "_float8", "numeric", "decimal":
//line views/templates/forms/json.qtpl:69
		qw422016.N().S(`"number"`)
//line views/templates/forms/json.qtpl:71
	case "date", "datetime", "timestampt", "timestamptz", "time", "_date", "_timestampt", "_timestamptz", "_time":
//line views/templates/forms/json.qtpl:71
		qw422016.N().S(`"datetime"`)
//line views/templates/forms/json.qtpl:73
	case "boolean", "bool":
//line views/templates/forms/json.qtpl:73
		qw422016.N().S(`"checkbox"`)
//line views/templates/forms/json.qtpl:74
	case "tel":
//line views/templates/forms/json.qtpl:74
		qw422016.N().S(`"tel"`)
//line views/templates/forms/json.qtpl:75
	case "password":
//line views/templates/forms/json.qtpl:75
		qw422016.N().S(`"password"`)
//line views/templates/forms/json.qtpl:76
	default:
//line views/templates/forms/json.qtpl:76
		qw422016.N().S(`"text"`)
//line views/templates/forms/json.qtpl:77
	}
//line views/templates/forms/json.qtpl:78
}

//line views/templates/forms/json.qtpl:78
func WriteConvertType(qq422016 qtio422016.Writer, t string) {
//line views/templates/forms/json.qtpl:78
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/forms/json.qtpl:78
	StreamConvertType(qw422016, t)
//line views/templates/forms/json.qtpl:78
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/json.qtpl:78
}

//line views/templates/forms/json.qtpl:78
func ConvertType(t string) string {
//line views/templates/forms/json.qtpl:78
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/forms/json.qtpl:78
	WriteConvertType(qb422016, t)
//line views/templates/forms/json.qtpl:78
	qs422016 := string(qb422016.B)
//line views/templates/forms/json.qtpl:78
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/forms/json.qtpl:78
	return qs422016
//line views/templates/forms/json.qtpl:78
}
