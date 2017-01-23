// This file is automatically generated by qtc from "anyForm.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line views/templates/forms/anyForm.qtpl:1
package forms

//line views/templates/forms/anyForm.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/templates/forms/anyForm.qtpl:1
import (
	"fmt"
	"strings"
)

// Здесь рисуем элементы ввода для полей, обозначающие связи таблицы многие-к-многим
// При этом имя поля однозначно определяет тип связи и таблицу, с которой связываемся

//line views/templates/forms/anyForm.qtpl:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/forms/anyForm.qtpl:7
func (field *FieldStructure) StreamRenderMultiSelect(qw422016 *qt422016.Writer, ns *FieldsTable, tablePrefix, key, val, titleLabel, required string) {
	//line views/templates/forms/anyForm.qtpl:7
	qw422016.N().S(`
    <div class="dropdown">
        <a class="dropdown-toggle" role="button" class="btn" data-toggle="modal" data-target="#`)
	//line views/templates/forms/anyForm.qtpl:9
	qw422016.E().S(tablePrefix + key)
	//line views/templates/forms/anyForm.qtpl:9
	qw422016.N().S(`" >
            <span>`)
	//line views/templates/forms/anyForm.qtpl:10
	qw422016.E().S(titleLabel)
	//line views/templates/forms/anyForm.qtpl:10
	qw422016.N().S(`</span>
            <b class="caret"></b>
        </a>
        <div class="modal hide" id="`)
	//line views/templates/forms/anyForm.qtpl:13
	qw422016.E().S(tablePrefix + key)
	//line views/templates/forms/anyForm.qtpl:13
	qw422016.N().S(`" tabindex="-1" role="dialog"
             aria-labelledby="`)
	//line views/templates/forms/anyForm.qtpl:14
	qw422016.E().S(tablePrefix + key)
	//line views/templates/forms/anyForm.qtpl:14
	qw422016.N().S(`ModalLabel" aria-hidden="true">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">X</button>
                <h3 id="`)
	//line views/templates/forms/anyForm.qtpl:17
	qw422016.E().S(tablePrefix + key)
	//line views/templates/forms/anyForm.qtpl:17
	qw422016.N().S(`ModalLabel">`)
	//line views/templates/forms/anyForm.qtpl:17
	qw422016.E().S(titleLabel)
	//line views/templates/forms/anyForm.qtpl:17
	qw422016.N().S(`</h3>
            </div>
            <div class="modal-body">
                <ul class="hor-menu" role="menu" aria-labelledby="dLabel">
                `)
	//line views/templates/forms/anyForm.qtpl:21
	field.getMultiSelect(ns)

	//line views/templates/forms/anyForm.qtpl:21
	qw422016.N().S(`
                `)
	//line views/templates/forms/anyForm.qtpl:22
	qw422016.N().S(field.Html)
	//line views/templates/forms/anyForm.qtpl:22
	qw422016.N().S(`
                </ul>
            </div>
            <div class="modal-footer">
                <input >
                <button onclick="return addNewItems(this);">Добавить</button>
            </div>
        </div>
    </div>
`)
//line views/templates/forms/anyForm.qtpl:31
}

//line views/templates/forms/anyForm.qtpl:31
func (field *FieldStructure) WriteRenderMultiSelect(qq422016 qtio422016.Writer, ns *FieldsTable, tablePrefix, key, val, titleLabel, required string) {
	//line views/templates/forms/anyForm.qtpl:31
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/forms/anyForm.qtpl:31
	field.StreamRenderMultiSelect(qw422016, ns, tablePrefix, key, val, titleLabel, required)
	//line views/templates/forms/anyForm.qtpl:31
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/anyForm.qtpl:31
}

//line views/templates/forms/anyForm.qtpl:31
func (field *FieldStructure) RenderMultiSelect(ns *FieldsTable, tablePrefix, key, val, titleLabel, required string) string {
	//line views/templates/forms/anyForm.qtpl:31
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/forms/anyForm.qtpl:31
	field.WriteRenderMultiSelect(qb422016, ns, tablePrefix, key, val, titleLabel, required)
	//line views/templates/forms/anyForm.qtpl:31
	qs422016 := string(qb422016.B)
	//line views/templates/forms/anyForm.qtpl:31
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/forms/anyForm.qtpl:31
	return qs422016
//line views/templates/forms/anyForm.qtpl:31
}

