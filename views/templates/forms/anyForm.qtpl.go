// This file is automatically generated by qtc from "anyForm.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line anyForm.qtpl:1
package forms

//line anyForm.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line anyForm.qtpl:1
import (
	"fmt"
	"strings"
)

// Показываем связанную таблицу RenderTable

//line anyForm.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line anyForm.qtpl:6
func (field *FieldStructure) StreamRenderTable(qw422016 *qt422016.Writer, ns *FieldsTable, tablePrefix, key, val, titleLabel, required string) {
	//line anyForm.qtpl:6
	qw422016.N().S(`
    <link href="/tables.css" rel="stylesheet">
    `)
	//line anyForm.qtpl:8
	if field.CSSClass == "" {
		//line anyForm.qtpl:8
		qw422016.N().S(`
        <div class="dropdown">
            <a class="dropdown-toggle" role="button" class="btn" data-toggle="modal" data-target="#div`)
		//line anyForm.qtpl:10
		qw422016.E().S(tablePrefix + key)
		//line anyForm.qtpl:10
		qw422016.N().S(`" >
                <span>`)
		//line anyForm.qtpl:11
		qw422016.E().S(titleLabel)
		//line anyForm.qtpl:11
		qw422016.N().S(`</span>
                <b class="caret"></b>
            </a>
    `)
		//line anyForm.qtpl:14
	}
	//line anyForm.qtpl:14
	qw422016.N().S(`
            <div class="`)
	//line anyForm.qtpl:15
	if field.CSSClass == "" {
		//line anyForm.qtpl:15
		qw422016.N().S(`modal hide`)
		//line anyForm.qtpl:15
	} else {
		//line anyForm.qtpl:15
		qw422016.E().S(field.CSSClass)
		//line anyForm.qtpl:15
	}
	//line anyForm.qtpl:15
	qw422016.N().S(`"
                id="div`)
	//line anyForm.qtpl:16
	qw422016.E().S(tablePrefix + key)
	//line anyForm.qtpl:16
	qw422016.N().S(`" tabindex="-1" role="dialog" aria-labelledby="`)
	//line anyForm.qtpl:16
	qw422016.E().S(tablePrefix + key)
	//line anyForm.qtpl:16
	qw422016.N().S(`ModalLabel"
                aria-hidden="true">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-hidden="true">X</button>
                     <h3 id="`)
	//line anyForm.qtpl:20
	qw422016.E().S(tablePrefix + key)
	//line anyForm.qtpl:20
	qw422016.N().S(`ModalLabel">`)
	//line anyForm.qtpl:20
	qw422016.E().S(titleLabel)
	//line anyForm.qtpl:20
	qw422016.N().S(`</h3>
                </div>
                <div class="modal-body">
                    <table  class="table table-striped table-bordered table-hover table-condensed" role="menu" aria-labelledby="dLabel">
                        `)
	//line anyForm.qtpl:24
	field.getTableFrom(ns, tablePrefix, key)

	//line anyForm.qtpl:24
	qw422016.N().S(`
                        `)
	//line anyForm.qtpl:25
	qw422016.N().S(field.Html)
	//line anyForm.qtpl:25
	qw422016.N().S(`
                    </table>
                </div>
                <div class="modal-footer">
                    <button onclick="return addNewRowTableID(this);" data-last-tr="tr`)
	//line anyForm.qtpl:29
	qw422016.E().S(tablePrefix + key)
	//line anyForm.qtpl:29
	qw422016.N().S(`">Добавить</button>
                </div>
            </div>
    `)
	//line anyForm.qtpl:32
	if field.CSSClass == "" {
		//line anyForm.qtpl:32
		qw422016.N().S(`
        </div> `)
		//line anyForm.qtpl:33
		qw422016.N().S(`
    `)
		//line anyForm.qtpl:34
	}
	//line anyForm.qtpl:34
	qw422016.N().S(`
`)
//line anyForm.qtpl:35
}

//line anyForm.qtpl:35
func (field *FieldStructure) WriteRenderTable(qq422016 qtio422016.Writer, ns *FieldsTable, tablePrefix, key, val, titleLabel, required string) {
	//line anyForm.qtpl:35
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line anyForm.qtpl:35
	field.StreamRenderTable(qw422016, ns, tablePrefix, key, val, titleLabel, required)
	//line anyForm.qtpl:35
	qt422016.ReleaseWriter(qw422016)
//line anyForm.qtpl:35
}

//line anyForm.qtpl:35
func (field *FieldStructure) RenderTable(ns *FieldsTable, tablePrefix, key, val, titleLabel, required string) string {
	//line anyForm.qtpl:35
	qb422016 := qt422016.AcquireByteBuffer()
	//line anyForm.qtpl:35
	field.WriteRenderTable(qb422016, ns, tablePrefix, key, val, titleLabel, required)
	//line anyForm.qtpl:35
	qs422016 := string(qb422016.B)
	//line anyForm.qtpl:35
	qt422016.ReleaseByteBuffer(qb422016)
	//line anyForm.qtpl:35
	return qs422016
//line anyForm.qtpl:35
}

// Здесь рисуем элементы ввода для полей, обозначающие связи таблицы многие-к-многим
// При этом имя поля однозначно определяет тип связи и таблицу, с которой связываемся

