// Code generated by qtc from "json.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// All the text outside function templates is treated as comments,
// i.e. it is just ignored by quicktemplate compiler (`qtc`). It is for humans.
//

//line views/templates/forms/json.qtpl:4
package forms

//line views/templates/forms/json.qtpl:4
import (
	"github.com/ruslanBik4/httpgo/views/templates/json"
)

//line views/templates/forms/json.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/templates/forms/json.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/forms/json.qtpl:9
type FormField struct {
	Title, Action, Method, Description string
}

// json for front forms https://storybook.pm-db.net/?path=/story/form-types--page

//line views/templates/forms/json.qtpl:17
func (f *FormField) StreamFormJSON(qw422016 *qt422016.Writer, blocks ...BlockColumns) {
//line views/templates/forms/json.qtpl:17
	qw422016.N().S(`{"title" : "`)
//line views/templates/forms/json.qtpl:19
	qw422016.N().S(f.Title)
//line views/templates/forms/json.qtpl:19
	qw422016.N().S(`","action": "`)
//line views/templates/forms/json.qtpl:20
	qw422016.N().S(f.Action)
//line views/templates/forms/json.qtpl:20
	qw422016.N().S(`","description": "`)
//line views/templates/forms/json.qtpl:21
	qw422016.N().S(f.Description)
//line views/templates/forms/json.qtpl:21
	qw422016.N().S(`","method": "`)
//line views/templates/forms/json.qtpl:22
	qw422016.N().S(f.Method)
//line views/templates/forms/json.qtpl:22
	qw422016.N().S(`","blocks": [`)
//line views/templates/forms/json.qtpl:24
	for i, block := range blocks {
//line views/templates/forms/json.qtpl:26
		if i > 0 {
//line views/templates/forms/json.qtpl:26
			qw422016.N().S(`,`)
//line views/templates/forms/json.qtpl:28
		}
//line views/templates/forms/json.qtpl:28
		qw422016.N().S(`{"id": "`)
//line views/templates/forms/json.qtpl:30
		qw422016.N().D(block.Id)
//line views/templates/forms/json.qtpl:30
		qw422016.N().S(`","title": "`)
//line views/templates/forms/json.qtpl:31
		qw422016.N().S(block.Title)
//line views/templates/forms/json.qtpl:31
		qw422016.N().S(`","description": "`)
//line views/templates/forms/json.qtpl:32
		qw422016.N().S(block.Description)
//line views/templates/forms/json.qtpl:32
		qw422016.N().S(`","fields": [`)
//line views/templates/forms/json.qtpl:34
		for j, col := range block.Columns {
//line views/templates/forms/json.qtpl:36
			if j > 0 {
//line views/templates/forms/json.qtpl:36
				qw422016.N().S(`,`)
//line views/templates/forms/json.qtpl:38
			}
//line views/templates/forms/json.qtpl:38
			qw422016.N().S(`{`)
//line views/templates/forms/json.qtpl:41
			values := col.GetValues()

//line views/templates/forms/json.qtpl:42
			qw422016.N().S(`"name": "`)
//line views/templates/forms/json.qtpl:43
			qw422016.E().S(col.InputName(-1))
//line views/templates/forms/json.qtpl:43
			qw422016.N().S(`","required":`)
//line views/templates/forms/json.qtpl:44
			qw422016.E().V(col.Required())
//line views/templates/forms/json.qtpl:44
			qw422016.N().S(`,"type": "`)
//line views/templates/forms/json.qtpl:45
			col.StreamInputTypeForJSON(qw422016)
//line views/templates/forms/json.qtpl:45
			qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:46
			col.StreamDataForJSON(qw422016)
//line views/templates/forms/json.qtpl:47
			if col.CharacterMaximumLength() > 0 {
//line views/templates/forms/json.qtpl:47
				qw422016.N().S(`, "maxLength":`)
//line views/templates/forms/json.qtpl:48
				qw422016.N().D(col.CharacterMaximumLength())
//line views/templates/forms/json.qtpl:49
			}
//line views/templates/forms/json.qtpl:50
			switch len(values) {
//line views/templates/forms/json.qtpl:51
			case 0:
//line views/templates/forms/json.qtpl:52
			case 1:
//line views/templates/forms/json.qtpl:53
				if values[0] != nil {
//line views/templates/forms/json.qtpl:53
					qw422016.N().S(`, "value": "`)
//line views/templates/forms/json.qtpl:54
					qw422016.E().V(values[0])
//line views/templates/forms/json.qtpl:54
					qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:55
				}
//line views/templates/forms/json.qtpl:56
			default:
//line views/templates/forms/json.qtpl:56
				qw422016.N().S(`, "value":`)
//line views/templates/forms/json.qtpl:57
				json.StreamSimpleDimension(qw422016, values)
//line views/templates/forms/json.qtpl:58
			}
//line views/templates/forms/json.qtpl:58
			qw422016.N().S(`, "title": "`)
//line views/templates/forms/json.qtpl:60
			qw422016.E().S(col.Label)
//line views/templates/forms/json.qtpl:60
			qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:62
			if p := col.Pattern(); p > "" {
//line views/templates/forms/json.qtpl:62
				qw422016.N().S(`, "pattern":"`)
//line views/templates/forms/json.qtpl:63
				qw422016.E().J(p)
//line views/templates/forms/json.qtpl:63
				qw422016.N().S(`"`)
//line views/templates/forms/json.qtpl:64
			}
//line views/templates/forms/json.qtpl:66
			if col.IsReadOnly {
//line views/templates/forms/json.qtpl:66
				qw422016.N().S(`, "readonly": true`)
//line views/templates/forms/json.qtpl:68
			}
//line views/templates/forms/json.qtpl:68
			qw422016.N().S(`}`)
//line views/templates/forms/json.qtpl:70
		}
//line views/templates/forms/json.qtpl:71
		for _, button := range block.Buttons {
//line views/templates/forms/json.qtpl:71
			qw422016.N().S(`,{"type":  "`)
//line views/templates/forms/json.qtpl:74
			qw422016.E().S(button.ButtonType)
//line views/templates/forms/json.qtpl:74
			qw422016.N().S(`","title": "`)
//line views/templates/forms/json.qtpl:75
			qw422016.E().S(button.Title)
//line views/templates/forms/json.qtpl:75
			qw422016.N().S(`","position":`)
//line views/templates/forms/json.qtpl:76
			if button.Position {
//line views/templates/forms/json.qtpl:76
				qw422016.N().S(`"right"`)
//line views/templates/forms/json.qtpl:76
			} else {
//line views/templates/forms/json.qtpl:76
				qw422016.N().S(`"left"`)
//line views/templates/forms/json.qtpl:76
			}
//line views/templates/forms/json.qtpl:76
			qw422016.N().S(`}`)
//line views/templates/forms/json.qtpl:78
		}
//line views/templates/forms/json.qtpl:78
		qw422016.N().S(`]}`)
//line views/templates/forms/json.qtpl:81
	}
//line views/templates/forms/json.qtpl:81
	qw422016.N().S(`]}`)
//line views/templates/forms/json.qtpl:84
}