//line views/templates/forms/anyForm.qtpl:32
func (field *FieldStructure) StreamRenderForeignSelect(qw422016 *qt422016.Writer, tablePrefix, key, val, titleLabel, required string) {
	//line views/templates/forms/anyForm.qtpl:32
	qw422016.N().S(`
    <label class="control-label" for="`)
	//line views/templates/forms/anyForm.qtpl:33
	qw422016.E().S(tablePrefix + key)
	//line views/templates/forms/anyForm.qtpl:33
	qw422016.N().S(`">`)
	//line views/templates/forms/anyForm.qtpl:33
	qw422016.E().S(titleLabel)
	//line views/templates/forms/anyForm.qtpl:33
	qw422016.N().S(`</label>
    <select id="`)
	//line views/templates/forms/anyForm.qtpl:34
	qw422016.E().S(key)
	//line views/templates/forms/anyForm.qtpl:34
	qw422016.N().S(`" name="`)
	//line views/templates/forms/anyForm.qtpl:34
	qw422016.E().S(tablePrefix + key)
	//line views/templates/forms/anyForm.qtpl:34
	qw422016.N().S(`" class="controls" `)
	//line views/templates/forms/anyForm.qtpl:34
	qw422016.E().S(required)
	//line views/templates/forms/anyForm.qtpl:34
	qw422016.N().S(`>
        `)
	//line views/templates/forms/anyForm.qtpl:35
	if field.IS_NULLABLE == "YES" && val == "" {
		//line views/templates/forms/anyForm.qtpl:35
		qw422016.N().S(`
            <option value="" selected>Значение можно не указывать</option>
        `)
		//line views/templates/forms/anyForm.qtpl:37
	} else {
		//line views/templates/forms/anyForm.qtpl:37
		qw422016.N().S(`
            <option disabled >Выберите значение из списка</option>
        `)
		//line views/templates/forms/anyForm.qtpl:39
	}
	//line views/templates/forms/anyForm.qtpl:39
	qw422016.N().S(`
        `)
	//line views/templates/forms/anyForm.qtpl:40
	field.getOptions(key[3:], val)

	//line views/templates/forms/anyForm.qtpl:40
	qw422016.N().S(`
        `)
	//line views/templates/forms/anyForm.qtpl:41
	qw422016.N().S(field.Html)
	//line views/templates/forms/anyForm.qtpl:41
	qw422016.N().S(`
    </select>
`)
//line views/templates/forms/anyForm.qtpl:43
}

//line views/templates/forms/anyForm.qtpl:43
func (field *FieldStructure) WriteRenderForeignSelect(qq422016 qtio422016.Writer, tablePrefix, key, val, titleLabel, required string) {
	//line views/templates/forms/anyForm.qtpl:43
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/forms/anyForm.qtpl:43
	field.StreamRenderForeignSelect(qw422016, tablePrefix, key, val, titleLabel, required)
	//line views/templates/forms/anyForm.qtpl:43
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/anyForm.qtpl:43
}

//line views/templates/forms/anyForm.qtpl:43
func (field *FieldStructure) RenderForeignSelect(tablePrefix, key, val, titleLabel, required string) string {
	//line views/templates/forms/anyForm.qtpl:43
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/forms/anyForm.qtpl:43
	field.WriteRenderForeignSelect(qb422016, tablePrefix, key, val, titleLabel, required)
	//line views/templates/forms/anyForm.qtpl:43
	qs422016 := string(qb422016.B)
	//line views/templates/forms/anyForm.qtpl:43
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/forms/anyForm.qtpl:43
	return qs422016
//line views/templates/forms/anyForm.qtpl:43
}

//line views/templates/forms/anyForm.qtpl:45
func (field *FieldStructure) streamrenderParentSelect(qw422016 *qt422016.Writer, nameTable, key, val, titleLabel, required string) {
	//line views/templates/forms/anyForm.qtpl:45
	qw422016.N().S(`
    <label class="control-label" for="`)
	//line views/templates/forms/anyForm.qtpl:46
	qw422016.E().S(key)
	//line views/templates/forms/anyForm.qtpl:46
	qw422016.N().S(`">`)
	//line views/templates/forms/anyForm.qtpl:46
	qw422016.E().S(titleLabel)
	//line views/templates/forms/anyForm.qtpl:46
	qw422016.N().S(`</label>
    <select id="`)
	//line views/templates/forms/anyForm.qtpl:47
	qw422016.E().S(key)
	//line views/templates/forms/anyForm.qtpl:47
	qw422016.N().S(`" name="`)
	//line views/templates/forms/anyForm.qtpl:47
	qw422016.E().S(key)
	//line views/templates/forms/anyForm.qtpl:47
	qw422016.N().S(`" class="controls" `)
	//line views/templates/forms/anyForm.qtpl:47
	qw422016.E().S(required)
	//line views/templates/forms/anyForm.qtpl:47
	qw422016.N().S(`>
        `)
	//line views/templates/forms/anyForm.qtpl:48
	if field.IS_NULLABLE == "YES" && val == "" {
		//line views/templates/forms/anyForm.qtpl:48
		qw422016.N().S(`
            <option value="" selected>Значение можно не указывать</option>
        `)
		//line views/templates/forms/anyForm.qtpl:50
	} else {
		//line views/templates/forms/anyForm.qtpl:50
		qw422016.N().S(`
            <option disabled >Выберите значение из списка</option>
        `)
		//line views/templates/forms/anyForm.qtpl:52
	}
	//line views/templates/forms/anyForm.qtpl:52
	qw422016.N().S(`
        `)
	//line views/templates/forms/anyForm.qtpl:54
	if nameTable == "" {
		nameTable = field.TableName
	}

	//line views/templates/forms/anyForm.qtpl:57
	qw422016.N().S(`
        `)
	//line views/templates/forms/anyForm.qtpl:58
	field.getOptions(nameTable, val)

	//line views/templates/forms/anyForm.qtpl:58
	qw422016.N().S(`
        `)
	//line views/templates/forms/anyForm.qtpl:59
	qw422016.N().S(field.Html)
	//line views/templates/forms/anyForm.qtpl:59
	qw422016.N().S(`
    </select>
`)
//line views/templates/forms/anyForm.qtpl:61
}