//line anyForm.qtpl:38
func (field *FieldStructure) StreamRenderMultiSelect(qw422016 *qt422016.Writer, ns *FieldsTable, tablePrefix, key, val, titleLabel, required string) {
	//line anyForm.qtpl:38
	qw422016.N().S(`
    `)
	//line anyForm.qtpl:39
	if field.CSSClass == "" {
		//line anyForm.qtpl:39
		qw422016.N().S(`
    <div class="dropdown">
        <a class="dropdown-toggle" role="button" class="btn" data-toggle="modal" data-target="#div`)
		//line anyForm.qtpl:41
		qw422016.E().S(tablePrefix + key)
		//line anyForm.qtpl:41
		qw422016.N().S(`" >
            <span>`)
		//line anyForm.qtpl:42
		qw422016.E().S(titleLabel)
		//line anyForm.qtpl:42
		qw422016.N().S(`</span>
            <b class="caret"></b>
        </a>
    `)
		//line anyForm.qtpl:45
	}
	//line anyForm.qtpl:45
	qw422016.N().S(`
        <div class="`)
	//line anyForm.qtpl:46
	if field.CSSClass == "" {
		//line anyForm.qtpl:46
		qw422016.N().S(`modal hide`)
		//line anyForm.qtpl:46
	} else {
		//line anyForm.qtpl:46
		qw422016.E().S(field.CSSClass)
		//line anyForm.qtpl:46
	}
	//line anyForm.qtpl:46
	qw422016.N().S(`"
            id="div`)
	//line anyForm.qtpl:47
	qw422016.E().S(tablePrefix + key)
	//line anyForm.qtpl:47
	qw422016.N().S(`" tabindex="-1" role="dialog"
             aria-labelledby="`)
	//line anyForm.qtpl:48
	qw422016.E().S(tablePrefix + key)
	//line anyForm.qtpl:48
	qw422016.N().S(`ModalLabel" aria-hidden="true">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">X</button>
                <h3 id="`)
	//line anyForm.qtpl:51
	qw422016.E().S(tablePrefix + key)
	//line anyForm.qtpl:51
	qw422016.N().S(`ModalLabel">`)
	//line anyForm.qtpl:51
	qw422016.E().S(titleLabel)
	//line anyForm.qtpl:51
	qw422016.N().S(`</h3>
            </div>
            <div class="modal-body">
                <ul class="hor-menu" role="menu" aria-labelledby="dLabel">
                `)
	//line anyForm.qtpl:55
	field.getMultiSelect(ns, key)

	//line anyForm.qtpl:55
	qw422016.N().S(`
                `)
	//line anyForm.qtpl:56
	qw422016.N().S(field.Html)
	//line anyForm.qtpl:56
	qw422016.N().S(`
                </ul>
            </div>
            <div class="modal-footer">
                <input >
                <button onclick="return addNewItems(this);" data-parent-div="div`)
	//line anyForm.qtpl:61
	qw422016.E().S(tablePrefix + key)
	//line anyForm.qtpl:61
	qw422016.N().S(`">Добавить</button>
            </div>
        </div>
    `)
	//line anyForm.qtpl:64
	if field.CSSClass == "" {
		//line anyForm.qtpl:64
		qw422016.N().S(`
        </div> `)
		//line anyForm.qtpl:65
		qw422016.N().S(`
    `)
		//line anyForm.qtpl:66
	}
	//line anyForm.qtpl:66
	qw422016.N().S(`
`)
//line anyForm.qtpl:67
}

//line anyForm.qtpl:67
func (field *FieldStructure) WriteRenderMultiSelect(qq422016 qtio422016.Writer, ns *FieldsTable, tablePrefix, key, val, titleLabel, required string) {
	//line anyForm.qtpl:67
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line anyForm.qtpl:67
	field.StreamRenderMultiSelect(qw422016, ns, tablePrefix, key, val, titleLabel, required)
	//line anyForm.qtpl:67
	qt422016.ReleaseWriter(qw422016)
//line anyForm.qtpl:67
}

//line anyForm.qtpl:67
func (field *FieldStructure) RenderMultiSelect(ns *FieldsTable, tablePrefix, key, val, titleLabel, required string) string {
	//line anyForm.qtpl:67
	qb422016 := qt422016.AcquireByteBuffer()
	//line anyForm.qtpl:67
	field.WriteRenderMultiSelect(qb422016, ns, tablePrefix, key, val, titleLabel, required)
	//line anyForm.qtpl:67
	qs422016 := string(qb422016.B)
	//line anyForm.qtpl:67
	qt422016.ReleaseByteBuffer(qb422016)
	//line anyForm.qtpl:67
	return qs422016
//line anyForm.qtpl:67
}

//line anyForm.qtpl:68
func (field *FieldStructure) StreamRenderForeignSelect(qw422016 *qt422016.Writer, tablePrefix, key, val, titleLabel, required, events, dataJson string) {
	//line anyForm.qtpl:68
	qw422016.N().S(`
    <label class="control-label" for="`)
	//line anyForm.qtpl:69
	qw422016.E().S(tablePrefix + key)
	//line anyForm.qtpl:69
	qw422016.N().S(`">`)
	//line anyForm.qtpl:69
	qw422016.E().S(titleLabel)
	//line anyForm.qtpl:69
	qw422016.N().S(`</label>
    <select id="`)
	//line anyForm.qtpl:70
	qw422016.E().S(key)
	//line anyForm.qtpl:70
	qw422016.N().S(`" name="`)
	//line anyForm.qtpl:70
	qw422016.E().S(tablePrefix + key)
	//line anyForm.qtpl:70
	qw422016.N().S(`" class="controls" `)
	//line anyForm.qtpl:70
	qw422016.E().S(required)
	//line anyForm.qtpl:70
	qw422016.N().S(`  `)
	//line anyForm.qtpl:70
	qw422016.N().S(events)
	//line anyForm.qtpl:70
	qw422016.N().S(` `)
	//line anyForm.qtpl:70
	qw422016.N().S(dataJson)
	//line anyForm.qtpl:70
	qw422016.N().S(`>
        `)
	//line anyForm.qtpl:71
	if field.IS_NULLABLE == "YES" && val == "" {
		//line anyForm.qtpl:71
		qw422016.N().S(`
            <option value="" selected>Значение можно не указывать</option>
        `)
		//line anyForm.qtpl:73
	} else {
		//line anyForm.qtpl:73
		qw422016.N().S(`
            <option disabled >Выберите значение из списка</option>
        `)
		//line anyForm.qtpl:75
	}
	//line anyForm.qtpl:75
	qw422016.N().S(`
        `)
	//line anyForm.qtpl:76
	field.getOptions(key[3:], val)

	//line anyForm.qtpl:76
	qw422016.N().S(`
        `)
	//line anyForm.qtpl:77
	qw422016.N().S(field.Html)
	//line anyForm.qtpl:77
	qw422016.N().S(`
    </select>
`)
//line anyForm.qtpl:79
}

