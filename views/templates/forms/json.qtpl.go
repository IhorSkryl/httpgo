// Code generated by qtc from "json.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// All the text outside function templates is treated as comments,
// i.e. it is just ignored by quicktemplate compiler (`qtc`). It is for humans.
//

//line views/templates/forms/json.qtpl:4
package forms

//line views/templates/forms/json.qtpl:4
import (
	"go/types"

	"github.com/ruslanBik4/httpgo/typesExt"
	"github.com/ruslanBik4/httpgo/views/templates/json"
)

// json for front forms.

//line views/templates/forms/json.qtpl:15
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/templates/forms/json.qtpl:15
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/forms/json.qtpl:15
func StreamFormJSON(qw422016 *qt422016.Writer, title, action, method string, columns []ColumnDecor) {
//line views/templates/forms/json.qtpl:15
	qw422016.N().S(`{"title" : '`)
//line views/templates/forms/json.qtpl:17
	qw422016.N().S(title)
//line views/templates/forms/json.qtpl:17
	qw422016.N().S(`',"action": '`)
//line views/templates/forms/json.qtpl:18
	qw422016.N().S(action)
//line views/templates/forms/json.qtpl:18
	qw422016.N().S(`',"method": '`)
//line views/templates/forms/json.qtpl:19
	qw422016.N().S(method)
//line views/templates/forms/json.qtpl:19
	qw422016.N().S(`',"fields": [`)
//line views/templates/forms/json.qtpl:21
	for i, col := range columns {
//line views/templates/forms/json.qtpl:23
		if i > 0 {
//line views/templates/forms/json.qtpl:23
			qw422016.N().S(`,`)
//line views/templates/forms/json.qtpl:25
		}
//line views/templates/forms/json.qtpl:25
		qw422016.N().S(`{`)
//line views/templates/forms/json.qtpl:28
		values := col.GetValues()

//line views/templates/forms/json.qtpl:29
		qw422016.N().S(`"name": "`)
//line views/templates/forms/json.qtpl:30
		qw422016.E().S(col.InputName(-1))
//line views/templates/forms/json.qtpl:30
		qw422016.N().S(`","required":`)
//line views/templates/forms/json.qtpl:31
		qw422016.E().V(col.Required())
//line views/templates/forms/json.qtpl:31
		qw422016.N().S(`,"type": "`)
//line views/templates/forms/json.qtpl:32
		col.StreamInputType(qw422016)
//line views/templates/forms/json.qtpl:32
		qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:33
		if col.CharacterMaximumLength() > 0 {
//line views/templates/forms/json.qtpl:33
			qw422016.N().S(`, "maxLength":`)
//line views/templates/forms/json.qtpl:34
			qw422016.N().D(col.CharacterMaximumLength())
//line views/templates/forms/json.qtpl:35
		}
//line views/templates/forms/json.qtpl:36
		switch len(values) {
//line views/templates/forms/json.qtpl:37
		case 0:
//line views/templates/forms/json.qtpl:38
		case 1:
//line views/templates/forms/json.qtpl:38
			qw422016.N().S(`, "value": "`)
//line views/templates/forms/json.qtpl:39
			qw422016.E().V(values)
//line views/templates/forms/json.qtpl:39
			qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:40
		default:
//line views/templates/forms/json.qtpl:40
			qw422016.N().S(`, "value":`)
//line views/templates/forms/json.qtpl:41
			json.StreamSimpleDimension(qw422016, values)
//line views/templates/forms/json.qtpl:42
		}
//line views/templates/forms/json.qtpl:44
		if col.Comment() > "" {
//line views/templates/forms/json.qtpl:44
			qw422016.N().S(`, "title": "`)
//line views/templates/forms/json.qtpl:45
			qw422016.E().S(col.Comment())
//line views/templates/forms/json.qtpl:45
			qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:46
		}
//line views/templates/forms/json.qtpl:47
		if col.Primary() {
//line views/templates/forms/json.qtpl:47
			qw422016.N().S(`, "hidden": "hidden"`)
//line views/templates/forms/json.qtpl:49
		}
//line views/templates/forms/json.qtpl:49
		qw422016.N().S(`}`)
//line views/templates/forms/json.qtpl:51
	}
//line views/templates/forms/json.qtpl:51
	qw422016.N().S(`]}`)
//line views/templates/forms/json.qtpl:54
}

//line views/templates/forms/json.qtpl:54
func WriteFormJSON(qq422016 qtio422016.Writer, title, action, method string, columns []ColumnDecor) {
//line views/templates/forms/json.qtpl:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/forms/json.qtpl:54
	StreamFormJSON(qw422016, title, action, method, columns)
//line views/templates/forms/json.qtpl:54
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/json.qtpl:54
}

