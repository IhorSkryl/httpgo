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

//line views/templates/forms/json.qtpl:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/templates/forms/json.qtpl:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/forms/json.qtpl:12
type FormField struct {
	Title, Action, Method, Description string
}

// json for front forms https://storybook.pm-db.net/?path=/story/form-types--page

//line views/templates/forms/json.qtpl:19
func (f *FormField) StreamFormJSON(qw422016 *qt422016.Writer, blocks ...BlockColumns) {
//line views/templates/forms/json.qtpl:19
	qw422016.N().S(`{"title" : "`)
//line views/templates/forms/json.qtpl:21
	qw422016.N().S(f.Title)
//line views/templates/forms/json.qtpl:21
	qw422016.N().S(`","action": "`)
//line views/templates/forms/json.qtpl:22
	qw422016.N().S(f.Action)
//line views/templates/forms/json.qtpl:22
	qw422016.N().S(`","description": "`)
//line views/templates/forms/json.qtpl:23
	qw422016.N().S(f.Description)
//line views/templates/forms/json.qtpl:23
	qw422016.N().S(`","method": "`)
//line views/templates/forms/json.qtpl:24
	qw422016.N().S(f.Method)
//line views/templates/forms/json.qtpl:24
	qw422016.N().S(`","blocks": [`)
//line views/templates/forms/json.qtpl:26
	for i, block := range blocks {
//line views/templates/forms/json.qtpl:28
		if i > 0 {
//line views/templates/forms/json.qtpl:28
			qw422016.N().S(`,`)
//line views/templates/forms/json.qtpl:30
		}
//line views/templates/forms/json.qtpl:30
		qw422016.N().S(`{"id": "`)
//line views/templates/forms/json.qtpl:32
		qw422016.N().D(block.Id)
//line views/templates/forms/json.qtpl:32
		qw422016.N().S(`","title": "`)
//line views/templates/forms/json.qtpl:33
		qw422016.N().S(block.Title)
//line views/templates/forms/json.qtpl:33
		qw422016.N().S(`","description": "`)
//line views/templates/forms/json.qtpl:34
		qw422016.N().S(block.Description)
//line views/templates/forms/json.qtpl:34
		qw422016.N().S(`","fields": [`)
//line views/templates/forms/json.qtpl:36
		for j, col := range block.Columns {
//line views/templates/forms/json.qtpl:38
			if j > 0 {
//line views/templates/forms/json.qtpl:38
				qw422016.N().S(`,`)
//line views/templates/forms/json.qtpl:40
			}
//line views/templates/forms/json.qtpl:40
			qw422016.N().S(`{`)
//line views/templates/forms/json.qtpl:43
			values := col.GetValues()

//line views/templates/forms/json.qtpl:44
			qw422016.N().S(`"name": "`)
//line views/templates/forms/json.qtpl:45
			qw422016.E().S(col.InputName(-1))
//line views/templates/forms/json.qtpl:45
			qw422016.N().S(`","required":`)
//line views/templates/forms/json.qtpl:46
			qw422016.E().V(col.Required())
//line views/templates/forms/json.qtpl:46
			qw422016.N().S(`,"type": "`)
//line views/templates/forms/json.qtpl:47
			col.StreamInputType(qw422016)
//line views/templates/forms/json.qtpl:47
			qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:48
			if col.CharacterMaximumLength() > 0 {
//line views/templates/forms/json.qtpl:48
				qw422016.N().S(`, "maxLength":`)
//line views/templates/forms/json.qtpl:49
				qw422016.N().D(col.CharacterMaximumLength())
//line views/templates/forms/json.qtpl:50
			}
//line views/templates/forms/json.qtpl:51
			switch len(values) {
//line views/templates/forms/json.qtpl:52
			case 0:
//line views/templates/forms/json.qtpl:53
			case 1:
//line views/templates/forms/json.qtpl:54
				if values[0] != nil {
//line views/templates/forms/json.qtpl:54
					qw422016.N().S(`, "value": "`)
//line views/templates/forms/json.qtpl:55
					qw422016.E().V(values[0])
//line views/templates/forms/json.qtpl:55
					qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:56
				}
//line views/templates/forms/json.qtpl:57
			default:
//line views/templates/forms/json.qtpl:57
				qw422016.N().S(`, "value":`)
//line views/templates/forms/json.qtpl:58
				json.StreamSimpleDimension(qw422016, values)
//line views/templates/forms/json.qtpl:59
			}
//line views/templates/forms/json.qtpl:59
			qw422016.N().S(`, "title": "`)
//line views/templates/forms/json.qtpl:61
			qw422016.E().S(col.Label())
//line views/templates/forms/json.qtpl:61
			qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:63
			if p := col.Pattern(); p > "" {
//line views/templates/forms/json.qtpl:63
				qw422016.N().S(`, "pattern":"`)
//line views/templates/forms/json.qtpl:64
				qw422016.E().J(p)
//line views/templates/forms/json.qtpl:64
				qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:65
			}
//line views/templates/forms/json.qtpl:67
			if col.IsReadOnly {
//line views/templates/forms/json.qtpl:67
				qw422016.N().S(`, "readonly": true`)
//line views/templates/forms/json.qtpl:69
			}
//line views/templates/forms/json.qtpl:69
			qw422016.N().S(`}`)
//line views/templates/forms/json.qtpl:71
		}
//line views/templates/forms/json.qtpl:72
		for _, button := range block.Buttons {
//line views/templates/forms/json.qtpl:72
			qw422016.N().S(`,{"type":  "button","title": "`)
//line views/templates/forms/json.qtpl:76
			qw422016.E().S(button.Title)
//line views/templates/forms/json.qtpl:76
			qw422016.N().S(`","position":`)
//line views/templates/forms/json.qtpl:77
			if button.Position {
//line views/templates/forms/json.qtpl:77
				qw422016.N().S(`"right"`)
//line views/templates/forms/json.qtpl:77
			} else {
//line views/templates/forms/json.qtpl:77
				qw422016.N().S(`"left"`)
//line views/templates/forms/json.qtpl:77
			}
//line views/templates/forms/json.qtpl:77
			qw422016.N().S(`}`)
//line views/templates/forms/json.qtpl:79
		}
//line views/templates/forms/json.qtpl:79
		qw422016.N().S(`]}`)
//line views/templates/forms/json.qtpl:82
	}
//line views/templates/forms/json.qtpl:82
	qw422016.N().S(`]}`)
//line views/templates/forms/json.qtpl:85
}