//line anyForm.qtpl:79
func (field *FieldStructure) WriteRenderForeignSelect(qq422016 qtio422016.Writer, tablePrefix, key, val, titleLabel, required, events, dataJson string) {
	//line anyForm.qtpl:79
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line anyForm.qtpl:79
	field.StreamRenderForeignSelect(qw422016, tablePrefix, key, val, titleLabel, required, events, dataJson)
	//line anyForm.qtpl:79
	qt422016.ReleaseWriter(qw422016)
//line anyForm.qtpl:79
}

//line anyForm.qtpl:79
func (field *FieldStructure) RenderForeignSelect(tablePrefix, key, val, titleLabel, required, events, dataJson string) string {
	//line anyForm.qtpl:79
	qb422016 := qt422016.AcquireByteBuffer()
	//line anyForm.qtpl:79
	field.WriteRenderForeignSelect(qb422016, tablePrefix, key, val, titleLabel, required, events, dataJson)
	//line anyForm.qtpl:79
	qs422016 := string(qb422016.B)
	//line anyForm.qtpl:79
	qt422016.ReleaseByteBuffer(qb422016)
	//line anyForm.qtpl:79
	return qs422016
//line anyForm.qtpl:79
}

//line anyForm.qtpl:81
func (field *FieldStructure) StreamRenderParentSelect(qw422016 *qt422016.Writer, nameTable, key, val, titleLabel, required, events, dataJson string) {
	//line anyForm.qtpl:81
	qw422016.N().S(`
    <label class="control-label" for="`)
	//line anyForm.qtpl:82
	qw422016.E().S(key)
	//line anyForm.qtpl:82
	qw422016.N().S(`">`)
	//line anyForm.qtpl:82
	qw422016.E().S(titleLabel)
	//line anyForm.qtpl:82
	qw422016.N().S(`</label>
    <select id="`)
	//line anyForm.qtpl:83
	qw422016.E().S(key)
	//line anyForm.qtpl:83
	qw422016.N().S(`" name="`)
	//line anyForm.qtpl:83
	qw422016.E().S(key)
	//line anyForm.qtpl:83
	qw422016.N().S(`" class="controls" `)
	//line anyForm.qtpl:83
	qw422016.E().S(required)
	//line anyForm.qtpl:83
	qw422016.N().S(`  `)
	//line anyForm.qtpl:83
	qw422016.N().S(events)
	//line anyForm.qtpl:83
	qw422016.N().S(` `)
	//line anyForm.qtpl:83
	qw422016.N().S(dataJson)
	//line anyForm.qtpl:83
	qw422016.N().S(`>
        `)
	//line anyForm.qtpl:84
	if field.IS_NULLABLE == "YES" && val == "" {
		//line anyForm.qtpl:84
		qw422016.N().S(`
            <option value="" selected>Значение можно не указывать</option>
        `)
		//line anyForm.qtpl:86
	} else {
		//line anyForm.qtpl:86
		qw422016.N().S(`
            <option disabled >Выберите значение из списка</option>
        `)
		//line anyForm.qtpl:88
	}
	//line anyForm.qtpl:88
	qw422016.N().S(`
        `)
	//line anyForm.qtpl:90
	if nameTable == "" {
		nameTable = field.TableName
	}

	//line anyForm.qtpl:93
	qw422016.N().S(`
        `)
	//line anyForm.qtpl:94
	field.getOptions(nameTable, val)

	//line anyForm.qtpl:94
	qw422016.N().S(`
        `)
	//line anyForm.qtpl:95
	qw422016.N().S(field.Html)
	//line anyForm.qtpl:95
	qw422016.N().S(`
    </select>
`)
//line anyForm.qtpl:97
}

//line anyForm.qtpl:97
func (field *FieldStructure) WriteRenderParentSelect(qq422016 qtio422016.Writer, nameTable, key, val, titleLabel, required, events, dataJson string) {
	//line anyForm.qtpl:97
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line anyForm.qtpl:97
	field.StreamRenderParentSelect(qw422016, nameTable, key, val, titleLabel, required, events, dataJson)
	//line anyForm.qtpl:97
	qt422016.ReleaseWriter(qw422016)
//line anyForm.qtpl:97
}

//line anyForm.qtpl:97
func (field *FieldStructure) RenderParentSelect(nameTable, key, val, titleLabel, required, events, dataJson string) string {
	//line anyForm.qtpl:97
	qb422016 := qt422016.AcquireByteBuffer()
	//line anyForm.qtpl:97
	field.WriteRenderParentSelect(qb422016, nameTable, key, val, titleLabel, required, events, dataJson)
	//line anyForm.qtpl:97
	qs422016 := string(qb422016.B)
	//line anyForm.qtpl:97
	qt422016.ReleaseByteBuffer(qb422016)
	//line anyForm.qtpl:97
	return qs422016
//line anyForm.qtpl:97
}