//line views/templates/forms/json.qtpl:54
func FormJSON(title, action, method string, columns []ColumnDecor) string {
//line views/templates/forms/json.qtpl:54
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/forms/json.qtpl:54
	WriteFormJSON(qb422016, title, action, method, columns)
//line views/templates/forms/json.qtpl:54
	qs422016 := string(qb422016.B)
//line views/templates/forms/json.qtpl:54
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/forms/json.qtpl:54
	return qs422016
//line views/templates/forms/json.qtpl:54
}

//line views/templates/forms/json.qtpl:56
func StreamFormHTML(qw422016 *qt422016.Writer, title, action, method string, columns []ColumnDecor) {
//line views/templates/forms/json.qtpl:56
	qw422016.N().S(`<form id="`)
//line views/templates/forms/json.qtpl:57
	qw422016.E().S(title)
//line views/templates/forms/json.qtpl:57
	qw422016.N().S(`form" name="`)
//line views/templates/forms/json.qtpl:57
	qw422016.E().S(title)
//line views/templates/forms/json.qtpl:57
	qw422016.N().S(`" role='form' class="form-horizontal row-fluid" target="content"action="`)
//line views/templates/forms/json.qtpl:58
	qw422016.E().S(action)
//line views/templates/forms/json.qtpl:58
	qw422016.N().S(`" method="`)
//line views/templates/forms/json.qtpl:58
	qw422016.N().S(method)
//line views/templates/forms/json.qtpl:58
	qw422016.N().S(`" enctype="multipart/form-data"onsubmit="return saveForm(this, afterSaveAnyForm);"  caption="`)
//line views/templates/forms/json.qtpl:59
	qw422016.E().S(title)
//line views/templates/forms/json.qtpl:59
	qw422016.N().S(`" >`)
//line views/templates/forms/json.qtpl:61
	for i, col := range columns {
//line views/templates/forms/json.qtpl:61
		qw422016.N().S(`<div id="divField`)
//line views/templates/forms/json.qtpl:62
		qw422016.N().D(i)
//line views/templates/forms/json.qtpl:62
		qw422016.N().S(`" class="input-wrap"`)
//line views/templates/forms/json.qtpl:63
		if col.IsHidden {
//line views/templates/forms/json.qtpl:63
			qw422016.N().S(`style="display:none"`)
//line views/templates/forms/json.qtpl:63
		}
//line views/templates/forms/json.qtpl:63
		qw422016.N().S(`<label class="input-label" for="`)
//line views/templates/forms/json.qtpl:64
		qw422016.E().S(col.Name())
//line views/templates/forms/json.qtpl:64
		qw422016.N().S(`">`)
//line views/templates/forms/json.qtpl:64
		qw422016.E().S(col.Label())
//line views/templates/forms/json.qtpl:64
		qw422016.N().S(`</label>`)
//line views/templates/forms/json.qtpl:65
		col.StreamRenderInputs(qw422016)
//line views/templates/forms/json.qtpl:65
		qw422016.N().S(`</div>`)
//line views/templates/forms/json.qtpl:67
	}
//line views/templates/forms/json.qtpl:67
	qw422016.N().S(`<div class="form-actions"><button class="main-btn" type="submit">Сохранить</button></div></form>`)
//line views/templates/forms/json.qtpl:73
}

//line views/templates/forms/json.qtpl:73
func WriteFormHTML(qq422016 qtio422016.Writer, title, action, method string, columns []ColumnDecor) {
//line views/templates/forms/json.qtpl:73
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/forms/json.qtpl:73
	StreamFormHTML(qw422016, title, action, method, columns)
//line views/templates/forms/json.qtpl:73
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/json.qtpl:73
}

//line views/templates/forms/json.qtpl:73
func FormHTML(title, action, method string, columns []ColumnDecor) string {
//line views/templates/forms/json.qtpl:73
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/forms/json.qtpl:73
	WriteFormHTML(qb422016, title, action, method, columns)
//line views/templates/forms/json.qtpl:73
	qs422016 := string(qb422016.B)
//line views/templates/forms/json.qtpl:73
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/forms/json.qtpl:73
	return qs422016
//line views/templates/forms/json.qtpl:73
}