//line views/templates/forms/anyForm.qtpl:61
func (field *FieldStructure) writerenderParentSelect(qq422016 qtio422016.Writer, nameTable, key, val, titleLabel, required string) {
	//line views/templates/forms/anyForm.qtpl:61
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/forms/anyForm.qtpl:61
	field.streamrenderParentSelect(qw422016, nameTable, key, val, titleLabel, required)
	//line views/templates/forms/anyForm.qtpl:61
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/anyForm.qtpl:61
}

//line views/templates/forms/anyForm.qtpl:61
func (field *FieldStructure) renderParentSelect(nameTable, key, val, titleLabel, required string) string {
	//line views/templates/forms/anyForm.qtpl:61
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/forms/anyForm.qtpl:61
	field.writerenderParentSelect(qb422016, nameTable, key, val, titleLabel, required)
	//line views/templates/forms/anyForm.qtpl:61
	qs422016 := string(qb422016.B)
	//line views/templates/forms/anyForm.qtpl:61
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/forms/anyForm.qtpl:61
	return qs422016
//line views/templates/forms/anyForm.qtpl:61
}

//line views/templates/forms/anyForm.qtpl:62
func (ns *FieldsTable) StreamShowAnyForm(qw422016 *qt422016.Writer, Action, Title string) {
	//line views/templates/forms/anyForm.qtpl:62
	qw422016.N().S(`
<form name="f`)
	//line views/templates/forms/anyForm.qtpl:63
	qw422016.E().S(ns.Name)
	//line views/templates/forms/anyForm.qtpl:63
	qw422016.N().S(`" role='form' class="form-horizontal row-fluid" target="content" action="`)
	//line views/templates/forms/anyForm.qtpl:63
	qw422016.E().S(Action)
	//line views/templates/forms/anyForm.qtpl:63
	qw422016.N().S(`" method="post"
      onsubmit="return saveForm(this, afterSaveAnyForm);" >
    <h2 class="form-signin-heading">`)
	//line views/templates/forms/anyForm.qtpl:65
	qw422016.E().S(Title)
	//line views/templates/forms/anyForm.qtpl:65
	qw422016.N().S(`</h2>
    `)
	//line views/templates/forms/anyForm.qtpl:66
	if ns.Name > "" {
		//line views/templates/forms/anyForm.qtpl:66
		qw422016.N().S(`
        <input type="hidden" name="table" value="`)
		//line views/templates/forms/anyForm.qtpl:67
		qw422016.E().S(ns.Name)
		//line views/templates/forms/anyForm.qtpl:67
		qw422016.N().S(`" >
    `)
		//line views/templates/forms/anyForm.qtpl:68
	}
	//line views/templates/forms/anyForm.qtpl:68
	qw422016.N().S(`
    `)
	//line views/templates/forms/anyForm.qtpl:69
	for idx, field := range ns.Rows {
		//line views/templates/forms/anyForm.qtpl:69
		qw422016.N().S(`

        `)
		//line views/templates/forms/anyForm.qtpl:72
		key := field.COLUMN_NAME
		titleFull, titleLabel, placeholder, pattern, dataJson := field.GetColumnTitles()
		val := field.Value
		tablePrefix := ""
		required := ""

		if field.IS_NULLABLE == "NO" {
			required = "required"
		}
		// TODO: check for new record

		if (val == "") && (field.COLUMN_DEFAULT > "") {
			val = field.COLUMN_DEFAULT
		}
		if dataJson > "" {
			dataJson = fmt.Sprintf(`data-names="{% s}"`, dataJson)
		}
		if field.TableName > "" {
			tablePrefix = field.TableName + ":"
		}
		nameInput := tablePrefix + key
		events := ""
		for name, funcName := range field.Events {
			events += fmt.Sprintf(`%s="return %s;"`, name, funcName)
		}

		//line views/templates/forms/anyForm.qtpl:97
		qw422016.N().S(`

        `)
		//line views/templates/forms/anyForm.qtpl:99
		if (val > "") && ((key == "id") || field.IsHidden) {
			//line views/templates/forms/anyForm.qtpl:99
			qw422016.N().S(`
            <input type="hidden" name="`)
			//line views/templates/forms/anyForm.qtpl:100
			qw422016.E().S(nameInput)
			//line views/templates/forms/anyForm.qtpl:100
			qw422016.N().S(`" `)
			//line views/templates/forms/anyForm.qtpl:100
			if val > "" {
				//line views/templates/forms/anyForm.qtpl:100
				qw422016.N().S(`value="`)
				//line views/templates/forms/anyForm.qtpl:100
				qw422016.E().S(val)
				//line views/templates/forms/anyForm.qtpl:100
				qw422016.N().S(`"`)
				//line views/templates/forms/anyForm.qtpl:100
			}
			//line views/templates/forms/anyForm.qtpl:100
			qw422016.N().S(` >
            `)
			//line views/templates/forms/anyForm.qtpl:101
			continue
			//line views/templates/forms/anyForm.qtpl:102
		} else if key == "id" {
			//line views/templates/forms/anyForm.qtpl:102
			qw422016.N().S(`
            `)
			//line views/templates/forms/anyForm.qtpl:103
			continue
			//line views/templates/forms/anyForm.qtpl:104
		}
		//line views/templates/forms/anyForm.qtpl:104
		qw422016.N().S(`

        <div id="divField`)
		//line views/templates/forms/anyForm.qtpl:106
		qw422016.N().D(idx)
		//line views/templates/forms/anyForm.qtpl:106
		qw422016.N().S(`" class="control-group `)
		//line views/templates/forms/anyForm.qtpl:106
		qw422016.N().S(required)
		//line views/templates/forms/anyForm.qtpl:106
		qw422016.N().S(` field-`)
		//line views/templates/forms/anyForm.qtpl:106
		qw422016.E().S(nameInput)
		//line views/templates/forms/anyForm.qtpl:106
		qw422016.N().S(` `)
		//line views/templates/forms/anyForm.qtpl:106
		qw422016.E().S(field.CSSClass)
		//line views/templates/forms/anyForm.qtpl:106
		qw422016.N().S(`"
            `)
		//line views/templates/forms/anyForm.qtpl:107
		if field.IsHidden {
			//line views/templates/forms/anyForm.qtpl:107
			qw422016.N().S(` style="display:none" `)
			//line views/templates/forms/anyForm.qtpl:107
		}
		//line views/templates/forms/anyForm.qtpl:107
		qw422016.N().S(`
            data-toggle="tooltip" title="`)
		//line views/templates/forms/anyForm.qtpl:108
		qw422016.E().S(titleFull)
		//line views/templates/forms/anyForm.qtpl:108
		qw422016.N().S(`"
        >
        `)
		//line views/templates/forms/anyForm.qtpl:110
		if key == "parent_id" {
			//line views/templates/forms/anyForm.qtpl:110
			qw422016.N().S(`
            `)
			//line views/templates/forms/anyForm.qtpl:111
			field.streamrenderParentSelect(qw422016, ns.Name, key, val, titleLabel, required)
			//line views/templates/forms/anyForm.qtpl:111
			qw422016.N().S(`
        `)
			//line views/templates/forms/anyForm.qtpl:112
		} else if strings.HasPrefix(key, "id_") {
			//line views/templates/forms/anyForm.qtpl:112
			qw422016.N().S(`
            `)
			//line views/templates/forms/anyForm.qtpl:113
			field.StreamRenderForeignSelect(qw422016, tablePrefix, key, val, titleLabel, required)
			//line views/templates/forms/anyForm.qtpl:113
			qw422016.N().S(`
        `)
			//line views/templates/forms/anyForm.qtpl:114
		} else if strings.HasPrefix(key, "setid_") {
			//line views/templates/forms/anyForm.qtpl:114
			qw422016.N().S(`
            `)
			//line views/templates/forms/anyForm.qtpl:115
			field.StreamRenderMultiSelect(qw422016, ns, tablePrefix, key, val, titleLabel, required)
			//line views/templates/forms/anyForm.qtpl:115
			qw422016.N().S(`
        `)
			//line views/templates/forms/anyForm.qtpl:116
		} else {
			//line views/templates/forms/anyForm.qtpl:116
			qw422016.N().S(`

                `)
			//line views/templates/forms/anyForm.qtpl:118
			switch field.DATA_TYPE {
			//line views/templates/forms/anyForm.qtpl:119
			case "tinyint":
				//line views/templates/forms/anyForm.qtpl:119
				qw422016.N().S(`
                        `)
				//line views/templates/forms/anyForm.qtpl:121
				checked := ""
				if val > "" {
					checked = "checked"
				}

				//line views/templates/forms/anyForm.qtpl:125
				qw422016.N().S(`
                        `)
				//line views/templates/forms/anyForm.qtpl:126
				streamrenderCheckBox(qw422016, nameInput, titleLabel, titleLabel, 1, checked, events, dataJson)
				//line views/templates/forms/anyForm.qtpl:126
				qw422016.N().S(`
                `)
			//line views/templates/forms/anyForm.qtpl:127
			case "enum":
				//line views/templates/forms/anyForm.qtpl:127
				qw422016.N().S(`
                    <label class="control-label" for="`)
				//line views/templates/forms/anyForm.qtpl:128
				qw422016.E().S(key)
				//line views/templates/forms/anyForm.qtpl:128
				qw422016.N().S(`">`)
				//line views/templates/forms/anyForm.qtpl:128
				qw422016.E().S(titleLabel)
				//line views/templates/forms/anyForm.qtpl:128
				qw422016.N().S(`:</label>
                    `)
				//line views/templates/forms/anyForm.qtpl:129
				t := field.renderEnum(nameInput, val, required, events, dataJson)

				//line views/templates/forms/anyForm.qtpl:129
				qw422016.N().S(`
                    `)
				//line views/templates/forms/anyForm.qtpl:130
				qw422016.N().S(t)
				//line views/templates/forms/anyForm.qtpl:130
				qw422016.N().S(`
                `)
			//line views/templates/forms/anyForm.qtpl:131
			case "set":
				//line views/templates/forms/anyForm.qtpl:131
				qw422016.N().S(`
                    <label class="control-label" for="`)
				//line views/templates/forms/anyForm.qtpl:132
				qw422016.E().S(key)
				//line views/templates/forms/anyForm.qtpl:132
				qw422016.N().S(`">`)
				//line views/templates/forms/anyForm.qtpl:132
				qw422016.E().S(titleLabel)
				//line views/templates/forms/anyForm.qtpl:132
				qw422016.N().S(`:</label>
                    `)
				//line views/templates/forms/anyForm.qtpl:133
				t := field.renderSet(nameInput, val, required, events, dataJson)

				//line views/templates/forms/anyForm.qtpl:133
				qw422016.N().S(`
                    `)
				//line views/templates/forms/anyForm.qtpl:134
				qw422016.N().S(t)
				//line views/templates/forms/anyForm.qtpl:134
				qw422016.N().S(`
                `)
			//line views/templates/forms/anyForm.qtpl:135
			case "blob", "text":
				//line views/templates/forms/anyForm.qtpl:135
				qw422016.N().S(`
                    <label class="control-label" for="`)
				//line views/templates/forms/anyForm.qtpl:136
				qw422016.E().S(key)
				//line views/templates/forms/anyForm.qtpl:136
				qw422016.N().S(`">`)
				//line views/templates/forms/anyForm.qtpl:136
				qw422016.E().S(titleLabel)
				//line views/templates/forms/anyForm.qtpl:136
				qw422016.N().S(`:</label>
                    <textarea id="`)
				//line views/templates/forms/anyForm.qtpl:137
				qw422016.E().S(key)
				//line views/templates/forms/anyForm.qtpl:137
				qw422016.N().S(`" name="`)
				//line views/templates/forms/anyForm.qtpl:137
				qw422016.E().S(nameInput)
				//line views/templates/forms/anyForm.qtpl:137
				qw422016.N().S(`" class="controls" placeholder="`)
				//line views/templates/forms/anyForm.qtpl:137
				qw422016.E().S(placeholder)
				//line views/templates/forms/anyForm.qtpl:137
				qw422016.N().S(`">
                        `)
				//line views/templates/forms/anyForm.qtpl:138
				qw422016.N().S(val)
				//line views/templates/forms/anyForm.qtpl:138
				qw422016.N().S(`
                    </textarea>
               `)
			//line views/templates/forms/anyForm.qtpl:140
			default:
				//line views/templates/forms/anyForm.qtpl:140
				qw422016.N().S(`
                    <label class="control-label" for="`)
				//line views/templates/forms/anyForm.qtpl:141
				qw422016.E().S(key)
				//line views/templates/forms/anyForm.qtpl:141
				qw422016.N().S(`">`)
				//line views/templates/forms/anyForm.qtpl:141
				qw422016.E().S(titleLabel)
				//line views/templates/forms/anyForm.qtpl:141
				qw422016.N().S(`</label>
                    <input id="`)
				//line views/templates/forms/anyForm.qtpl:142
				qw422016.E().S(key)
				//line views/templates/forms/anyForm.qtpl:142
				qw422016.N().S(`" name="`)
				//line views/templates/forms/anyForm.qtpl:142
				qw422016.E().S(nameInput)
				//line views/templates/forms/anyForm.qtpl:142
				qw422016.N().S(`" class="controls" placeholder="`)
				//line views/templates/forms/anyForm.qtpl:142
				qw422016.E().S(placeholder)
				//line views/templates/forms/anyForm.qtpl:142
				qw422016.N().S(`"
                        `)
				//line views/templates/forms/anyForm.qtpl:143
				if val > "" {
					//line views/templates/forms/anyForm.qtpl:143
					qw422016.N().S(` value="`)
					//line views/templates/forms/anyForm.qtpl:143
					qw422016.E().S(val)
					//line views/templates/forms/anyForm.qtpl:143
					qw422016.N().S(`" `)
					//line views/templates/forms/anyForm.qtpl:143
				}
				//line views/templates/forms/anyForm.qtpl:143
				qw422016.N().S(`
                        `)
				//line views/templates/forms/anyForm.qtpl:144
				qw422016.N().S(required)
				//line views/templates/forms/anyForm.qtpl:144
				qw422016.N().S(` `)
				//line views/templates/forms/anyForm.qtpl:144
				qw422016.N().S(events)
				//line views/templates/forms/anyForm.qtpl:144
				qw422016.N().S(` `)
				//line views/templates/forms/anyForm.qtpl:144
				qw422016.N().S(dataJson)
				//line views/templates/forms/anyForm.qtpl:144
				qw422016.N().S(` type="

                        `)
				//line views/templates/forms/anyForm.qtpl:146
				if field.DATA_TYPE == "int" || field.DATA_TYPE == "double" {
					//line views/templates/forms/anyForm.qtpl:146
					qw422016.N().S(`
                            number"
                               `)
					//line views/templates/forms/anyForm.qtpl:148
					if strings.Contains(field.COLUMN_TYPE, "unsigned") {
						//line views/templates/forms/anyForm.qtpl:148
						qw422016.N().S(`min="0"`)
						//line views/templates/forms/anyForm.qtpl:148
					}
					//line views/templates/forms/anyForm.qtpl:148
					qw422016.N().S(`
                        `)
					//line views/templates/forms/anyForm.qtpl:149
				} else if field.DATA_TYPE == "date" {
					//line views/templates/forms/anyForm.qtpl:149
					qw422016.N().S(`
                            date"
                        `)
					//line views/templates/forms/anyForm.qtpl:151
				} else if field.DATA_TYPE == "datetime" {
					//line views/templates/forms/anyForm.qtpl:151
					qw422016.N().S(`
                            datetime"
                        `)
					//line views/templates/forms/anyForm.qtpl:153
				} else if strings.Contains(key, "email") {
					//line views/templates/forms/anyForm.qtpl:153
					qw422016.N().S(`
                            email"
                        `)
					//line views/templates/forms/anyForm.qtpl:155
				} else {
					//line views/templates/forms/anyForm.qtpl:155
					qw422016.N().S(`
                            text"
                               `)
					//line views/templates/forms/anyForm.qtpl:157
					if field.CHARACTER_MAXIMUM_LENGTH > 0 {
						//line views/templates/forms/anyForm.qtpl:157
						qw422016.N().S(`
                                    maxlength="`)
						//line views/templates/forms/anyForm.qtpl:158
						qw422016.N().D(field.CHARACTER_MAXIMUM_LENGTH)
						//line views/templates/forms/anyForm.qtpl:158
						qw422016.N().S(`"
                                    `)
						//line views/templates/forms/anyForm.qtpl:159
						if field.CHARACTER_MAXIMUM_LENGTH == 255 {
							//line views/templates/forms/anyForm.qtpl:159
							qw422016.N().S(`
                                        class="input-xlarge"
                                    `)
							//line views/templates/forms/anyForm.qtpl:161
						}
						//line views/templates/forms/anyForm.qtpl:161
						qw422016.N().S(`
                               `)
						//line views/templates/forms/anyForm.qtpl:162
					}
					//line views/templates/forms/anyForm.qtpl:162
					qw422016.N().S(`
                               `)
					//line views/templates/forms/anyForm.qtpl:163
					if pattern > "" {
						//line views/templates/forms/anyForm.qtpl:163
						qw422016.N().S(`pattern="`)
						//line views/templates/forms/anyForm.qtpl:163
						qw422016.N().S(pattern)
						//line views/templates/forms/anyForm.qtpl:163
						qw422016.N().S(`" onkeyup="return validatePattern(this);"`)
						//line views/templates/forms/anyForm.qtpl:163
					}
					//line views/templates/forms/anyForm.qtpl:163
					qw422016.N().S(`
                        `)
					//line views/templates/forms/anyForm.qtpl:164
				}
				//line views/templates/forms/anyForm.qtpl:164
				qw422016.N().S(`
                    />
                `)
				//line views/templates/forms/anyForm.qtpl:166
			}
			//line views/templates/forms/anyForm.qtpl:166
			qw422016.N().S(`

        `)
			//line views/templates/forms/anyForm.qtpl:168
		}
		//line views/templates/forms/anyForm.qtpl:168
		qw422016.N().S(`
        </div>
    `)
		//line views/templates/forms/anyForm.qtpl:170
	}
	//line views/templates/forms/anyForm.qtpl:170
	qw422016.N().S(`
    <div class="form-actions">
        <button class="btn btn-large btn-primary" type="submit">Сохранить</button>

    </div>
</form>
`)
//line views/templates/forms/anyForm.qtpl:176
}