//line anyForm.qtpl:98
func (ns *FieldsTable) StreamShowAnyForm(qw422016 *qt422016.Writer, Action, Title string) {
	//line anyForm.qtpl:98
	qw422016.N().S(`
`)
	//line anyForm.qtpl:100
	var figure string
	if ns.SaveFormEvents == nil {
		ns.SaveFormEvents = make(map[string]string, 1)
		ns.SaveFormEvents["successSaveForm"] = "afterSaveAnyForm"
	}
	formName := "f" + ns.Name

	//line anyForm.qtpl:106
	qw422016.N().S(`
<h2 class="form-signin-heading">`)
	//line anyForm.qtpl:107
	qw422016.E().S(Title)
	//line anyForm.qtpl:107
	qw422016.N().S(`</h2>
<form id=="`)
	//line anyForm.qtpl:108
	qw422016.E().S(formName)
	//line anyForm.qtpl:108
	qw422016.N().S(`" name="`)
	//line anyForm.qtpl:108
	qw422016.E().S(formName)
	//line anyForm.qtpl:108
	qw422016.N().S(`" role='form' class="form-horizontal row-fluid" target="content" action="`)
	//line anyForm.qtpl:108
	qw422016.E().S(Action)
	//line anyForm.qtpl:108
	qw422016.N().S(`" method="post"
      onsubmit="return saveForm(this, `)
	//line anyForm.qtpl:109
	qw422016.N().S(ns.SaveFormEvents["successSaveForm"])
	//line anyForm.qtpl:109
	qw422016.N().S(`);" >
      `)
	//line anyForm.qtpl:111
	setDatePicker := ""

	//line anyForm.qtpl:112
	qw422016.N().S(`
    `)
	//line anyForm.qtpl:113
	if ns.Name > "" {
		//line anyForm.qtpl:113
		qw422016.N().S(`
        <input type="hidden" name="table" value="`)
		//line anyForm.qtpl:114
		qw422016.E().S(ns.Name)
		//line anyForm.qtpl:114
		qw422016.N().S(`" >
    `)
		//line anyForm.qtpl:115
	}
	//line anyForm.qtpl:115
	qw422016.N().S(`
    `)
	//line anyForm.qtpl:116
	for idx, field := range ns.Rows {
		//line anyForm.qtpl:116
		qw422016.N().S(`

        `)
		//line anyForm.qtpl:119
		key := field.COLUMN_NAME

		if field.DATA_TYPE == "date" {
			setDatePicker += "$('#" + key + `').datetimepicker({format:'Y-m-d'});
                `
		} else if field.DATA_TYPE == "datetime" {
			setDatePicker += "$('#" + key + `').datetimepicker({format:'Y-m-d H:i'});
                `
		}

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

		//line anyForm.qtpl:153
		qw422016.N().S(`

        `)
		//line anyForm.qtpl:155
		if figure != field.Figure {
			//line anyForm.qtpl:155
			qw422016.N().S(`
            `)
			//line anyForm.qtpl:156
			if figure > "" {
				//line anyForm.qtpl:156
				qw422016.N().S(`
                </figure>
            `)
				//line anyForm.qtpl:158
			}
			//line anyForm.qtpl:158
			qw422016.N().S(`
            `)
			//line anyForm.qtpl:160
			figure = field.Figure

			//line anyForm.qtpl:161
			qw422016.N().S(`
            `)
			//line anyForm.qtpl:162
			if figure > "" {
				//line anyForm.qtpl:162
				qw422016.N().S(`
                <figure class="`)
				//line anyForm.qtpl:163
				qw422016.E().S(field.CSSClass)
				//line anyForm.qtpl:163
				qw422016.N().S(`">
                <figcaption>`)
				//line anyForm.qtpl:164
				qw422016.E().S(figure)
				//line anyForm.qtpl:164
				qw422016.N().S(`</figcaption>
            `)
				//line anyForm.qtpl:165
			}
			//line anyForm.qtpl:165
			qw422016.N().S(`
        `)
			//line anyForm.qtpl:166
		}
		//line anyForm.qtpl:166
		qw422016.N().S(`


        `)
		//line anyForm.qtpl:169
		if (val > "") && ((key == "id") || field.IsHidden) {
			//line anyForm.qtpl:169
			qw422016.N().S(`
            <input type="hidden" name="`)
			//line anyForm.qtpl:170
			qw422016.E().S(nameInput)
			//line anyForm.qtpl:170
			qw422016.N().S(`" `)
			//line anyForm.qtpl:170
			if val > "" {
				//line anyForm.qtpl:170
				qw422016.N().S(`value="`)
				//line anyForm.qtpl:170
				qw422016.E().S(val)
				//line anyForm.qtpl:170
				qw422016.N().S(`"`)
				//line anyForm.qtpl:170
			}
			//line anyForm.qtpl:170
			qw422016.N().S(` >
            `)
			//line anyForm.qtpl:171
			continue
			//line anyForm.qtpl:172
		} else if key == "id" {
			//line anyForm.qtpl:172
			qw422016.N().S(`
            `)
			//line anyForm.qtpl:173
			continue
			//line anyForm.qtpl:174
		}
		//line anyForm.qtpl:174
		qw422016.N().S(`

        <div id="divField`)
		//line anyForm.qtpl:176
		qw422016.N().D(idx)
		//line anyForm.qtpl:176
		qw422016.N().S(`" class="input-wrap `)
		//line anyForm.qtpl:176
		qw422016.N().S(required)
		//line anyForm.qtpl:176
		qw422016.N().S(` field-`)
		//line anyForm.qtpl:176
		qw422016.E().S(nameInput)
		//line anyForm.qtpl:176
		qw422016.N().S(` `)
		//line anyForm.qtpl:176
		qw422016.E().S(field.CSSClass)
		//line anyForm.qtpl:176
		qw422016.N().S(`"
            `)
		//line anyForm.qtpl:177
		if field.IsHidden {
			//line anyForm.qtpl:177
			qw422016.N().S(` style="display:none" `)
			//line anyForm.qtpl:177
		}
		//line anyForm.qtpl:177
		qw422016.N().S(`
            data-toggle="tooltip" title="`)
		//line anyForm.qtpl:178
		qw422016.E().S(titleFull)
		//line anyForm.qtpl:178
		qw422016.N().S(`"
        >
        `)
		//line anyForm.qtpl:180
		if key == "parent_id" {
			//line anyForm.qtpl:180
			qw422016.N().S(`
            `)
			//line anyForm.qtpl:181
			field.StreamRenderParentSelect(qw422016, ns.Name, key, val, titleLabel, required, events, dataJson)
			//line anyForm.qtpl:181
			qw422016.N().S(`
        `)
			//line anyForm.qtpl:182
		} else if strings.HasPrefix(key, "id_") {
			//line anyForm.qtpl:182
			qw422016.N().S(`
            `)
			//line anyForm.qtpl:183
			field.StreamRenderForeignSelect(qw422016, tablePrefix, key, val, titleLabel, required, events, dataJson)
			//line anyForm.qtpl:183
			qw422016.N().S(`
        `)
			//line anyForm.qtpl:184
		} else if strings.HasPrefix(key, "setid_") || strings.HasPrefix(key, "nodeid_") {
			//line anyForm.qtpl:184
			qw422016.N().S(`
            `)
			//line anyForm.qtpl:185
			field.StreamRenderMultiSelect(qw422016, ns, tablePrefix, key, val, titleLabel, required)
			//line anyForm.qtpl:185
			qw422016.N().S(`
        `)
			//line anyForm.qtpl:186
		} else if strings.HasPrefix(key, "tableid_") {
			//line anyForm.qtpl:186
			qw422016.N().S(`
            `)
			//line anyForm.qtpl:187
			field.StreamRenderTable(qw422016, ns, tablePrefix, key, val, titleLabel, required)
			//line anyForm.qtpl:187
			qw422016.N().S(`
        `)
			//line anyForm.qtpl:188
		} else {
			//line anyForm.qtpl:188
			qw422016.N().S(`

                `)
			//line anyForm.qtpl:190
			switch field.DATA_TYPE {
			//line anyForm.qtpl:191
			case "tinyint":
				//line anyForm.qtpl:191
				qw422016.N().S(`
                        `)
				//line anyForm.qtpl:193
				checked := ""
				if val == "1" {
					checked = "checked"
				}

				//line anyForm.qtpl:197
				qw422016.N().S(`
                        `)
				//line anyForm.qtpl:198
				StreamRenderCheckBox(qw422016, nameInput, "1", titleLabel, 1, checked, events, dataJson)
				//line anyForm.qtpl:198
				qw422016.N().S(`
                `)
			//line anyForm.qtpl:199
			case "enum":
				//line anyForm.qtpl:199
				qw422016.N().S(`
                    <label class="input-label" for="`)
				//line anyForm.qtpl:200
				qw422016.E().S(key)
				//line anyForm.qtpl:200
				qw422016.N().S(`">`)
				//line anyForm.qtpl:200
				qw422016.E().S(titleLabel)
				//line anyForm.qtpl:200
				qw422016.N().S(`:</label>
                    `)
				//line anyForm.qtpl:201
				t := field.RenderEnum(nameInput, val, required, events, dataJson)

				//line anyForm.qtpl:201
				qw422016.N().S(`
                    `)
				//line anyForm.qtpl:202
				qw422016.N().S(t)
				//line anyForm.qtpl:202
				qw422016.N().S(`
                `)
			//line anyForm.qtpl:203
			case "set":
				//line anyForm.qtpl:203
				qw422016.N().S(`
                    <label class="input-label" for="`)
				//line anyForm.qtpl:204
				qw422016.E().S(key)
				//line anyForm.qtpl:204
				qw422016.N().S(`">`)
				//line anyForm.qtpl:204
				qw422016.E().S(titleLabel)
				//line anyForm.qtpl:204
				qw422016.N().S(`:</label>
                    `)
				//line anyForm.qtpl:205
				t := field.RenderSet(nameInput, val, required, events, dataJson)

				//line anyForm.qtpl:205
				qw422016.N().S(`
                    `)
				//line anyForm.qtpl:206
				qw422016.N().S(t)
				//line anyForm.qtpl:206
				qw422016.N().S(`
                `)
			//line anyForm.qtpl:207
			case "blob":
				//line anyForm.qtpl:207
				qw422016.N().S(`
                     <label class="input-label" for="`)
				//line anyForm.qtpl:208
				qw422016.E().S(key)
				//line anyForm.qtpl:208
				qw422016.N().S(`">`)
				//line anyForm.qtpl:208
				qw422016.E().S(titleLabel)
				//line anyForm.qtpl:208
				qw422016.N().S(`:</label>
                     <input type="file"
                        id="`)
				//line anyForm.qtpl:210
				qw422016.E().S(nameInput)
				//line anyForm.qtpl:210
				qw422016.N().S(`" name="`)
				//line anyForm.qtpl:210
				qw422016.E().S(nameInput)
				//line anyForm.qtpl:210
				qw422016.N().S(`" class="controls"
                                                `)
				//line anyForm.qtpl:211
				qw422016.N().S(required)
				//line anyForm.qtpl:211
				qw422016.N().S(` `)
				//line anyForm.qtpl:211
				qw422016.N().S(events)
				//line anyForm.qtpl:211
				qw422016.N().S(` `)
				//line anyForm.qtpl:211
				qw422016.N().S(dataJson)
				//line anyForm.qtpl:211
				qw422016.N().S(`
                     />
               `)
			//line anyForm.qtpl:213
			case "text":
				//line anyForm.qtpl:213
				qw422016.N().S(`
                        <label class="control-label" for="`)
				//line anyForm.qtpl:214
				qw422016.E().S(key)
				//line anyForm.qtpl:214
				qw422016.N().S(`">`)
				//line anyForm.qtpl:214
				qw422016.E().S(titleLabel)
				//line anyForm.qtpl:214
				qw422016.N().S(`:</label>
                    <textarea id="`)
				//line anyForm.qtpl:215
				qw422016.E().S(key)
				//line anyForm.qtpl:215
				qw422016.N().S(`" name="`)
				//line anyForm.qtpl:215
				qw422016.E().S(nameInput)
				//line anyForm.qtpl:215
				qw422016.N().S(`" class="controls" placeholder="`)
				//line anyForm.qtpl:215
				qw422016.E().S(placeholder)
				//line anyForm.qtpl:215
				qw422016.N().S(`" `)
				//line anyForm.qtpl:215
				qw422016.N().S(events)
				//line anyForm.qtpl:215
				qw422016.N().S(` `)
				//line anyForm.qtpl:215
				qw422016.N().S(dataJson)
				//line anyForm.qtpl:215
				qw422016.N().S(`>
                        `)
				//line anyForm.qtpl:216
				qw422016.N().S(val)
				//line anyForm.qtpl:216
				qw422016.N().S(`
                    </textarea>
               `)
			//line anyForm.qtpl:218
			default:
				//line anyForm.qtpl:218
				qw422016.N().S(`
                    `)
				//line anyForm.qtpl:219
				field.StreamRenderInputFromType(qw422016, nameInput, val, titleLabel, placeholder, pattern, required, events, dataJson)
				//line anyForm.qtpl:219
				qw422016.N().S(`
               `)
				//line anyForm.qtpl:220
			}
			//line anyForm.qtpl:220
			qw422016.N().S(`

        `)
			//line anyForm.qtpl:222
		}
		//line anyForm.qtpl:222
		qw422016.N().S(`
        </div>
    `)
		//line anyForm.qtpl:224
	}
	//line anyForm.qtpl:224
	qw422016.N().S(`
    `)
	//line anyForm.qtpl:225
	if figure > "" {
		//line anyForm.qtpl:225
		qw422016.N().S(`
        </figure>
    `)
		//line anyForm.qtpl:227
	}
	//line anyForm.qtpl:227
	qw422016.N().S(`
    <div class="form-actions">
        <button class="main-btn" type="submit">Сохранить</button>

    </div>
</form>
`)
	//line anyForm.qtpl:233
	if setDatePicker > "" {
		//line anyForm.qtpl:233
		qw422016.N().S(`
    <head>
        <script src="/jquery.datetimepicker.js"></script>
    </head>
`)
		//line anyForm.qtpl:237
	}
	//line anyForm.qtpl:237
	qw422016.N().S(`
<script>
 $(function(){
    $(".business-form-select").styler();
    $(".custom-input-label").click(function(){
    	$(this).addClass('small-label');
    });
    $('.business-form-input').keyup(function(){
    	if($(this).val() != ''){
    		$(this).next(".custom-input-label").addClass('small-label');
    	} else {
    		$(this).next(".custom-input-label").removeClass('small-label');
    		}
    });

    `)
	//line anyForm.qtpl:252
	if setDatePicker > "" {
		//line anyForm.qtpl:252
		qw422016.N().S(`
            `)
		//line anyForm.qtpl:253
		qw422016.N().S(setDatePicker)
		//line anyForm.qtpl:253
		qw422016.N().S(`
    `)
		//line anyForm.qtpl:254
	}
	//line anyForm.qtpl:254
	qw422016.N().S(`
    `)
	//line anyForm.qtpl:255
	if onload, ok := ns.DataJSOM["onload"]; ok {
		//line anyForm.qtpl:255
		qw422016.N().S(`
             `)
		//line anyForm.qtpl:256
		qw422016.N().S(onload.(string))
		//line anyForm.qtpl:256
		qw422016.N().S(`
    `)
		//line anyForm.qtpl:257
	}
	//line anyForm.qtpl:257
	qw422016.N().S(`
})
</script>
`)
//line anyForm.qtpl:260
}