//line views/templates/forms/json.qtpl:85
func (f *FormField) WriteFormJSON(qq422016 qtio422016.Writer, blocks ...BlockColumns) {
//line views/templates/forms/json.qtpl:85
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/forms/json.qtpl:85
	f.StreamFormJSON(qw422016, blocks...)
//line views/templates/forms/json.qtpl:85
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/json.qtpl:85
}

//line views/templates/forms/json.qtpl:85
func (f *FormField) FormJSON(blocks ...BlockColumns) string {
//line views/templates/forms/json.qtpl:85
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/forms/json.qtpl:85
	f.WriteFormJSON(qb422016, blocks...)
//line views/templates/forms/json.qtpl:85
	qs422016 := string(qb422016.B)
//line views/templates/forms/json.qtpl:85
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/forms/json.qtpl:85
	return qs422016
//line views/templates/forms/json.qtpl:85
}

//line views/templates/forms/json.qtpl:88
func (col *ColumnDecor) StreamRenderInputs(qw422016 *qt422016.Writer) {
//line views/templates/forms/json.qtpl:89
	for i, val := range col.GetValues() {
//line views/templates/forms/json.qtpl:89
		qw422016.N().S(`<input type="`)
//line views/templates/forms/json.qtpl:90
		col.StreamInputType(qw422016)
//line views/templates/forms/json.qtpl:90
		qw422016.N().S(`" name="`)
//line views/templates/forms/json.qtpl:90
		qw422016.E().S(col.InputName(i))
//line views/templates/forms/json.qtpl:90
		qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:91
		if col.CharacterMaximumLength() > 0 {
//line views/templates/forms/json.qtpl:91
			qw422016.N().S(`max = "`)
//line views/templates/forms/json.qtpl:91
			qw422016.N().D(col.CharacterMaximumLength())
//line views/templates/forms/json.qtpl:91
			qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:91
		}
//line views/templates/forms/json.qtpl:92
		if val != nil {
//line views/templates/forms/json.qtpl:92
			qw422016.N().S(`value = "`)
//line views/templates/forms/json.qtpl:92
			qw422016.E().V(val)
//line views/templates/forms/json.qtpl:92
			qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:92
		}
//line views/templates/forms/json.qtpl:93
		if col.Required() {
//line views/templates/forms/json.qtpl:93
			qw422016.N().S(`required="true"`)
//line views/templates/forms/json.qtpl:93
		}
//line views/templates/forms/json.qtpl:94
		if col.Pattern() > "" {
//line views/templates/forms/json.qtpl:94
			qw422016.N().S(`pattern="`)
//line views/templates/forms/json.qtpl:94
			qw422016.E().S(col.Pattern())
//line views/templates/forms/json.qtpl:94
			qw422016.N().S(`"  onkeyup="return validatePattern(this);"`)
//line views/templates/forms/json.qtpl:94
		}
//line views/templates/forms/json.qtpl:95
		if col.Placeholder() > "" {
//line views/templates/forms/json.qtpl:95
			qw422016.N().S(`placeholder="`)
//line views/templates/forms/json.qtpl:95
			qw422016.E().S(col.Placeholder())
//line views/templates/forms/json.qtpl:95
			qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:95
		}
//line views/templates/forms/json.qtpl:96
		if col.IsHidden {
//line views/templates/forms/json.qtpl:96
			qw422016.N().S(`hidden`)
//line views/templates/forms/json.qtpl:96
		}
//line views/templates/forms/json.qtpl:97
		if col.IsReadOnly {
//line views/templates/forms/json.qtpl:97
			qw422016.N().S(`readonly disabled`)
//line views/templates/forms/json.qtpl:97
		}
//line views/templates/forms/json.qtpl:97
		qw422016.N().S(`>`)
//line views/templates/forms/json.qtpl:99
	}
//line views/templates/forms/json.qtpl:100
	if col.isSlice {
//line views/templates/forms/json.qtpl:100
		qw422016.N().S(`<button class="main-btn" type="button"onclick="this.parentNode.insertBefore(this.previousElementSibling.cloneNode(), this); return false;">+</button>`)
//line views/templates/forms/json.qtpl:104
	}
//line views/templates/forms/json.qtpl:105
}