//line views/templates/forms/json.qtpl:84
func (f *FormField) WriteFormJSON(qq422016 qtio422016.Writer, blocks ...BlockColumns) {
//line views/templates/forms/json.qtpl:84
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/forms/json.qtpl:84
	f.StreamFormJSON(qw422016, blocks...)
//line views/templates/forms/json.qtpl:84
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/json.qtpl:84
}

//line views/templates/forms/json.qtpl:84
func (f *FormField) FormJSON(blocks ...BlockColumns) string {
//line views/templates/forms/json.qtpl:84
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/forms/json.qtpl:84
	f.WriteFormJSON(qb422016, blocks...)
//line views/templates/forms/json.qtpl:84
	qs422016 := string(qb422016.B)
//line views/templates/forms/json.qtpl:84
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/forms/json.qtpl:84
	return qs422016
//line views/templates/forms/json.qtpl:84
}

//line views/templates/forms/json.qtpl:87
func (col *ColumnDecor) StreamDataForJSON(qw422016 *qt422016.Writer) {
//line views/templates/forms/json.qtpl:88
	if col.SelectOptions != nil {
//line views/templates/forms/json.qtpl:88
		qw422016.N().S(`, "data": [`)
//line views/templates/forms/json.qtpl:90
		for title, val := range col.SelectOptions {
//line views/templates/forms/json.qtpl:90
			qw422016.N().S(`{"title": "`)
//line views/templates/forms/json.qtpl:92
			qw422016.E().S(title)
//line views/templates/forms/json.qtpl:92
			qw422016.N().S(`","value": "`)
//line views/templates/forms/json.qtpl:93
			qw422016.E().S(val)
//line views/templates/forms/json.qtpl:93
			qw422016.N().S(`"},`)
//line views/templates/forms/json.qtpl:95
		}
//line views/templates/forms/json.qtpl:95
		qw422016.N().S(`{"title": "add new item","value": "new"}]`)
//line views/templates/forms/json.qtpl:101
	}
//line views/templates/forms/json.qtpl:102
}