//line anyForm.qtpl:260
func (ns *FieldsTable) WriteShowAnyForm(qq422016 qtio422016.Writer, Action, Title string) {
	//line anyForm.qtpl:260
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line anyForm.qtpl:260
	ns.StreamShowAnyForm(qw422016, Action, Title)
	//line anyForm.qtpl:260
	qt422016.ReleaseWriter(qw422016)
//line anyForm.qtpl:260
}

//line anyForm.qtpl:260
func (ns *FieldsTable) ShowAnyForm(Action, Title string) string {
	//line anyForm.qtpl:260
	qb422016 := qt422016.AcquireByteBuffer()
	//line anyForm.qtpl:260
	ns.WriteShowAnyForm(qb422016, Action, Title)
	//line anyForm.qtpl:260
	qs422016 := string(qb422016.B)
	//line anyForm.qtpl:260
	qt422016.ReleaseByteBuffer(qb422016)
	//line anyForm.qtpl:260
	return qs422016
//line anyForm.qtpl:260
}

//line anyForm.qtpl:261
func (fieldStrc *FieldStructure) StreamRenderInputFromType(qw422016 *qt422016.Writer, nameInput, val, title, placeholder, pattern, required, events, dataJson string) {
	//line anyForm.qtpl:261
	qw422016.N().S(`
                    <span class="input-label">`)
	//line anyForm.qtpl:262
	qw422016.E().S(title)
	//line anyForm.qtpl:262
	qw422016.N().S(`</span>
                    <div class="form-items-wrap">
                        <label class="custom-input-label" for="`)
	//line anyForm.qtpl:264
	qw422016.E().S(nameInput)
	//line anyForm.qtpl:264
	qw422016.N().S(`">`)
	//line anyForm.qtpl:264
	qw422016.E().S(placeholder)
	//line anyForm.qtpl:264
	qw422016.N().S(`</label>
                        <input type=

                        `)
	//line anyForm.qtpl:267
	if fieldStrc.InputType > "" {
		//line anyForm.qtpl:267
		qw422016.N().S(`
                            "`)
		//line anyForm.qtpl:268
		qw422016.E().S(fieldStrc.InputType)
		//line anyForm.qtpl:268
		qw422016.N().S(`"
                        `)
		//line anyForm.qtpl:269
	} else if fieldStrc.DATA_TYPE == "int" || fieldStrc.DATA_TYPE == "double" {
		//line anyForm.qtpl:269
		qw422016.N().S(`
                            "number"
                               `)
		//line anyForm.qtpl:271
		if strings.Contains(fieldStrc.COLUMN_TYPE, "unsigned") {
			//line anyForm.qtpl:271
			qw422016.N().S(`min="0"`)
			//line anyForm.qtpl:271
		}
		//line anyForm.qtpl:271
		qw422016.N().S(`
                        `)
		//line anyForm.qtpl:272
	} else if fieldStrc.DATA_TYPE == "date" {
		//line anyForm.qtpl:272
		qw422016.N().S(`
                            "date"
                        `)
		//line anyForm.qtpl:274
	} else if fieldStrc.DATA_TYPE == "datetime" {
		//line anyForm.qtpl:274
		qw422016.N().S(`
                            "datetime"
                        `)
		//line anyForm.qtpl:276
	} else if strings.Contains(nameInput, "email") {
		//line anyForm.qtpl:276
		qw422016.N().S(`
                            "email"
                        `)
		//line anyForm.qtpl:278
	} else {
		//line anyForm.qtpl:278
		qw422016.N().S(`
                            "text"
                               `)
		//line anyForm.qtpl:280
		if fieldStrc.CHARACTER_MAXIMUM_LENGTH > 0 {
			//line anyForm.qtpl:280
			qw422016.N().S(`
                                    maxlength="`)
			//line anyForm.qtpl:281
			qw422016.N().D(fieldStrc.CHARACTER_MAXIMUM_LENGTH)
			//line anyForm.qtpl:281
			qw422016.N().S(`"
                               `)
			//line anyForm.qtpl:282
		}
		//line anyForm.qtpl:282
		qw422016.N().S(`
                               `)
		//line anyForm.qtpl:283
		if pattern > "" {
			//line anyForm.qtpl:283
			qw422016.N().S(`pattern="`)
			//line anyForm.qtpl:283
			qw422016.N().S(pattern)
			//line anyForm.qtpl:283
			qw422016.N().S(`" onkeyup="return validatePattern(this);"`)
			//line anyForm.qtpl:283
		}
		//line anyForm.qtpl:283
		qw422016.N().S(`
                        `)
		//line anyForm.qtpl:284
	}
	//line anyForm.qtpl:284
	qw422016.N().S(`
                        id="`)
	//line anyForm.qtpl:285
	qw422016.E().S(nameInput)
	//line anyForm.qtpl:285
	qw422016.N().S(`" name="`)
	//line anyForm.qtpl:285
	qw422016.E().S(nameInput)
	//line anyForm.qtpl:285
	qw422016.N().S(`" class="business-form-input"
                                                `)
	//line anyForm.qtpl:286
	if val > "" {
		//line anyForm.qtpl:286
		qw422016.N().S(` value="`)
		//line anyForm.qtpl:286
		qw422016.E().S(val)
		//line anyForm.qtpl:286
		qw422016.N().S(`" `)
		//line anyForm.qtpl:286
	}
	//line anyForm.qtpl:286
	qw422016.N().S(`
                                                `)
	//line anyForm.qtpl:287
	qw422016.N().S(required)
	//line anyForm.qtpl:287
	qw422016.N().S(` `)
	//line anyForm.qtpl:287
	qw422016.N().S(events)
	//line anyForm.qtpl:287
	qw422016.N().S(` `)
	//line anyForm.qtpl:287
	qw422016.N().S(dataJson)
	//line anyForm.qtpl:287
	qw422016.N().S(`
                    />
                    </div>
`)
//line anyForm.qtpl:290
}

