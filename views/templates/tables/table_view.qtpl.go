// This file is automatically generated by qtc from "table_view.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line views/templates/tables/table_view.qtpl:1
package tables

//line views/templates/tables/table_view.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// All the text outside function templates is treated as comments,
// i.e. it is just ignored by quicktemplate compiler (`qtc`). It is for humans.
//
// показ табличных данных.

//line views/templates/tables/table_view.qtpl:5
import (
	"database/sql"
	"fmt"
	"github.com/ruslanBik4/httpgo/views/templates/forms"
	"log"
	"strconv"
	"strings"
)

//line views/templates/tables/table_view.qtpl:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/tables/table_view.qtpl:14
func (query *QueryStruct) streamrenderHeadTables(qw422016 *qt422016.Writer) {
	//line views/templates/tables/table_view.qtpl:14
	qw422016.N().S(`
<link href="/tables.css" rel="stylesheet">
<table class="table table-striped table-bordered table-hover table-condensed">
    <thead>
        <tr>
            `)
	//line views/templates/tables/table_view.qtpl:19
	for idx, fieldName := range query.columns {
		//line views/templates/tables/table_view.qtpl:19
		qw422016.N().S(`
                <td>`)
		//line views/templates/tables/table_view.qtpl:21
		fieldStruct := query.fields[idx]
		titleFull, titleLabel := fieldStruct.COLUMN_COMMENT, fieldStruct.COLUMN_COMMENT

		if len(titleFull) > 50 {

			titleLabel = titleFull[:strings.LastIndex(titleFull[:51], " ")] + "..."
		}

		if titleLabel == "" {
			titleLabel = fieldName
		}

		//line views/templates/tables/table_view.qtpl:33
		qw422016.N().S(`
                    <a href="`)
		//line views/templates/tables/table_view.qtpl:34
		qw422016.E().S(query.Href)
		//line views/templates/tables/table_view.qtpl:34
		qw422016.N().S(`/?order=`)
		//line views/templates/tables/table_view.qtpl:34
		qw422016.E().S(fieldName)
		//line views/templates/tables/table_view.qtpl:34
		qw422016.N().S(`" title="`)
		//line views/templates/tables/table_view.qtpl:34
		qw422016.E().S(titleFull)
		//line views/templates/tables/table_view.qtpl:34
		qw422016.N().S(`">`)
		//line views/templates/tables/table_view.qtpl:34
		qw422016.E().S(titleLabel)
		//line views/templates/tables/table_view.qtpl:34
		qw422016.N().S(`</a>
                </td>
            `)
		//line views/templates/tables/table_view.qtpl:36
	}
	//line views/templates/tables/table_view.qtpl:36
	qw422016.N().S(`
        </tr>
    </thead>
`)
//line views/templates/tables/table_view.qtpl:39
}

//line views/templates/tables/table_view.qtpl:39
func (query *QueryStruct) writerenderHeadTables(qq422016 qtio422016.Writer) {
	//line views/templates/tables/table_view.qtpl:39
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:39
	query.streamrenderHeadTables(qw422016)
	//line views/templates/tables/table_view.qtpl:39
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:39
}

//line views/templates/tables/table_view.qtpl:39
func (query *QueryStruct) renderHeadTables() string {
	//line views/templates/tables/table_view.qtpl:39
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:39
	query.writerenderHeadTables(qb422016)
	//line views/templates/tables/table_view.qtpl:39
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:39
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:39
	return qs422016
//line views/templates/tables/table_view.qtpl:39
}

//line views/templates/tables/table_view.qtpl:40
func (query *QueryStruct) StreamRenderTable(qw422016 *qt422016.Writer) {
	//line views/templates/tables/table_view.qtpl:40
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:42
	if err := query.beforeRender(); err != nil {
		return
	}

	//line views/templates/tables/table_view.qtpl:45
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:46
	query.streamrenderHeadTables(qw422016)
	//line views/templates/tables/table_view.qtpl:46
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:47
	query.streamrenderBodyTables(qw422016)
	//line views/templates/tables/table_view.qtpl:47
	qw422016.N().S(`
`)
//line views/templates/tables/table_view.qtpl:48
}