//line views/templates/forms/anyForm.qtpl:176
func (ns *FieldsTable) WriteShowAnyForm(qq422016 qtio422016.Writer, Action, Title string) {
	//line views/templates/forms/anyForm.qtpl:176
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/forms/anyForm.qtpl:176
	ns.StreamShowAnyForm(qw422016, Action, Title)
	//line views/templates/forms/anyForm.qtpl:176
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/anyForm.qtpl:176
}

//line views/templates/forms/anyForm.qtpl:176
func (ns *FieldsTable) ShowAnyForm(Action, Title string) string {
	//line views/templates/forms/anyForm.qtpl:176
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/forms/anyForm.qtpl:176
	ns.WriteShowAnyForm(qb422016, Action, Title)
	//line views/templates/forms/anyForm.qtpl:176
	qs422016 := string(qb422016.B)
	//line views/templates/forms/anyForm.qtpl:176
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/forms/anyForm.qtpl:176
	return qs422016
//line views/templates/forms/anyForm.qtpl:176
}

//line views/templates/forms/anyForm.qtpl:177
func streamrenderCheckBox(qw422016 *qt422016.Writer, key, val, title string, idx int, checked, events, dataJson string) {
	//line views/templates/forms/anyForm.qtpl:177
	qw422016.N().S(`
    <label class="checkbox" for="`)
	//line views/templates/forms/anyForm.qtpl:178
	qw422016.E().S(key)
	//line views/templates/forms/anyForm.qtpl:178
	qw422016.N().D(idx)
	//line views/templates/forms/anyForm.qtpl:178
	qw422016.N().S(`">
        <input type="checkbox" id="`)
	//line views/templates/forms/anyForm.qtpl:179
	qw422016.E().S(key)
	//line views/templates/forms/anyForm.qtpl:179
	qw422016.N().D(idx)
	//line views/templates/forms/anyForm.qtpl:179
	qw422016.N().S(`" name="`)
	//line views/templates/forms/anyForm.qtpl:179
	qw422016.E().S(key)
	//line views/templates/forms/anyForm.qtpl:179
	qw422016.N().S(`" value="`)
	//line views/templates/forms/anyForm.qtpl:179
	qw422016.E().S(val)
	//line views/templates/forms/anyForm.qtpl:179
	qw422016.N().S(`" `)
	//line views/templates/forms/anyForm.qtpl:179
	qw422016.N().S(checked)
	//line views/templates/forms/anyForm.qtpl:179
	qw422016.N().S(`
                `)
	//line views/templates/forms/anyForm.qtpl:180
	qw422016.N().S(events)
	//line views/templates/forms/anyForm.qtpl:180
	qw422016.N().S(` `)
	//line views/templates/forms/anyForm.qtpl:180
	qw422016.N().S(dataJson)
	//line views/templates/forms/anyForm.qtpl:180
	qw422016.N().S(`
        />
        `)
	//line views/templates/forms/anyForm.qtpl:182
	qw422016.E().S(title)
	//line views/templates/forms/anyForm.qtpl:182
	qw422016.N().S(`
    </label>
`)
//line views/templates/forms/anyForm.qtpl:184
}