//line anyForm.qtpl:290
func (fieldStrc *FieldStructure) WriteRenderInputFromType(qq422016 qtio422016.Writer, nameInput, val, title, placeholder, pattern, required, events, dataJson string) {
	//line anyForm.qtpl:290
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line anyForm.qtpl:290
	fieldStrc.StreamRenderInputFromType(qw422016, nameInput, val, title, placeholder, pattern, required, events, dataJson)
	//line anyForm.qtpl:290
	qt422016.ReleaseWriter(qw422016)
//line anyForm.qtpl:290
}

//line anyForm.qtpl:290
func (fieldStrc *FieldStructure) RenderInputFromType(nameInput, val, title, placeholder, pattern, required, events, dataJson string) string {
	//line anyForm.qtpl:290
	qb422016 := qt422016.AcquireByteBuffer()
	//line anyForm.qtpl:290
	fieldStrc.WriteRenderInputFromType(qb422016, nameInput, val, title, placeholder, pattern, required, events, dataJson)
	//line anyForm.qtpl:290
	qs422016 := string(qb422016.B)
	//line anyForm.qtpl:290
	qt422016.ReleaseByteBuffer(qb422016)
	//line anyForm.qtpl:290
	return qs422016
//line anyForm.qtpl:290
}

//line anyForm.qtpl:292
func StreamRenderCheckBox(qw422016 *qt422016.Writer, key, val, title string, idx int, checked, events, dataJson string) {
	//line anyForm.qtpl:292
	qw422016.N().S(`
    <label class="checkbox" for="`)
	//line anyForm.qtpl:293
	qw422016.E().S(key)
	//line anyForm.qtpl:293
	qw422016.N().D(idx)
	//line anyForm.qtpl:293
	qw422016.N().S(`">
        <input type="checkbox" id="`)
	//line anyForm.qtpl:294
	qw422016.E().S(key)
	//line anyForm.qtpl:294
	qw422016.N().D(idx)
	//line anyForm.qtpl:294
	qw422016.N().S(`" name="`)
	//line anyForm.qtpl:294
	qw422016.E().S(key)
	//line anyForm.qtpl:294
	qw422016.N().S(`" value="`)
	//line anyForm.qtpl:294
	qw422016.E().S(val)
	//line anyForm.qtpl:294
	qw422016.N().S(`" `)
	//line anyForm.qtpl:294
	qw422016.N().S(checked)
	//line anyForm.qtpl:294
	qw422016.N().S(`
                `)
	//line anyForm.qtpl:295
	qw422016.N().S(events)
	//line anyForm.qtpl:295
	qw422016.N().S(` `)
	//line anyForm.qtpl:295
	qw422016.N().S(dataJson)
	//line anyForm.qtpl:295
	qw422016.N().S(`
        />
        `)
	//line anyForm.qtpl:297
	qw422016.E().S(title)
	//line anyForm.qtpl:297
	qw422016.N().S(`
    </label>
`)
//line anyForm.qtpl:299
}

