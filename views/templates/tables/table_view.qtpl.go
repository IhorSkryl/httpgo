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
	"github.com/ruslanBik4/httpgo/views/templates/forms"
	"log"
	"strconv"
	"strings"
)

//line views/templates/tables/table_view.qtpl:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/tables/table_view.qtpl:13
func (query *QueryStruct) streamrenderHeadTables(qw422016 *qt422016.Writer) {
	//line views/templates/tables/table_view.qtpl:13
	qw422016.N().S(`
<link href="/tables.css" rel="stylesheet">
<table class="table table-striped table-bordered table-hover table-condensed">
    <thead>
        <tr>
            `)
	//line views/templates/tables/table_view.qtpl:18
	for idx, fieldName := range query.columns {
		//line views/templates/tables/table_view.qtpl:18
		qw422016.N().S(`
                <td>`)
		//line views/templates/tables/table_view.qtpl:20
		fieldStruct := query.fields[idx]
		titleFull, titleLabel := fieldStruct.COLUMN_COMMENT, fieldStruct.COLUMN_COMMENT

		if len(titleFull) > 50 {

			titleLabel = titleFull[:strings.LastIndex(titleFull[:51], " ")] + "..."
		}

		if titleLabel == "" {
			titleLabel = fieldName
		}

		//line views/templates/tables/table_view.qtpl:32
		qw422016.N().S(`
                    <a href="`)
		//line views/templates/tables/table_view.qtpl:33
		qw422016.E().S(query.Href)
		//line views/templates/tables/table_view.qtpl:33
		qw422016.N().S(`/?order=`)
		//line views/templates/tables/table_view.qtpl:33
		qw422016.E().S(fieldName)
		//line views/templates/tables/table_view.qtpl:33
		qw422016.N().S(`" title="`)
		//line views/templates/tables/table_view.qtpl:33
		qw422016.E().S(titleFull)
		//line views/templates/tables/table_view.qtpl:33
		qw422016.N().S(`">`)
		//line views/templates/tables/table_view.qtpl:33
		qw422016.E().S(titleLabel)
		//line views/templates/tables/table_view.qtpl:33
		qw422016.N().S(`</a>
                </td>
            `)
		//line views/templates/tables/table_view.qtpl:35
	}
	//line views/templates/tables/table_view.qtpl:35
	qw422016.N().S(`
        </tr>
    </thead>
`)
//line views/templates/tables/table_view.qtpl:38
}

//line views/templates/tables/table_view.qtpl:38
func (query *QueryStruct) writerenderHeadTables(qq422016 qtio422016.Writer) {
	//line views/templates/tables/table_view.qtpl:38
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:38
	query.streamrenderHeadTables(qw422016)
	//line views/templates/tables/table_view.qtpl:38
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:38
}

//line views/templates/tables/table_view.qtpl:38
func (query *QueryStruct) renderHeadTables() string {
	//line views/templates/tables/table_view.qtpl:38
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:38
	query.writerenderHeadTables(qb422016)
	//line views/templates/tables/table_view.qtpl:38
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:38
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:38
	return qs422016
//line views/templates/tables/table_view.qtpl:38
}

//line views/templates/tables/table_view.qtpl:39
func (query *QueryStruct) StreamRenderTable(qw422016 *qt422016.Writer) {
	//line views/templates/tables/table_view.qtpl:39
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:41
	if err := query.beforeRender(); err != nil {
		return
	}

	//line views/templates/tables/table_view.qtpl:44
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:45
	query.streamrenderHeadTables(qw422016)
	//line views/templates/tables/table_view.qtpl:45
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:46
	query.streamrenderBodyTables(qw422016)
	//line views/templates/tables/table_view.qtpl:46
	qw422016.N().S(`
`)
//line views/templates/tables/table_view.qtpl:47
}

//line views/templates/tables/table_view.qtpl:47
func (query *QueryStruct) WriteRenderTable(qq422016 qtio422016.Writer) {
	//line views/templates/tables/table_view.qtpl:47
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:47
	query.StreamRenderTable(qw422016)
	//line views/templates/tables/table_view.qtpl:47
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:47
}