//line views/templates/tables/table_view.qtpl:48
func (query *QueryStruct) WriteRenderTable(qq422016 qtio422016.Writer) {
	//line views/templates/tables/table_view.qtpl:48
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:48
	query.StreamRenderTable(qw422016)
	//line views/templates/tables/table_view.qtpl:48
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:48
}

//line views/templates/tables/table_view.qtpl:48
func (query *QueryStruct) RenderTable() string {
	//line views/templates/tables/table_view.qtpl:48
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:48
	query.WriteRenderTable(qb422016)
	//line views/templates/tables/table_view.qtpl:48
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:48
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:48
	return qs422016
//line views/templates/tables/table_view.qtpl:48
}

//line views/templates/tables/table_view.qtpl:49
func StreamShowTable(qw422016 *qt422016.Writer, tableName string, fields forms.FieldsTable, rows *sql.Rows) {
	//line views/templates/tables/table_view.qtpl:49
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:51
	var query QueryStruct
	query.Rows = rows
	query.Href = "/admin/table/" + tableName
	query.HrefEdit = "/admin/row/edit/?table=" + tableName + "&id="
	query.Tables = append(query.Tables, &fields)

	if err := query.beforeRender(); err != nil {
		return
	}

	//line views/templates/tables/table_view.qtpl:60
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:61
	query.streamrenderHeadTables(qw422016)
	//line views/templates/tables/table_view.qtpl:61
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:62
	query.streamrenderBodyTables(qw422016)
	//line views/templates/tables/table_view.qtpl:62
	qw422016.N().S(`
`)
//line views/templates/tables/table_view.qtpl:63
}

//line views/templates/tables/table_view.qtpl:63
func WriteShowTable(qq422016 qtio422016.Writer, tableName string, fields forms.FieldsTable, rows *sql.Rows) {
	//line views/templates/tables/table_view.qtpl:63
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:63
	StreamShowTable(qw422016, tableName, fields, rows)
	//line views/templates/tables/table_view.qtpl:63
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:63
}

//line views/templates/tables/table_view.qtpl:63
func ShowTable(tableName string, fields forms.FieldsTable, rows *sql.Rows) string {
	//line views/templates/tables/table_view.qtpl:63
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:63
	WriteShowTable(qb422016, tableName, fields, rows)
	//line views/templates/tables/table_view.qtpl:63
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:63
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:63
	return qs422016
//line views/templates/tables/table_view.qtpl:63
}