//line views/templates/forms/json.qtpl:102
func (col *ColumnDecor) WriteDataForJSON(qq422016 qtio422016.Writer) {
//line views/templates/forms/json.qtpl:102
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/forms/json.qtpl:102
	col.StreamDataForJSON(qw422016)
//line views/templates/forms/json.qtpl:102
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/json.qtpl:102
}

//line views/templates/forms/json.qtpl:102
func (col *ColumnDecor) DataForJSON() string {
//line views/templates/forms/json.qtpl:102
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/forms/json.qtpl:102
	col.WriteDataForJSON(qb422016)
//line views/templates/forms/json.qtpl:102
	qs422016 := string(qb422016.B)
//line views/templates/forms/json.qtpl:102
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/forms/json.qtpl:102
	return qs422016
//line views/templates/forms/json.qtpl:102
}

//line views/templates/forms/json.qtpl:104
func (col *ColumnDecor) StreamInputTypeForJSON(qw422016 *qt422016.Writer) {
//line views/templates/forms/json.qtpl:105
	switch {
//line views/templates/forms/json.qtpl:106
	case col.InputType == "checkbox":
//line views/templates/forms/json.qtpl:106
		qw422016.N().S(`switch`)
//line views/templates/forms/json.qtpl:108
	case col.SelectOptions != nil:
//line views/templates/forms/json.qtpl:108
		qw422016.N().S(`select`)
//line views/templates/forms/json.qtpl:110
	default:
//line views/templates/forms/json.qtpl:111
		qw422016.N().S(col.InputType)
//line views/templates/forms/json.qtpl:112
	}
//line views/templates/forms/json.qtpl:113
}

//line views/templates/forms/json.qtpl:113
func (col *ColumnDecor) WriteInputTypeForJSON(qq422016 qtio422016.Writer) {
//line views/templates/forms/json.qtpl:113
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/forms/json.qtpl:113
	col.StreamInputTypeForJSON(qw422016)
//line views/templates/forms/json.qtpl:113
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/json.qtpl:113
}

//line views/templates/forms/json.qtpl:113
func (col *ColumnDecor) InputTypeForJSON() string {
//line views/templates/forms/json.qtpl:113
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/forms/json.qtpl:113
	col.WriteInputTypeForJSON(qb422016)
//line views/templates/forms/json.qtpl:113
	qs422016 := string(qb422016.B)
//line views/templates/forms/json.qtpl:113
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/forms/json.qtpl:113
	return qs422016
//line views/templates/forms/json.qtpl:113
}