//line anyForm.qtpl:299
func WriteRenderCheckBox(qq422016 qtio422016.Writer, key, val, title string, idx int, checked, events, dataJson string) {
	//line anyForm.qtpl:299
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line anyForm.qtpl:299
	StreamRenderCheckBox(qw422016, key, val, title, idx, checked, events, dataJson)
	//line anyForm.qtpl:299
	qt422016.ReleaseWriter(qw422016)
//line anyForm.qtpl:299
}

//line anyForm.qtpl:299
func RenderCheckBox(key, val, title string, idx int, checked, events, dataJson string) string {
	//line anyForm.qtpl:299
	qb422016 := qt422016.AcquireByteBuffer()
	//line anyForm.qtpl:299
	WriteRenderCheckBox(qb422016, key, val, title, idx, checked, events, dataJson)
	//line anyForm.qtpl:299
	qs422016 := string(qb422016.B)
	//line anyForm.qtpl:299
	qt422016.ReleaseByteBuffer(qb422016)
	//line anyForm.qtpl:299
	return qs422016
//line anyForm.qtpl:299
}

//line anyForm.qtpl:300
func streamrenderRadioBox(qw422016 *qt422016.Writer, key, val, title string, idx int, checked, events, dataJson string) {
	//line anyForm.qtpl:300
	qw422016.N().S(`
    <label class="checkbox" for="`)
	//line anyForm.qtpl:301
	qw422016.E().S(key)
	//line anyForm.qtpl:301
	qw422016.N().D(idx)
	//line anyForm.qtpl:301
	qw422016.N().S(`">
        <input type="radio" id="`)
	//line anyForm.qtpl:302
	qw422016.E().S(key)
	//line anyForm.qtpl:302
	qw422016.N().D(idx)
	//line anyForm.qtpl:302
	qw422016.N().S(`" name="`)
	//line anyForm.qtpl:302
	qw422016.E().S(key)
	//line anyForm.qtpl:302
	qw422016.N().S(`" value="`)
	//line anyForm.qtpl:302
	qw422016.E().S(val)
	//line anyForm.qtpl:302
	qw422016.N().S(`" `)
	//line anyForm.qtpl:302
	qw422016.N().S(checked)
	//line anyForm.qtpl:302
	qw422016.N().S(`
                `)
	//line anyForm.qtpl:303
	qw422016.N().S(events)
	//line anyForm.qtpl:303
	qw422016.N().S(` `)
	//line anyForm.qtpl:303
	qw422016.N().S(dataJson)
	//line anyForm.qtpl:303
	qw422016.N().S(`
        />
        `)
	//line anyForm.qtpl:305
	qw422016.E().S(title)
	//line anyForm.qtpl:305
	qw422016.N().S(`
    </label>
`)
//line anyForm.qtpl:307
}