//line views/templates/forms/anyForm.qtpl:184
func writerenderCheckBox(qq422016 qtio422016.Writer, key, val, title string, idx int, checked, events, dataJson string) {
	//line views/templates/forms/anyForm.qtpl:184
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/forms/anyForm.qtpl:184
	streamrenderCheckBox(qw422016, key, val, title, idx, checked, events, dataJson)
	//line views/templates/forms/anyForm.qtpl:184
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/anyForm.qtpl:184
}

//line views/templates/forms/anyForm.qtpl:184
func renderCheckBox(key, val, title string, idx int, checked, events, dataJson string) string {
	//line views/templates/forms/anyForm.qtpl:184
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/forms/anyForm.qtpl:184
	writerenderCheckBox(qb422016, key, val, title, idx, checked, events, dataJson)
	//line views/templates/forms/anyForm.qtpl:184
	qs422016 := string(qb422016.B)
	//line views/templates/forms/anyForm.qtpl:184
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/forms/anyForm.qtpl:184
	return qs422016
//line views/templates/forms/anyForm.qtpl:184
}

//line views/templates/forms/anyForm.qtpl:185
func streamrenderRadioBox(qw422016 *qt422016.Writer, key, val, title string, idx int, checked, events, dataJson string) {
	//line views/templates/forms/anyForm.qtpl:185
	qw422016.N().S(`
    <label class="checkbox" for="`)
	//line views/templates/forms/anyForm.qtpl:186
	qw422016.E().S(key)
	//line views/templates/forms/anyForm.qtpl:186
	qw422016.N().D(idx)
	//line views/templates/forms/anyForm.qtpl:186
	qw422016.N().S(`">
        <input type="radio" id="`)
	//line views/templates/forms/anyForm.qtpl:187
	qw422016.E().S(key)
	//line views/templates/forms/anyForm.qtpl:187
	qw422016.N().D(idx)
	//line views/templates/forms/anyForm.qtpl:187
	qw422016.N().S(`" name="`)
	//line views/templates/forms/anyForm.qtpl:187
	qw422016.E().S(key)
	//line views/templates/forms/anyForm.qtpl:187
	qw422016.N().S(`" value="`)
	//line views/templates/forms/anyForm.qtpl:187
	qw422016.E().S(val)
	//line views/templates/forms/anyForm.qtpl:187
	qw422016.N().S(`" `)
	//line views/templates/forms/anyForm.qtpl:187
	qw422016.N().S(checked)
	//line views/templates/forms/anyForm.qtpl:187
	qw422016.N().S(`
                `)
	//line views/templates/forms/anyForm.qtpl:188
	qw422016.N().S(events)
	//line views/templates/forms/anyForm.qtpl:188
	qw422016.N().S(` `)
	//line views/templates/forms/anyForm.qtpl:188
	qw422016.N().S(dataJson)
	//line views/templates/forms/anyForm.qtpl:188
	qw422016.N().S(`
        />
        `)
	//line views/templates/forms/anyForm.qtpl:190
	qw422016.E().S(title)
	//line views/templates/forms/anyForm.qtpl:190
	qw422016.N().S(`
    </label>
`)
//line views/templates/forms/anyForm.qtpl:192
}