//line views/templates/forms/json.qtpl:75
func (col *ColumnDecor) StreamRenderInputs(qw422016 *qt422016.Writer) {
//line views/templates/forms/json.qtpl:76
	for i, val := range col.GetValues() {
//line views/templates/forms/json.qtpl:76
		qw422016.N().S(`col.RenderInput(val)<input type="`)
//line views/templates/forms/json.qtpl:78
		col.StreamInputType(qw422016)
//line views/templates/forms/json.qtpl:78
		qw422016.N().S(`" name="`)
//line views/templates/forms/json.qtpl:78
		qw422016.E().S(col.InputName(i))
//line views/templates/forms/json.qtpl:78
		qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:79
		if col.CharacterMaximumLength() > 0 {
//line views/templates/forms/json.qtpl:79
			qw422016.N().S(`max = "`)
//line views/templates/forms/json.qtpl:79
			qw422016.N().D(col.CharacterMaximumLength())
//line views/templates/forms/json.qtpl:79
			qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:79
		}
//line views/templates/forms/json.qtpl:80
		if val != nil {
//line views/templates/forms/json.qtpl:80
			qw422016.N().S(`value = "`)
//line views/templates/forms/json.qtpl:80
			qw422016.E().V(val)
//line views/templates/forms/json.qtpl:80
			qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:80
		}
//line views/templates/forms/json.qtpl:81
		if col.Required() {
//line views/templates/forms/json.qtpl:81
			qw422016.N().S(`required="true"`)
//line views/templates/forms/json.qtpl:81
		}
//line views/templates/forms/json.qtpl:82
		if col.IsHidden {
//line views/templates/forms/json.qtpl:82
			qw422016.N().S(`hidden`)
//line views/templates/forms/json.qtpl:82
		}
//line views/templates/forms/json.qtpl:82
		qw422016.N().S(`>`)
//line views/templates/forms/json.qtpl:84
	}
//line views/templates/forms/json.qtpl:85
}

//line views/templates/forms/json.qtpl:85
func (col *ColumnDecor) WriteRenderInputs(qq422016 qtio422016.Writer) {
//line views/templates/forms/json.qtpl:85
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/forms/json.qtpl:85
	col.StreamRenderInputs(qw422016)
//line views/templates/forms/json.qtpl:85
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/json.qtpl:85
}

//line views/templates/forms/json.qtpl:85
func (col *ColumnDecor) RenderInputs() string {
//line views/templates/forms/json.qtpl:85
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/forms/json.qtpl:85
	col.WriteRenderInputs(qb422016)
//line views/templates/forms/json.qtpl:85
	qs422016 := string(qb422016.B)
//line views/templates/forms/json.qtpl:85
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/forms/json.qtpl:85
	return qs422016
//line views/templates/forms/json.qtpl:85
}

//line views/templates/forms/json.qtpl:87
func (col *ColumnDecor) StreamInputType(qw422016 *qt422016.Writer) {
//line views/templates/forms/json.qtpl:88
	switch {
//line views/templates/forms/json.qtpl:89
	case typesExt.IsNumeric(col.BasicTypeInfo()):
//line views/templates/forms/json.qtpl:89
		qw422016.N().S(`number`)
//line views/templates/forms/json.qtpl:91
	case col.BasicType() == types.Bool:
//line views/templates/forms/json.qtpl:91
		qw422016.N().S(`checkbox`)
//line views/templates/forms/json.qtpl:93
	default:
//line views/templates/forms/json.qtpl:94
		switch col.Type() {
//line views/templates/forms/json.qtpl:95
		case "date", "datetime", "timestampt", "timestamptz", "time", "_date", "_timestampt", "_timestamptz", "_time":
//line views/templates/forms/json.qtpl:95
			qw422016.N().S(`datetime`)
//line views/templates/forms/json.qtpl:97
		case "tel":
//line views/templates/forms/json.qtpl:97
			qw422016.N().S(`tel`)
//line views/templates/forms/json.qtpl:98
		case "password":
//line views/templates/forms/json.qtpl:98
			qw422016.N().S(`password`)
//line views/templates/forms/json.qtpl:99
		default:
//line views/templates/forms/json.qtpl:99
			qw422016.N().S(`text`)
//line views/templates/forms/json.qtpl:100
		}
//line views/templates/forms/json.qtpl:101
	}
//line views/templates/forms/json.qtpl:102
}

//line views/templates/forms/json.qtpl:102
func (col *ColumnDecor) WriteInputType(qq422016 qtio422016.Writer) {
//line views/templates/forms/json.qtpl:102
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/forms/json.qtpl:102
	col.StreamInputType(qw422016)
//line views/templates/forms/json.qtpl:102
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/json.qtpl:102
}

//line views/templates/forms/json.qtpl:102
func (col *ColumnDecor) InputType() string {
//line views/templates/forms/json.qtpl:102
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/forms/json.qtpl:102
	col.WriteInputType(qb422016)
//line views/templates/forms/json.qtpl:102
	qs422016 := string(qb422016.B)
//line views/templates/forms/json.qtpl:102
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/forms/json.qtpl:102
	return qs422016
//line views/templates/forms/json.qtpl:102
}