//line anyForm.qtpl:307
func writerenderRadioBox(qq422016 qtio422016.Writer, key, val, title string, idx int, checked, events, dataJson string) {
	//line anyForm.qtpl:307
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line anyForm.qtpl:307
	streamrenderRadioBox(qw422016, key, val, title, idx, checked, events, dataJson)
	//line anyForm.qtpl:307
	qt422016.ReleaseWriter(qw422016)
//line anyForm.qtpl:307
}

//line anyForm.qtpl:307
func renderRadioBox(key, val, title string, idx int, checked, events, dataJson string) string {
	//line anyForm.qtpl:307
	qb422016 := qt422016.AcquireByteBuffer()
	//line anyForm.qtpl:307
	writerenderRadioBox(qb422016, key, val, title, idx, checked, events, dataJson)
	//line anyForm.qtpl:307
	qs422016 := string(qb422016.B)
	//line anyForm.qtpl:307
	qt422016.ReleaseByteBuffer(qb422016)
	//line anyForm.qtpl:307
	return qs422016
//line anyForm.qtpl:307
}

//line anyForm.qtpl:308
func streamrenderSelect(qw422016 *qt422016.Writer, key, options, required, events, dataJson string) {
	//line anyForm.qtpl:308
	qw422016.N().S(`
    <select id="`)
	//line anyForm.qtpl:309
	qw422016.E().S(key)
	//line anyForm.qtpl:309
	qw422016.N().S(`" name="`)
	//line anyForm.qtpl:309
	qw422016.E().S(key)
	//line anyForm.qtpl:309
	qw422016.N().S(`" class="business-form-select" `)
	//line anyForm.qtpl:309
	qw422016.N().S(required)
	//line anyForm.qtpl:309
	qw422016.N().S(` `)
	//line anyForm.qtpl:309
	qw422016.N().S(events)
	//line anyForm.qtpl:309
	qw422016.N().S(` `)
	//line anyForm.qtpl:309
	qw422016.N().S(dataJson)
	//line anyForm.qtpl:309
	qw422016.N().S(`>
        `)
	//line anyForm.qtpl:310
	qw422016.N().S(options)
	//line anyForm.qtpl:310
	qw422016.N().S(`
    </select>
`)
//line anyForm.qtpl:312
}

//line anyForm.qtpl:312
func writerenderSelect(qq422016 qtio422016.Writer, key, options, required, events, dataJson string) {
	//line anyForm.qtpl:312
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line anyForm.qtpl:312
	streamrenderSelect(qw422016, key, options, required, events, dataJson)
	//line anyForm.qtpl:312
	qt422016.ReleaseWriter(qw422016)
//line anyForm.qtpl:312
}

//line anyForm.qtpl:312
func renderSelect(key, options, required, events, dataJson string) string {
	//line anyForm.qtpl:312
	qb422016 := qt422016.AcquireByteBuffer()
	//line anyForm.qtpl:312
	writerenderSelect(qb422016, key, options, required, events, dataJson)
	//line anyForm.qtpl:312
	qs422016 := string(qb422016.B)
	//line anyForm.qtpl:312
	qt422016.ReleaseByteBuffer(qb422016)
	//line anyForm.qtpl:312
	return qs422016
//line anyForm.qtpl:312
}

//line anyForm.qtpl:313
func streamrenderOption(qw422016 *qt422016.Writer, val, title, selected string) {
	//line anyForm.qtpl:313
	qw422016.N().S(`
    <option value="`)
	//line anyForm.qtpl:314
	qw422016.E().S(val)
	//line anyForm.qtpl:314
	qw422016.N().S(`" `)
	//line anyForm.qtpl:314
	qw422016.N().S(selected)
	//line anyForm.qtpl:314
	qw422016.N().S(` >`)
	//line anyForm.qtpl:314
	qw422016.E().S(title)
	//line anyForm.qtpl:314
	qw422016.N().S(`</option>
`)
//line anyForm.qtpl:315
}

//line anyForm.qtpl:315
func writerenderOption(qq422016 qtio422016.Writer, val, title, selected string) {
	//line anyForm.qtpl:315
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line anyForm.qtpl:315
	streamrenderOption(qw422016, val, title, selected)
	//line anyForm.qtpl:315
	qt422016.ReleaseWriter(qw422016)
//line anyForm.qtpl:315
}

//line anyForm.qtpl:315
func renderOption(val, title, selected string) string {
	//line anyForm.qtpl:315
	qb422016 := qt422016.AcquireByteBuffer()
	//line anyForm.qtpl:315
	writerenderOption(qb422016, val, title, selected)
	//line anyForm.qtpl:315
	qs422016 := string(qb422016.B)
	//line anyForm.qtpl:315
	qt422016.ReleaseByteBuffer(qb422016)
	//line anyForm.qtpl:315
	return qs422016
//line anyForm.qtpl:315
}