//line views/templates/forms/json.qtpl:105
func (col *ColumnDecor) WriteRenderInputs(qq422016 qtio422016.Writer) {
//line views/templates/forms/json.qtpl:105
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/forms/json.qtpl:105
	col.StreamRenderInputs(qw422016)
//line views/templates/forms/json.qtpl:105
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/json.qtpl:105
}

//line views/templates/forms/json.qtpl:105
func (col *ColumnDecor) RenderInputs() string {
//line views/templates/forms/json.qtpl:105
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/forms/json.qtpl:105
	col.WriteRenderInputs(qb422016)
//line views/templates/forms/json.qtpl:105
	qs422016 := string(qb422016.B)
//line views/templates/forms/json.qtpl:105
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/forms/json.qtpl:105
	return qs422016
//line views/templates/forms/json.qtpl:105
}

//line views/templates/forms/json.qtpl:107
func (col *ColumnDecor) StreamInputType(qw422016 *qt422016.Writer) {
//line views/templates/forms/json.qtpl:108
	if col.IsHidden {
//line views/templates/forms/json.qtpl:108
		qw422016.N().S(`hidden`)
//line views/templates/forms/json.qtpl:110
	} else {
//line views/templates/forms/json.qtpl:111
		switch {
//line views/templates/forms/json.qtpl:112
		case typesExt.IsNumeric(col.BasicTypeInfo()):
//line views/templates/forms/json.qtpl:112
			qw422016.N().S(`number`)
//line views/templates/forms/json.qtpl:114
		case col.BasicType() == types.Bool:
//line views/templates/forms/json.qtpl:114
			qw422016.N().S(`switch`)
//line views/templates/forms/json.qtpl:116
		default:
//line views/templates/forms/json.qtpl:117
			switch col.Type() {
//line views/templates/forms/json.qtpl:118
			case "date", "_date":
//line views/templates/forms/json.qtpl:118
				qw422016.N().S(`date`)
//line views/templates/forms/json.qtpl:120
			case "datetime", "timestampt", "timestamptz", "time", "_timestampt", "_timestamptz", "_time":
//line views/templates/forms/json.qtpl:120
				qw422016.N().S(`datetime`)
//line views/templates/forms/json.qtpl:122
			case "email":
//line views/templates/forms/json.qtpl:122
				qw422016.N().S(`email`)
//line views/templates/forms/json.qtpl:123
			case "tel":
//line views/templates/forms/json.qtpl:123
				qw422016.N().S(`tel`)
//line views/templates/forms/json.qtpl:124
			case "password":
//line views/templates/forms/json.qtpl:124
				qw422016.N().S(`password`)
//line views/templates/forms/json.qtpl:125
			default:
//line views/templates/forms/json.qtpl:125
				qw422016.N().S(`text`)
//line views/templates/forms/json.qtpl:126
			}
//line views/templates/forms/json.qtpl:127
		}
//line views/templates/forms/json.qtpl:128
	}
//line views/templates/forms/json.qtpl:129
}

//line views/templates/forms/json.qtpl:129
func (col *ColumnDecor) WriteInputType(qq422016 qtio422016.Writer) {
//line views/templates/forms/json.qtpl:129
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/forms/json.qtpl:129
	col.StreamInputType(qw422016)
//line views/templates/forms/json.qtpl:129
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/json.qtpl:129
}