//line views/templates/tables/table_view.qtpl:47
func (query *QueryStruct) RenderTable() string {
	//line views/templates/tables/table_view.qtpl:47
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:47
	query.WriteRenderTable(qb422016)
	//line views/templates/tables/table_view.qtpl:47
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:47
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:47
	return qs422016
//line views/templates/tables/table_view.qtpl:47
}

//line views/templates/tables/table_view.qtpl:48
func StreamShowTable(qw422016 *qt422016.Writer, tableName string, fields forms.FieldsTable, rows *sql.Rows) {
	//line views/templates/tables/table_view.qtpl:48
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:50
	var query QueryStruct
	query.Rows = rows
	query.Href = "/admin/table/" + tableName
	query.HrefEdit = "/admin/row/edit/?table=" + tableName + "&id="
	query.Tables = append(query.Tables, &fields)

	if err := query.beforeRender(); err != nil {
		return
	}

	//line views/templates/tables/table_view.qtpl:59
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:60
	query.streamrenderHeadTables(qw422016)
	//line views/templates/tables/table_view.qtpl:60
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:61
	query.streamrenderBodyTables(qw422016)
	//line views/templates/tables/table_view.qtpl:61
	qw422016.N().S(`
`)
//line views/templates/tables/table_view.qtpl:62
}

//line views/templates/tables/table_view.qtpl:62
func WriteShowTable(qq422016 qtio422016.Writer, tableName string, fields forms.FieldsTable, rows *sql.Rows) {
	//line views/templates/tables/table_view.qtpl:62
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:62
	StreamShowTable(qw422016, tableName, fields, rows)
	//line views/templates/tables/table_view.qtpl:62
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:62
}

//line views/templates/tables/table_view.qtpl:62
func ShowTable(tableName string, fields forms.FieldsTable, rows *sql.Rows) string {
	//line views/templates/tables/table_view.qtpl:62
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:62
	WriteShowTable(qb422016, tableName, fields, rows)
	//line views/templates/tables/table_view.qtpl:62
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:62
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:62
	return qs422016
//line views/templates/tables/table_view.qtpl:62
}