//line views/templates/forms/anyForm.qtpl:192
func writerenderRadioBox(qq422016 qtio422016.Writer, key, val, title string, idx int, checked, events, dataJson string) {
	//line views/templates/forms/anyForm.qtpl:192
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/forms/anyForm.qtpl:192
	streamrenderRadioBox(qw422016, key, val, title, idx, checked, events, dataJson)
	//line views/templates/forms/anyForm.qtpl:192
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/anyForm.qtpl:192
}

//line views/templates/forms/anyForm.qtpl:192
func renderRadioBox(key, val, title string, idx int, checked, events, dataJson string) string {
	//line views/templates/forms/anyForm.qtpl:192
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/forms/anyForm.qtpl:192
	writerenderRadioBox(qb422016, key, val, title, idx, checked, events, dataJson)
	//line views/templates/forms/anyForm.qtpl:192
	qs422016 := string(qb422016.B)
	//line views/templates/forms/anyForm.qtpl:192
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/forms/anyForm.qtpl:192
	return qs422016
//line views/templates/forms/anyForm.qtpl:192
}

//line views/templates/forms/anyForm.qtpl:193
func streamrenderSelect(qw422016 *qt422016.Writer, key, options, required, events, dataJson string) {
	//line views/templates/forms/anyForm.qtpl:193
	qw422016.N().S(`
    <select id="`)
	//line views/templates/forms/anyForm.qtpl:194
	qw422016.E().S(key)
	//line views/templates/forms/anyForm.qtpl:194
	qw422016.N().S(`" name="`)
	//line views/templates/forms/anyForm.qtpl:194
	qw422016.E().S(key)
	//line views/templates/forms/anyForm.qtpl:194
	qw422016.N().S(`" class="controls" `)
	//line views/templates/forms/anyForm.qtpl:194
	qw422016.N().S(required)
	//line views/templates/forms/anyForm.qtpl:194
	qw422016.N().S(` `)
	//line views/templates/forms/anyForm.qtpl:194
	qw422016.N().S(events)
	//line views/templates/forms/anyForm.qtpl:194
	qw422016.N().S(` `)
	//line views/templates/forms/anyForm.qtpl:194
	qw422016.N().S(dataJson)
	//line views/templates/forms/anyForm.qtpl:194
	qw422016.N().S(`>
        `)
	//line views/templates/forms/anyForm.qtpl:195
	qw422016.N().S(options)
	//line views/templates/forms/anyForm.qtpl:195
	qw422016.N().S(`
    </select>
`)
//line views/templates/forms/anyForm.qtpl:197
}