//line views/templates/tables/table_view.qtpl:64
func (query *QueryStruct) streamrenderBodyTables(qw422016 *qt422016.Writer) {
	//line views/templates/tables/table_view.qtpl:64
	qw422016.N().S(`
    <tbody>
        `)
	//line views/templates/tables/table_view.qtpl:66
	for query.Rows.Next() {
		//line views/templates/tables/table_view.qtpl:66
		qw422016.N().S(`
            `)
		//line views/templates/tables/table_view.qtpl:68
		if err := query.Rows.Scan(query.row...); err != nil {
			log.Println(err)
			continue
		}

		required, tablePrefix, titleLabel := "", "", ""

		//line views/templates/tables/table_view.qtpl:74
		qw422016.N().S(`
        <tr>
            `)
		//line views/templates/tables/table_view.qtpl:76
		for idx, rawField := range query.row {
			//line views/templates/tables/table_view.qtpl:76
			qw422016.N().S(`
                `)
			//line views/templates/tables/table_view.qtpl:78
			var fieldStruct *forms.FieldStructure

			switch rawField.(type) {
			case *forms.FieldStructure:
				fieldStruct = rawField.(*forms.FieldStructure)
			default:
				fieldStruct = query.fields[idx]
			}
			key := fieldStruct.COLUMN_NAME
			val := fieldStruct.Value
			fields := fieldStruct.Table
			nameInput := tablePrefix + key
			dataJson := ""
			if dataJson > "" {
				dataJson = fmt.Sprintf(`data-names="{% s}"`, dataJson)
			}
			events := ""
			for name, funcName := range fieldStruct.Events {
				events += fmt.Sprintf(`%s="return %s;"`, name, funcName)
			}

			//line views/templates/tables/table_view.qtpl:98
			qw422016.N().S(`
                `)
			//line views/templates/tables/table_view.qtpl:99
			if idx == 0 {
				//line views/templates/tables/table_view.qtpl:99
				qw422016.N().S(`
                    <td><a href="`)
				//line views/templates/tables/table_view.qtpl:100
				qw422016.N().S(query.HrefEdit)
				//line views/templates/tables/table_view.qtpl:100
				qw422016.E().S(val)
				//line views/templates/tables/table_view.qtpl:100
				qw422016.N().S(`" target="content">`)
				//line views/templates/tables/table_view.qtpl:100
				qw422016.E().S(val)
				//line views/templates/tables/table_view.qtpl:100
				qw422016.N().S(`</a>
                    </td>
                    `)
				//line views/templates/tables/table_view.qtpl:103
				tablePrefix = fields.Name
				fields.ID, _ = strconv.Atoi(val)

				//line views/templates/tables/table_view.qtpl:105
				qw422016.N().S(`
                `)
				//line views/templates/tables/table_view.qtpl:106
			} else {
				//line views/templates/tables/table_view.qtpl:106
				qw422016.N().S(`
                `)
				//line views/templates/tables/table_view.qtpl:109
				qw422016.N().S(`
                    <td class="`)
				//line views/templates/tables/table_view.qtpl:110
				qw422016.N().S(required)
				//line views/templates/tables/table_view.qtpl:110
				qw422016.N().S(` field-`)
				//line views/templates/tables/table_view.qtpl:110
				qw422016.E().S(nameInput)
				//line views/templates/tables/table_view.qtpl:110
				qw422016.N().S(` `)
				//line views/templates/tables/table_view.qtpl:110
				qw422016.E().S(fieldStruct.CSSClass)
				//line views/templates/tables/table_view.qtpl:110
				qw422016.N().S(`">
                        `)
				//line views/templates/tables/table_view.qtpl:111
				if key == "parent_id" {
					//line views/templates/tables/table_view.qtpl:111
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:112
					fieldStruct.StreamRenderParentSelect(qw422016, fields.Name, key, val, titleLabel, required, events, dataJson)
					//line views/templates/tables/table_view.qtpl:112
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:113
				} else if strings.HasPrefix(key, "id_") {
					//line views/templates/tables/table_view.qtpl:113
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:114
					fieldStruct.StreamRenderForeignSelect(qw422016, tablePrefix, key, val, titleLabel, required, events, dataJson)
					//line views/templates/tables/table_view.qtpl:114
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:115
				} else if strings.HasPrefix(key, "setid_") {
					//line views/templates/tables/table_view.qtpl:115
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:116
					fieldStruct.StreamRenderMultiSelect(qw422016, fields, tablePrefix, key, val, "См. ", required)
					//line views/templates/tables/table_view.qtpl:116
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:117
				} else if strings.HasPrefix(key, "tableid_") {
					//line views/templates/tables/table_view.qtpl:117
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:118
					fieldStruct.StreamRenderTable(qw422016, fields, tablePrefix, key, val, "Табл", required)
					//line views/templates/tables/table_view.qtpl:118
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:119
				} else {
					//line views/templates/tables/table_view.qtpl:119
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:120
					switch fieldStruct.DATA_TYPE {
					//line views/templates/tables/table_view.qtpl:121
					case "tinyint":
						//line views/templates/tables/table_view.qtpl:121
						qw422016.N().S(`
                                    `)
						//line views/templates/tables/table_view.qtpl:123
						checked := ""
						if val == "1" {
							checked = "checked"
						}

						//line views/templates/tables/table_view.qtpl:127
						qw422016.N().S(`
                                    `)
						//line views/templates/tables/table_view.qtpl:128
						forms.StreamRenderCheckBox(qw422016, nameInput, "1", titleLabel, 1, checked, required, events, dataJson)
						//line views/templates/tables/table_view.qtpl:128
						qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:129
					case "enum":
						//line views/templates/tables/table_view.qtpl:129
						qw422016.N().S(`
                            `)
						//line views/templates/tables/table_view.qtpl:130
						t := fieldStruct.RenderEnum(nameInput, val, required, events, dataJson)

						//line views/templates/tables/table_view.qtpl:130
						qw422016.N().S(`
                            `)
						//line views/templates/tables/table_view.qtpl:131
						qw422016.N().S(t)
						//line views/templates/tables/table_view.qtpl:131
						qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:132
					case "set":
						//line views/templates/tables/table_view.qtpl:132
						qw422016.N().S(`
                            `)
						//line views/templates/tables/table_view.qtpl:133
						t := fieldStruct.RenderSet(nameInput, val, required, events, dataJson)

						//line views/templates/tables/table_view.qtpl:133
						qw422016.N().S(`
                            `)
						//line views/templates/tables/table_view.qtpl:134
						qw422016.N().S(t)
						//line views/templates/tables/table_view.qtpl:134
						qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:135
					case "blob", "text":
						//line views/templates/tables/table_view.qtpl:135
						qw422016.N().S(`
                            <textarea id="`)
						//line views/templates/tables/table_view.qtpl:136
						qw422016.E().S(key)
						//line views/templates/tables/table_view.qtpl:136
						qw422016.N().S(`" name="`)
						//line views/templates/tables/table_view.qtpl:136
						qw422016.E().S(nameInput)
						//line views/templates/tables/table_view.qtpl:136
						qw422016.N().S(`" class="controls"  `)
						//line views/templates/tables/table_view.qtpl:136
						qw422016.N().S(events)
						//line views/templates/tables/table_view.qtpl:136
						qw422016.N().S(` `)
						//line views/templates/tables/table_view.qtpl:136
						qw422016.N().S(dataJson)
						//line views/templates/tables/table_view.qtpl:136
						qw422016.N().S(`>
                                `)
						//line views/templates/tables/table_view.qtpl:137
						qw422016.N().S(val)
						//line views/templates/tables/table_view.qtpl:137
						qw422016.N().S(`
                            </textarea>
                            `)
					//line views/templates/tables/table_view.qtpl:139
					default:
						//line views/templates/tables/table_view.qtpl:139
						qw422016.N().S(`
                                `)
						//line views/templates/tables/table_view.qtpl:140
						qw422016.E().S(val)
						//line views/templates/tables/table_view.qtpl:140
						qw422016.N().S(`
                            `)
						//line views/templates/tables/table_view.qtpl:141
					}
					//line views/templates/tables/table_view.qtpl:141
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:142
				}
				//line views/templates/tables/table_view.qtpl:142
				qw422016.N().S(`
                    </td>
                `)
				//line views/templates/tables/table_view.qtpl:144
			}
			//line views/templates/tables/table_view.qtpl:144
			qw422016.N().S(`
            `)
			//line views/templates/tables/table_view.qtpl:145
		}
		//line views/templates/tables/table_view.qtpl:145
		qw422016.N().S(`
        </tr>
        `)
		//line views/templates/tables/table_view.qtpl:147
	}
	//line views/templates/tables/table_view.qtpl:147
	qw422016.N().S(`
    </tbody>

</table>

`)
//line views/templates/tables/table_view.qtpl:152
}

//line views/templates/tables/table_view.qtpl:152
func (query *QueryStruct) writerenderBodyTables(qq422016 qtio422016.Writer) {
	//line views/templates/tables/table_view.qtpl:152
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:152
	query.streamrenderBodyTables(qw422016)
	//line views/templates/tables/table_view.qtpl:152
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:152
}

//line views/templates/tables/table_view.qtpl:152
func (query *QueryStruct) renderBodyTables() string {
	//line views/templates/tables/table_view.qtpl:152
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:152
	query.writerenderBodyTables(qb422016)
	//line views/templates/tables/table_view.qtpl:152
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:152
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:152
	return qs422016
//line views/templates/tables/table_view.qtpl:152
}