//line views/templates/tables/table_view.qtpl:63
func (query *QueryStruct) streamrenderBodyTables(qw422016 *qt422016.Writer) {
	//line views/templates/tables/table_view.qtpl:63
	qw422016.N().S(`
    <tbody>
        `)
	//line views/templates/tables/table_view.qtpl:65
	for query.Rows.Next() {
		//line views/templates/tables/table_view.qtpl:65
		qw422016.N().S(`
            `)
		//line views/templates/tables/table_view.qtpl:67
		if err := query.Rows.Scan(query.row...); err != nil {
			log.Println(err)
			continue
		}

		required, tablePrefix, titleLabel := "", "", ""

		//line views/templates/tables/table_view.qtpl:73
		qw422016.N().S(`
        <tr>
            `)
		//line views/templates/tables/table_view.qtpl:75
		for idx, rawField := range query.row {
			//line views/templates/tables/table_view.qtpl:75
			qw422016.N().S(`
                `)
			//line views/templates/tables/table_view.qtpl:77
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

			//line views/templates/tables/table_view.qtpl:88
			qw422016.N().S(`
                `)
			//line views/templates/tables/table_view.qtpl:89
			if idx == 0 {
				//line views/templates/tables/table_view.qtpl:89
				qw422016.N().S(`
                    <td><a href="`)
				//line views/templates/tables/table_view.qtpl:90
				qw422016.N().S(query.HrefEdit)
				//line views/templates/tables/table_view.qtpl:90
				qw422016.E().S(val)
				//line views/templates/tables/table_view.qtpl:90
				qw422016.N().S(`" target="content">`)
				//line views/templates/tables/table_view.qtpl:90
				qw422016.E().S(val)
				//line views/templates/tables/table_view.qtpl:90
				qw422016.N().S(`</a>
                    </td>
                    `)
				//line views/templates/tables/table_view.qtpl:93
				tablePrefix = fields.Name
				fields.ID, _ = strconv.Atoi(val)

				//line views/templates/tables/table_view.qtpl:95
				qw422016.N().S(`
                `)
				//line views/templates/tables/table_view.qtpl:96
			} else {
				//line views/templates/tables/table_view.qtpl:96
				qw422016.N().S(`
                    <td>
                        `)
				//line views/templates/tables/table_view.qtpl:98
				if key == "parent_id" {
					//line views/templates/tables/table_view.qtpl:98
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:99
					fieldStruct.StreamRenderParentSelect(qw422016, fields.Name, key, val, titleLabel, required)
					//line views/templates/tables/table_view.qtpl:99
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:100
				} else if strings.HasPrefix(key, "id_") {
					//line views/templates/tables/table_view.qtpl:100
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:101
					fieldStruct.StreamRenderForeignSelect(qw422016, tablePrefix, key, val, titleLabel, required)
					//line views/templates/tables/table_view.qtpl:101
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:102
				} else if strings.HasPrefix(key, "setid_") {
					//line views/templates/tables/table_view.qtpl:102
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:103
					fieldStruct.StreamRenderMultiSelect(qw422016, fields, tablePrefix, key, val, "См. ", required)
					//line views/templates/tables/table_view.qtpl:103
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:104
				} else if strings.HasPrefix(key, "tableid_") {
					//line views/templates/tables/table_view.qtpl:104
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:105
					fieldStruct.StreamRenderTable(qw422016, fields, tablePrefix, key, val, "Табл", required)
					//line views/templates/tables/table_view.qtpl:105
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:106
				} else {
					//line views/templates/tables/table_view.qtpl:106
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:107
					switch fieldStruct.DATA_TYPE {
					//line views/templates/tables/table_view.qtpl:108
					case "tinyint":
						//line views/templates/tables/table_view.qtpl:108
						qw422016.N().S(`
                                    `)
						//line views/templates/tables/table_view.qtpl:110
						checked := ""
						if val == "1" {
							checked = "checked"
						}

						//line views/templates/tables/table_view.qtpl:114
						qw422016.N().S(`
                                    `)
						//line views/templates/tables/table_view.qtpl:115
						forms.StreamRenderCheckBox(qw422016, tablePrefix+key, "1", titleLabel, 1, checked, "", "")
						//line views/templates/tables/table_view.qtpl:115
						qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:116
					default:
						//line views/templates/tables/table_view.qtpl:116
						qw422016.N().S(`
                                `)
						//line views/templates/tables/table_view.qtpl:117
						qw422016.E().S(val)
						//line views/templates/tables/table_view.qtpl:117
						qw422016.N().S(`
                            `)
						//line views/templates/tables/table_view.qtpl:118
					}
					//line views/templates/tables/table_view.qtpl:118
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:119
				}
				//line views/templates/tables/table_view.qtpl:119
				qw422016.N().S(`
                    </td>
                `)
				//line views/templates/tables/table_view.qtpl:121
			}
			//line views/templates/tables/table_view.qtpl:121
			qw422016.N().S(`
            `)
			//line views/templates/tables/table_view.qtpl:122
		}
		//line views/templates/tables/table_view.qtpl:122
		qw422016.N().S(`
        </tr>
        `)
		//line views/templates/tables/table_view.qtpl:124
	}
	//line views/templates/tables/table_view.qtpl:124
	qw422016.N().S(`
    </tbody>

</table>

`)
//line views/templates/tables/table_view.qtpl:129
}

//line views/templates/tables/table_view.qtpl:129
func (query *QueryStruct) writerenderBodyTables(qq422016 qtio422016.Writer) {
	//line views/templates/tables/table_view.qtpl:129
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:129
	query.streamrenderBodyTables(qw422016)
	//line views/templates/tables/table_view.qtpl:129
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:129
}

//line views/templates/tables/table_view.qtpl:129
func (query *QueryStruct) renderBodyTables() string {
	//line views/templates/tables/table_view.qtpl:129
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:129
	query.writerenderBodyTables(qb422016)
	//line views/templates/tables/table_view.qtpl:129
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:129
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:129
	return qs422016
//line views/templates/tables/table_view.qtpl:129
}