//line views/templates/forms/anyForm.qtpl:197
func writerenderSelect(qq422016 qtio422016.Writer, key, options, required, events, dataJson string) {
	//line views/templates/forms/anyForm.qtpl:197
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/forms/anyForm.qtpl:197
	streamrenderSelect(qw422016, key, options, required, events, dataJson)
	//line views/templates/forms/anyForm.qtpl:197
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/anyForm.qtpl:197
}

//line views/templates/forms/anyForm.qtpl:197
func renderSelect(key, options, required, events, dataJson string) string {
	//line views/templates/forms/anyForm.qtpl:197
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/forms/anyForm.qtpl:197
	writerenderSelect(qb422016, key, options, required, events, dataJson)
	//line views/templates/forms/anyForm.qtpl:197
	qs422016 := string(qb422016.B)
	//line views/templates/forms/anyForm.qtpl:197
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/forms/anyForm.qtpl:197
	return qs422016
//line views/templates/forms/anyForm.qtpl:197
}

//line views/templates/forms/anyForm.qtpl:198
func streamrenderOption(qw422016 *qt422016.Writer, val, title, selected string) {
	//line views/templates/forms/anyForm.qtpl:198
	qw422016.N().S(`
    <option value="`)
	//line views/templates/forms/anyForm.qtpl:199
	qw422016.E().S(val)
	//line views/templates/forms/anyForm.qtpl:199
	qw422016.N().S(`" `)
	//line views/templates/forms/anyForm.qtpl:199
	qw422016.N().S(selected)
	//line views/templates/forms/anyForm.qtpl:199
	qw422016.N().S(` >`)
	//line views/templates/forms/anyForm.qtpl:199
	qw422016.E().S(title)
	//line views/templates/forms/anyForm.qtpl:199
	qw422016.N().S(`</option>
`)
//line views/templates/forms/anyForm.qtpl:200
}

//line views/templates/forms/anyForm.qtpl:200
func writerenderOption(qq422016 qtio422016.Writer, val, title, selected string) {
	//line views/templates/forms/anyForm.qtpl:200
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/forms/anyForm.qtpl:200
	streamrenderOption(qw422016, val, title, selected)
	//line views/templates/forms/anyForm.qtpl:200
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/anyForm.qtpl:200
}

//line views/templates/forms/anyForm.qtpl:200
func renderOption(val, title, selected string) string {
	//line views/templates/forms/anyForm.qtpl:200
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/forms/anyForm.qtpl:200
	writerenderOption(qb422016, val, title, selected)
	//line views/templates/forms/anyForm.qtpl:200
	qs422016 := string(qb422016.B)
	//line views/templates/forms/anyForm.qtpl:200
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/forms/anyForm.qtpl:200
	return qs422016
//line views/templates/forms/anyForm.qtpl:200
}