//line views/templates/forms/json.qtpl:116
func (col *ColumnDecor) StreamRenderAttr(qw422016 *qt422016.Writer, i int) {
//line views/templates/forms/json.qtpl:116
	qw422016.N().S(` name="`)
//line views/templates/forms/json.qtpl:117
	qw422016.E().S(col.InputName(i))
//line views/templates/forms/json.qtpl:117
	qw422016.N().S(`" `)
//line views/templates/forms/json.qtpl:118
	if col.Required() {
//line views/templates/forms/json.qtpl:118
		qw422016.N().S(` required="true" `)
//line views/templates/forms/json.qtpl:118
	}
//line views/templates/forms/json.qtpl:118
	qw422016.N().S(` `)
//line views/templates/forms/json.qtpl:119
	if col.IsReadOnly {
//line views/templates/forms/json.qtpl:119
		qw422016.N().S(` readonly disabled `)
//line views/templates/forms/json.qtpl:119
	}
//line views/templates/forms/json.qtpl:119
	qw422016.N().S(` `)
//line views/templates/forms/json.qtpl:121
}

//line views/templates/forms/json.qtpl:121
func (col *ColumnDecor) WriteRenderAttr(qq422016 qtio422016.Writer, i int) {
//line views/templates/forms/json.qtpl:121
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/forms/json.qtpl:121
	col.StreamRenderAttr(qw422016, i)
//line views/templates/forms/json.qtpl:121
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/json.qtpl:121
}

//line views/templates/forms/json.qtpl:121
func (col *ColumnDecor) RenderAttr(i int) string {
//line views/templates/forms/json.qtpl:121
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/forms/json.qtpl:121
	col.WriteRenderAttr(qb422016, i)
//line views/templates/forms/json.qtpl:121
	qs422016 := string(qb422016.B)
//line views/templates/forms/json.qtpl:121
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/forms/json.qtpl:121
	return qs422016
//line views/templates/forms/json.qtpl:121
}

//line views/templates/forms/json.qtpl:123
func (col *ColumnDecor) StreamRenderInputs(qw422016 *qt422016.Writer) {
//line views/templates/forms/json.qtpl:123
	qw422016.N().S(` `)
//line views/templates/forms/json.qtpl:124
	for i, val := range col.GetValues() {
//line views/templates/forms/json.qtpl:124
		qw422016.N().S(` `)
//line views/templates/forms/json.qtpl:125
		if col.SelectOptions != nil {
//line views/templates/forms/json.qtpl:125
			qw422016.N().S(` <select `)
//line views/templates/forms/json.qtpl:126
			col.StreamRenderAttr(qw422016, i)
//line views/templates/forms/json.qtpl:126
			qw422016.N().S(` > `)
//line views/templates/forms/json.qtpl:127
			for title, value := range col.SelectOptions {
//line views/templates/forms/json.qtpl:127
				qw422016.N().S(` <option value="`)
//line views/templates/forms/json.qtpl:128
				qw422016.E().S(value)
//line views/templates/forms/json.qtpl:128
				qw422016.N().S(`" `)
//line views/templates/forms/json.qtpl:128
				if val != nil && val == value {
//line views/templates/forms/json.qtpl:128
					qw422016.N().S(` selected `)
//line views/templates/forms/json.qtpl:128
				}
//line views/templates/forms/json.qtpl:128
				qw422016.N().S(` >`)
//line views/templates/forms/json.qtpl:128
				qw422016.E().S(title)
//line views/templates/forms/json.qtpl:128
				qw422016.N().S(`</option> `)
//line views/templates/forms/json.qtpl:129
			}
//line views/templates/forms/json.qtpl:129
			qw422016.N().S(` </select> `)
//line views/templates/forms/json.qtpl:131
		} else {
//line views/templates/forms/json.qtpl:131
			qw422016.N().S(` <input type="`)
//line views/templates/forms/json.qtpl:132
			qw422016.E().S(col.InputType)
//line views/templates/forms/json.qtpl:132
			qw422016.N().S(`" `)
//line views/templates/forms/json.qtpl:132
			col.StreamRenderAttr(qw422016, i)
//line views/templates/forms/json.qtpl:132
			qw422016.N().S(` `)
//line views/templates/forms/json.qtpl:133
			if col.CharacterMaximumLength() > 0 {
//line views/templates/forms/json.qtpl:133
				qw422016.N().S(` max = "`)
//line views/templates/forms/json.qtpl:133
				qw422016.N().D(col.CharacterMaximumLength())
//line views/templates/forms/json.qtpl:133
				qw422016.N().S(`" `)
//line views/templates/forms/json.qtpl:133
			}
//line views/templates/forms/json.qtpl:133
			qw422016.N().S(` `)
//line views/templates/forms/json.qtpl:134
			if val != nil {
//line views/templates/forms/json.qtpl:134
				qw422016.N().S(` value = "`)
//line views/templates/forms/json.qtpl:134
				qw422016.E().V(val)
//line views/templates/forms/json.qtpl:134
				qw422016.N().S(`" `)
//line views/templates/forms/json.qtpl:134
			}
//line views/templates/forms/json.qtpl:134
			qw422016.N().S(` `)
//line views/templates/forms/json.qtpl:135
			if col.Pattern() > "" {
//line views/templates/forms/json.qtpl:135
				qw422016.N().S(` pattern="`)
//line views/templates/forms/json.qtpl:135
				qw422016.E().S(col.Pattern())
//line views/templates/forms/json.qtpl:135
				qw422016.N().S(`"  onkeyup="return validatePattern(this);" `)
//line views/templates/forms/json.qtpl:135
			}
//line views/templates/forms/json.qtpl:135
			qw422016.N().S(` `)
//line views/templates/forms/json.qtpl:136
			if col.Placeholder() > "" {
//line views/templates/forms/json.qtpl:136
				qw422016.N().S(` placeholder="`)
//line views/templates/forms/json.qtpl:136
				qw422016.E().S(col.Placeholder())
//line views/templates/forms/json.qtpl:136
				qw422016.N().S(`" `)
//line views/templates/forms/json.qtpl:136
			}
//line views/templates/forms/json.qtpl:136
			qw422016.N().S(` > `)
//line views/templates/forms/json.qtpl:138
		}
//line views/templates/forms/json.qtpl:138
		qw422016.N().S(` `)
//line views/templates/forms/json.qtpl:139
	}
//line views/templates/forms/json.qtpl:139
	qw422016.N().S(` `)
//line views/templates/forms/json.qtpl:140
	if col.IsSlice {
//line views/templates/forms/json.qtpl:140
		qw422016.N().S(` <button class="main-btn" type="button" onclick="this.parentNode.insertBefore(this.previousElementSibling.cloneNode(), this); return false;">+</button> `)
//line views/templates/forms/json.qtpl:144
	}
//line views/templates/forms/json.qtpl:144
	qw422016.N().S(` `)
//line views/templates/forms/json.qtpl:145
}