//line views/templates/forms/json.qtpl:129
func (col *ColumnDecor) InputType() string {
//line views/templates/forms/json.qtpl:129
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/forms/json.qtpl:129
	col.WriteInputType(qb422016)
//line views/templates/forms/json.qtpl:129
	qs422016 := string(qb422016.B)
//line views/templates/forms/json.qtpl:129
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/forms/json.qtpl:129
	return qs422016
//line views/templates/forms/json.qtpl:129
}

//line views/templates/forms/json.qtpl:131
func (f *FormField) StreamFormHTML(qw422016 *qt422016.Writer, columns []ColumnDecor) {
//line views/templates/forms/json.qtpl:131
	qw422016.N().S(`<form id="`)
//line views/templates/forms/json.qtpl:132
	qw422016.E().S(f.Title)
//line views/templates/forms/json.qtpl:132
	qw422016.N().S(`form" name="`)
//line views/templates/forms/json.qtpl:132
	qw422016.E().S(f.Title)
//line views/templates/forms/json.qtpl:132
	qw422016.N().S(`" role='form' class="form-horizontal row-fluid" target="content"action="`)
//line views/templates/forms/json.qtpl:133
	qw422016.E().S(f.Action)
//line views/templates/forms/json.qtpl:133
	qw422016.N().S(`" method="`)
//line views/templates/forms/json.qtpl:133
	qw422016.N().S(f.Method)
//line views/templates/forms/json.qtpl:133
	qw422016.N().S(`" enctype="multipart/form-data"onsubmit="return saveForm(this, afterSaveAnyForm);"  caption="`)
//line views/templates/forms/json.qtpl:134
	qw422016.E().S(f.Title)
//line views/templates/forms/json.qtpl:134
	qw422016.N().S(`" >`)
//line views/templates/forms/json.qtpl:136
	for i, col := range columns {
//line views/templates/forms/json.qtpl:136
		qw422016.N().S(`<div id="divField`)
//line views/templates/forms/json.qtpl:137
		qw422016.N().D(i)
//line views/templates/forms/json.qtpl:137
		qw422016.N().S(`" class="input-wrap"`)
//line views/templates/forms/json.qtpl:138
		if col.IsHidden {
//line views/templates/forms/json.qtpl:138
			qw422016.N().S(`style="display:none"`)
//line views/templates/forms/json.qtpl:138
		}
//line views/templates/forms/json.qtpl:138
		qw422016.N().S(`<label class="input-label" for="`)
//line views/templates/forms/json.qtpl:139
		qw422016.E().S(col.Name())
//line views/templates/forms/json.qtpl:139
		qw422016.N().S(`">`)
//line views/templates/forms/json.qtpl:139
		qw422016.E().S(col.Label())
//line views/templates/forms/json.qtpl:139
		qw422016.N().S(`</label>`)
//line views/templates/forms/json.qtpl:140
		col.StreamRenderInputs(qw422016)
//line views/templates/forms/json.qtpl:140
		qw422016.N().S(`</div>`)
//line views/templates/forms/json.qtpl:142
	}
//line views/templates/forms/json.qtpl:142
	qw422016.N().S(`<div class="form-actions"><button class="main-btn" type="submit">Сохранить</button></div></form><script>`)
//line views/templates/forms/json.qtpl:149
	qw422016.N().S(` `)
//line views/templates/forms/json.qtpl:150
	qw422016.N().S(`
function validatePattern(thisElem) {
    var re = thisElem.pattern,
        result = true;

    if (re == "") {
        return true;
    }

    try {

        re = new RegExp(re);
        result = re.test(thisElem.value);
        if(result){
            thisElem.style.borderColor = 'green';
        } else {
            thisElem.style.borderColor = 'red';
        }

    } catch (e) {
        console.log(e)
    }

    return result;
}
</script>
`)
//line views/templates/forms/json.qtpl:176
}

//line views/templates/forms/json.qtpl:176
func (f *FormField) WriteFormHTML(qq422016 qtio422016.Writer, columns []ColumnDecor) {
//line views/templates/forms/json.qtpl:176
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/forms/json.qtpl:176
	f.StreamFormHTML(qw422016, columns)
//line views/templates/forms/json.qtpl:176
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/json.qtpl:176
}

//line views/templates/forms/json.qtpl:176
func (f *FormField) FormHTML(columns []ColumnDecor) string {
//line views/templates/forms/json.qtpl:176
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/forms/json.qtpl:176
	f.WriteFormHTML(qb422016, columns)
//line views/templates/forms/json.qtpl:176
	qs422016 := string(qb422016.B)
//line views/templates/forms/json.qtpl:176
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/forms/json.qtpl:176
	return qs422016
//line views/templates/forms/json.qtpl:176
}