//line views/templates/forms/json.qtpl:145
func (col *ColumnDecor) WriteRenderInputs(qq422016 qtio422016.Writer) {
//line views/templates/forms/json.qtpl:145
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/forms/json.qtpl:145
	col.StreamRenderInputs(qw422016)
//line views/templates/forms/json.qtpl:145
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/json.qtpl:145
}

//line views/templates/forms/json.qtpl:145
func (col *ColumnDecor) RenderInputs() string {
//line views/templates/forms/json.qtpl:145
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/forms/json.qtpl:145
	col.WriteRenderInputs(qb422016)
//line views/templates/forms/json.qtpl:145
	qs422016 := string(qb422016.B)
//line views/templates/forms/json.qtpl:145
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/forms/json.qtpl:145
	return qs422016
//line views/templates/forms/json.qtpl:145
}

//line views/templates/forms/json.qtpl:147
func (f *FormField) StreamFormHTML(qw422016 *qt422016.Writer, blocks ...BlockColumns) {
//line views/templates/forms/json.qtpl:147
	qw422016.N().S(` <form id="`)
//line views/templates/forms/json.qtpl:148
	qw422016.E().S(f.Title)
//line views/templates/forms/json.qtpl:148
	qw422016.N().S(`form" name="`)
//line views/templates/forms/json.qtpl:148
	qw422016.E().S(f.Title)
//line views/templates/forms/json.qtpl:148
	qw422016.N().S(`" role='form' class="form-horizontal row-fluid" target="content" action="`)
//line views/templates/forms/json.qtpl:149
	qw422016.E().S(f.Action)
//line views/templates/forms/json.qtpl:149
	qw422016.N().S(`" method="`)
//line views/templates/forms/json.qtpl:149
	qw422016.N().S(f.Method)
//line views/templates/forms/json.qtpl:149
	qw422016.N().S(`" enctype="multipart/form-data" onsubmit="return saveForm(this, afterSaveAnyForm);"  caption="`)
//line views/templates/forms/json.qtpl:150
	qw422016.E().S(f.Title)
//line views/templates/forms/json.qtpl:150
	qw422016.N().S(`" > `)
//line views/templates/forms/json.qtpl:151
	for _, block := range blocks {
//line views/templates/forms/json.qtpl:151
		qw422016.N().S(` <div> <h2>`)
//line views/templates/forms/json.qtpl:153
		qw422016.E().S(block.Title)
//line views/templates/forms/json.qtpl:153
		qw422016.N().S(` </h2> <h3>`)
//line views/templates/forms/json.qtpl:154
		qw422016.E().S(block.Description)
//line views/templates/forms/json.qtpl:154
		qw422016.N().S(` </h3> `)
//line views/templates/forms/json.qtpl:155
		for i, col := range block.Columns {
//line views/templates/forms/json.qtpl:155
			qw422016.N().S(` <div id="divField`)
//line views/templates/forms/json.qtpl:156
			qw422016.N().D(i)
//line views/templates/forms/json.qtpl:156
			qw422016.N().S(`" class="input-wrap" `)
//line views/templates/forms/json.qtpl:157
			if col.IsHidden {
//line views/templates/forms/json.qtpl:157
				qw422016.N().S(` style="display:none" `)
//line views/templates/forms/json.qtpl:157
			}
//line views/templates/forms/json.qtpl:157
			qw422016.N().S(` > <label class="input-label" for="`)
//line views/templates/forms/json.qtpl:158
			qw422016.E().S(col.Name())
//line views/templates/forms/json.qtpl:158
			qw422016.N().S(`">`)
//line views/templates/forms/json.qtpl:158
			qw422016.E().S(col.Label)
//line views/templates/forms/json.qtpl:158
			qw422016.N().S(`</label> `)
//line views/templates/forms/json.qtpl:159
			col.StreamRenderInputs(qw422016)
//line views/templates/forms/json.qtpl:159
			qw422016.N().S(` </div> `)
//line views/templates/forms/json.qtpl:161
		}
//line views/templates/forms/json.qtpl:161
		qw422016.N().S(` </div> `)
//line views/templates/forms/json.qtpl:163
		if block.Multype {
//line views/templates/forms/json.qtpl:163
			qw422016.N().S(` <button class="main-btn" type="button" onclick="this.parentNode.insertBefore(this.previousElementSibling.cloneNode(true), this); return false;">+</button> `)
//line views/templates/forms/json.qtpl:167
		}
//line views/templates/forms/json.qtpl:167
		qw422016.N().S(` `)
//line views/templates/forms/json.qtpl:168
	}
//line views/templates/forms/json.qtpl:168
	qw422016.N().S(` <div class="form-actions"> <button class="main-btn" type="submit">Сохранить</button> </div> </form> <script> `)
//line views/templates/forms/json.qtpl:175
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
//line views/templates/forms/json.qtpl:201
}

//line views/templates/forms/json.qtpl:201
func (f *FormField) WriteFormHTML(qq422016 qtio422016.Writer, blocks ...BlockColumns) {
//line views/templates/forms/json.qtpl:201
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/forms/json.qtpl:201
	f.StreamFormHTML(qw422016, blocks...)
//line views/templates/forms/json.qtpl:201
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/json.qtpl:201
}

//line views/templates/forms/json.qtpl:201
func (f *FormField) FormHTML(blocks ...BlockColumns) string {
//line views/templates/forms/json.qtpl:201
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/forms/json.qtpl:201
	f.WriteFormHTML(qb422016, blocks...)
//line views/templates/forms/json.qtpl:201
	qs422016 := string(qb422016.B)
//line views/templates/forms/json.qtpl:201
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/forms/json.qtpl:201
	return qs422016
//line views/templates/forms/json.qtpl:201
}
