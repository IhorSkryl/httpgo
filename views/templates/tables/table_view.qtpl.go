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
<!--table class="table table-striped table-bordered table-hover table-condensed" -->
<div style='margin:0;' class='thead table row-fluid' >
    <div class='tr' style='$style_div' >
        `)
	//line views/templates/tables/table_view.qtpl:20
	var figure string

	//line views/templates/tables/table_view.qtpl:21
	qw422016.N().S(`
            `)
	//line views/templates/tables/table_view.qtpl:22
	for idx, fieldName := range query.columns {
		//line views/templates/tables/table_view.qtpl:22
		qw422016.N().S(`
                `)
		//line views/templates/tables/table_view.qtpl:24
		fieldStruct := query.fields[idx]
		titleFull, titleLabel := fieldStruct.COLUMN_COMMENT, fieldStruct.COLUMN_COMMENT

		if len(titleFull) > 50 {

			titleLabel = titleFull[:strings.LastIndex(titleFull[:51], " ")] + "..."
		}

		if titleLabel == "" {
			titleLabel = fieldName
		}
		if fieldStruct.InputType == "" {
			fieldStruct.InputType = forms.StyleInput(fieldStruct.DATA_TYPE)
		}
		width := forms.GetLengthFromType(fieldStruct.InputType)
		fieldStruct.CSSStyle = fmt.Sprintf("min-width:%dpx;", width)
		query.widthTable += width

		//line views/templates/tables/table_view.qtpl:41
		qw422016.N().S(`
                `)
		//line views/templates/tables/table_view.qtpl:42
		qw422016.N().S(`
                `)
		//line views/templates/tables/table_view.qtpl:43
		if figure != fieldStruct.Figure {
			//line views/templates/tables/table_view.qtpl:43
			qw422016.N().S(`
                    `)
			//line views/templates/tables/table_view.qtpl:44
			if figure > "" {
				//line views/templates/tables/table_view.qtpl:44
				qw422016.N().S(`
                        </div>
                    `)
				//line views/templates/tables/table_view.qtpl:46
			}
			//line views/templates/tables/table_view.qtpl:46
			qw422016.N().S(`
                    `)
			//line views/templates/tables/table_view.qtpl:48
			figure = fieldStruct.Figure

			//line views/templates/tables/table_view.qtpl:49
			qw422016.N().S(`
                    `)
			//line views/templates/tables/table_view.qtpl:50
			if figure > "" {
				//line views/templates/tables/table_view.qtpl:50
				qw422016.N().S(`
                            <div class='td' style="border: 1px inset gray; padding-top: 1px;"> <p class='name_group'> `)
				//line views/templates/tables/table_view.qtpl:51
				qw422016.E().S(figure)
				//line views/templates/tables/table_view.qtpl:51
				qw422016.N().S(`</p>
                    `)
				//line views/templates/tables/table_view.qtpl:52
			}
			//line views/templates/tables/table_view.qtpl:52
			qw422016.N().S(`
                `)
			//line views/templates/tables/table_view.qtpl:53
		}
		//line views/templates/tables/table_view.qtpl:53
		qw422016.N().S(`
                 <div title="`)
		//line views/templates/tables/table_view.qtpl:54
		qw422016.E().S(fieldName)
		//line views/templates/tables/table_view.qtpl:54
		qw422016.N().S(`" class="th" style="`)
		//line views/templates/tables/table_view.qtpl:54
		qw422016.E().S(fieldStruct.CSSStyle)
		//line views/templates/tables/table_view.qtpl:54
		qw422016.N().S(`">
                   <a href="`)
		//line views/templates/tables/table_view.qtpl:55
		qw422016.E().S(query.Href)
		//line views/templates/tables/table_view.qtpl:55
		qw422016.N().S(`/?order=`)
		//line views/templates/tables/table_view.qtpl:55
		qw422016.E().S(fieldName)
		//line views/templates/tables/table_view.qtpl:55
		qw422016.N().S(`" title="`)
		//line views/templates/tables/table_view.qtpl:55
		qw422016.E().S(titleFull)
		//line views/templates/tables/table_view.qtpl:55
		qw422016.N().S(`">`)
		//line views/templates/tables/table_view.qtpl:55
		qw422016.E().S(titleLabel)
		//line views/templates/tables/table_view.qtpl:55
		qw422016.N().S(`</a>
                </div>
            `)
		//line views/templates/tables/table_view.qtpl:57
	}
	//line views/templates/tables/table_view.qtpl:57
	qw422016.N().S(`
            `)
	//line views/templates/tables/table_view.qtpl:58
	if figure > "" {
		//line views/templates/tables/table_view.qtpl:58
		qw422016.N().S(`
                </div>
            `)
		//line views/templates/tables/table_view.qtpl:60
	}
	//line views/templates/tables/table_view.qtpl:60
	qw422016.N().S(`
    </div>
</div>
`)
//line views/templates/tables/table_view.qtpl:63
}

//line views/templates/tables/table_view.qtpl:63
func (query *QueryStruct) writerenderHeadTables(qq422016 qtio422016.Writer) {
	//line views/templates/tables/table_view.qtpl:63
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:63
	query.streamrenderHeadTables(qw422016)
	//line views/templates/tables/table_view.qtpl:63
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:63
}

//line views/templates/tables/table_view.qtpl:63
func (query *QueryStruct) renderHeadTables() string {
	//line views/templates/tables/table_view.qtpl:63
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:63
	query.writerenderHeadTables(qb422016)
	//line views/templates/tables/table_view.qtpl:63
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:63
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:63
	return qs422016
//line views/templates/tables/table_view.qtpl:63
}

//line views/templates/tables/table_view.qtpl:64
func (query *QueryStruct) StreamRenderTable(qw422016 *qt422016.Writer) {
	//line views/templates/tables/table_view.qtpl:64
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:66
	if err := query.beforeRender(); err != nil {
		return
	}

	//line views/templates/tables/table_view.qtpl:69
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:70
	query.streamrenderHeadTables(qw422016)
	//line views/templates/tables/table_view.qtpl:70
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:71
	query.streamrenderBodyTables(qw422016)
	//line views/templates/tables/table_view.qtpl:71
	qw422016.N().S(`
`)
//line views/templates/tables/table_view.qtpl:72
}

//line views/templates/tables/table_view.qtpl:72
func (query *QueryStruct) WriteRenderTable(qq422016 qtio422016.Writer) {
	//line views/templates/tables/table_view.qtpl:72
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:72
	query.StreamRenderTable(qw422016)
	//line views/templates/tables/table_view.qtpl:72
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:72
}

//line views/templates/tables/table_view.qtpl:72
func (query *QueryStruct) RenderTable() string {
	//line views/templates/tables/table_view.qtpl:72
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:72
	query.WriteRenderTable(qb422016)
	//line views/templates/tables/table_view.qtpl:72
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:72
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:72
	return qs422016
//line views/templates/tables/table_view.qtpl:72
}

//line views/templates/tables/table_view.qtpl:73
func StreamShowTable(qw422016 *qt422016.Writer, tableName string, fields forms.FieldsTable, rows *sql.Rows) {
	//line views/templates/tables/table_view.qtpl:73
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:75
	var query QueryStruct
	query.Rows = rows
	query.Href = "/admin/table/" + tableName
	query.HrefEdit = "/admin/row/edit/?table=" + tableName + "&id="
	query.Tables = append(query.Tables, &fields)

	if err := query.beforeRender(); err != nil {
		return
	}

	//line views/templates/tables/table_view.qtpl:84
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:85
	query.streamrenderHeadTables(qw422016)
	//line views/templates/tables/table_view.qtpl:85
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:86
	query.streamrenderBodyTables(qw422016)
	//line views/templates/tables/table_view.qtpl:86
	qw422016.N().S(`
`)
//line views/templates/tables/table_view.qtpl:87
}

//line views/templates/tables/table_view.qtpl:87
func WriteShowTable(qq422016 qtio422016.Writer, tableName string, fields forms.FieldsTable, rows *sql.Rows) {
	//line views/templates/tables/table_view.qtpl:87
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:87
	StreamShowTable(qw422016, tableName, fields, rows)
	//line views/templates/tables/table_view.qtpl:87
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:87
}

//line views/templates/tables/table_view.qtpl:87
func ShowTable(tableName string, fields forms.FieldsTable, rows *sql.Rows) string {
	//line views/templates/tables/table_view.qtpl:87
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:87
	WriteShowTable(qb422016, tableName, fields, rows)
	//line views/templates/tables/table_view.qtpl:87
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:87
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:87
	return qs422016
//line views/templates/tables/table_view.qtpl:87
}

//line views/templates/tables/table_view.qtpl:88
func (query *QueryStruct) streamrenderBodyTables(qw422016 *qt422016.Writer) {
	//line views/templates/tables/table_view.qtpl:88
	qw422016.N().S(`
<div id="table_body" style='width: `)
	//line views/templates/tables/table_view.qtpl:89
	qw422016.N().D(query.widthTable)
	//line views/templates/tables/table_view.qtpl:89
	qw422016.N().S(`px;height:350px;' class='panel-body scroll-pane '>
        `)
	//line views/templates/tables/table_view.qtpl:90
	for query.Rows.Next() {
		//line views/templates/tables/table_view.qtpl:90
		qw422016.N().S(`
            `)
		//line views/templates/tables/table_view.qtpl:92
		if err := query.Rows.Scan(query.row...); err != nil {
			log.Println(err)
			continue
		}

		required, tablePrefix, titleLabel := "", "", ""

		//line views/templates/tables/table_view.qtpl:98
		qw422016.N().S(`
        <div class='tr' style='$style_div'>
            `)
		//line views/templates/tables/table_view.qtpl:100
		for idx, rawField := range query.row {
			//line views/templates/tables/table_view.qtpl:100
			qw422016.N().S(`
                `)
			//line views/templates/tables/table_view.qtpl:102
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

			//line views/templates/tables/table_view.qtpl:122
			qw422016.N().S(`
                `)
			//line views/templates/tables/table_view.qtpl:123
			if idx == 0 {
				//line views/templates/tables/table_view.qtpl:123
				qw422016.N().S(`
                    <div name="`)
				//line views/templates/tables/table_view.qtpl:124
				qw422016.E().S(key)
				//line views/templates/tables/table_view.qtpl:124
				qw422016.N().S(`" class="td" style="`)
				//line views/templates/tables/table_view.qtpl:124
				qw422016.E().S(fieldStruct.CSSStyle)
				//line views/templates/tables/table_view.qtpl:124
				qw422016.N().S(`">
                        <a href="`)
				//line views/templates/tables/table_view.qtpl:125
				qw422016.N().S(query.HrefEdit)
				//line views/templates/tables/table_view.qtpl:125
				qw422016.E().S(val)
				//line views/templates/tables/table_view.qtpl:125
				qw422016.N().S(`" target="content">`)
				//line views/templates/tables/table_view.qtpl:125
				qw422016.E().S(val)
				//line views/templates/tables/table_view.qtpl:125
				qw422016.N().S(`</a>
                    </div>
                    `)
				//line views/templates/tables/table_view.qtpl:128
				tablePrefix = fields.Name
				fields.ID, _ = strconv.Atoi(val)

				//line views/templates/tables/table_view.qtpl:130
				qw422016.N().S(`
                `)
				//line views/templates/tables/table_view.qtpl:131
			} else {
				//line views/templates/tables/table_view.qtpl:131
				qw422016.N().S(`
                `)
				//line views/templates/tables/table_view.qtpl:134
				qw422016.N().S(`
                    <div name="`)
				//line views/templates/tables/table_view.qtpl:135
				qw422016.E().S(key)
				//line views/templates/tables/table_view.qtpl:135
				qw422016.N().S(`" class="td `)
				//line views/templates/tables/table_view.qtpl:135
				qw422016.N().S(required)
				//line views/templates/tables/table_view.qtpl:135
				qw422016.N().S(` field-`)
				//line views/templates/tables/table_view.qtpl:135
				qw422016.E().S(nameInput)
				//line views/templates/tables/table_view.qtpl:135
				qw422016.N().S(` `)
				//line views/templates/tables/table_view.qtpl:135
				qw422016.E().S(fieldStruct.CSSClass)
				//line views/templates/tables/table_view.qtpl:135
				qw422016.N().S(`"
                         style="`)
				//line views/templates/tables/table_view.qtpl:136
				qw422016.E().S(fieldStruct.CSSStyle)
				//line views/templates/tables/table_view.qtpl:136
				qw422016.N().S(`;margin:auto; height:auto;">
                        `)
				//line views/templates/tables/table_view.qtpl:137
				if key == "parent_id" {
					//line views/templates/tables/table_view.qtpl:137
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:138
					fieldStruct.StreamRenderParentSelect(qw422016, fields.Name, key, val, titleLabel, required, events, dataJson)
					//line views/templates/tables/table_view.qtpl:138
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:139
				} else if strings.HasPrefix(key, "id_") {
					//line views/templates/tables/table_view.qtpl:139
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:140
					fieldStruct.StreamRenderForeignSelect(qw422016, tablePrefix, key, val, titleLabel, required, events, dataJson)
					//line views/templates/tables/table_view.qtpl:140
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:141
				} else if strings.HasPrefix(key, "setid_") || strings.HasPrefix(key, "nodeid_") {
					//line views/templates/tables/table_view.qtpl:141
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:142
					fieldStruct.StreamRenderMultiSelect(qw422016, fields, tablePrefix, key, val, "См. ", required)
					//line views/templates/tables/table_view.qtpl:142
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:143
				} else if strings.HasPrefix(key, "tableid_") {
					//line views/templates/tables/table_view.qtpl:143
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:144
					fieldStruct.StreamRenderTable(qw422016, fields, tablePrefix, key, val, "Табл", required)
					//line views/templates/tables/table_view.qtpl:144
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:145
				} else {
					//line views/templates/tables/table_view.qtpl:145
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:146
					switch fieldStruct.DATA_TYPE {
					//line views/templates/tables/table_view.qtpl:147
					case "tinyint":
						//line views/templates/tables/table_view.qtpl:147
						qw422016.N().S(`
                                    `)
						//line views/templates/tables/table_view.qtpl:149
						checked := ""
						if val == "1" {
							checked = "checked"
						}

						//line views/templates/tables/table_view.qtpl:153
						qw422016.N().S(`
                                    `)
						//line views/templates/tables/table_view.qtpl:154
						forms.StreamRenderCheckBox(qw422016, nameInput, "1", titleLabel, 1, checked, required, events, dataJson)
						//line views/templates/tables/table_view.qtpl:154
						qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:155
					case "enum":
						//line views/templates/tables/table_view.qtpl:155
						qw422016.N().S(`
                                `)
						//line views/templates/tables/table_view.qtpl:156
						t := fieldStruct.RenderEnum(nameInput, val, required, events, dataJson)

						//line views/templates/tables/table_view.qtpl:156
						qw422016.N().S(`
                                `)
						//line views/templates/tables/table_view.qtpl:157
						qw422016.N().S(t)
						//line views/templates/tables/table_view.qtpl:157
						qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:158
					case "set":
						//line views/templates/tables/table_view.qtpl:158
						qw422016.N().S(`
                                `)
						//line views/templates/tables/table_view.qtpl:159
						t := fieldStruct.RenderSet(nameInput, val, required, events, dataJson)

						//line views/templates/tables/table_view.qtpl:159
						qw422016.N().S(`
                                `)
						//line views/templates/tables/table_view.qtpl:160
						qw422016.N().S(t)
						//line views/templates/tables/table_view.qtpl:160
						qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:161
					case "text":
						//line views/templates/tables/table_view.qtpl:161
						qw422016.N().S(`
                                <textarea id="`)
						//line views/templates/tables/table_view.qtpl:162
						qw422016.E().S(key)
						//line views/templates/tables/table_view.qtpl:162
						qw422016.N().S(`" name="`)
						//line views/templates/tables/table_view.qtpl:162
						qw422016.E().S(nameInput)
						//line views/templates/tables/table_view.qtpl:162
						qw422016.N().S(`" class="controls"  `)
						//line views/templates/tables/table_view.qtpl:162
						qw422016.N().S(events)
						//line views/templates/tables/table_view.qtpl:162
						qw422016.N().S(` `)
						//line views/templates/tables/table_view.qtpl:162
						qw422016.N().S(dataJson)
						//line views/templates/tables/table_view.qtpl:162
						qw422016.N().S(`>
                                    `)
						//line views/templates/tables/table_view.qtpl:163
						qw422016.N().S(val)
						//line views/templates/tables/table_view.qtpl:163
						qw422016.N().S(`
                                </textarea>
                            `)
					//line views/templates/tables/table_view.qtpl:165
					case "blob":
						//line views/templates/tables/table_view.qtpl:165
						qw422016.N().S(`
                                <image src="/images/expand_hover.png" />
                            `)
					//line views/templates/tables/table_view.qtpl:167
					default:
						//line views/templates/tables/table_view.qtpl:167
						qw422016.N().S(`
                                `)
						//line views/templates/tables/table_view.qtpl:168
						if fieldStruct.LinkTD > "" {
							//line views/templates/tables/table_view.qtpl:168
							qw422016.N().S(`
                                 <a href="`)
							//line views/templates/tables/table_view.qtpl:169
							qw422016.N().S(fieldStruct.LinkTD)
							//line views/templates/tables/table_view.qtpl:169
							qw422016.N().S(`?id=`)
							//line views/templates/tables/table_view.qtpl:169
							qw422016.N().D(fields.ID)
							//line views/templates/tables/table_view.qtpl:169
							qw422016.N().S(`" target="_blank" class="referal"> !! </a>
                                 `)
							//line views/templates/tables/table_view.qtpl:170
						}
						//line views/templates/tables/table_view.qtpl:170
						qw422016.N().S(`
                                `)
						//line views/templates/tables/table_view.qtpl:171
						qw422016.E().S(val)
						//line views/templates/tables/table_view.qtpl:171
						qw422016.N().S(`
                            `)
						//line views/templates/tables/table_view.qtpl:172
					}
					//line views/templates/tables/table_view.qtpl:172
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:173
				}
				//line views/templates/tables/table_view.qtpl:173
				qw422016.N().S(`
                    </div>
                `)
				//line views/templates/tables/table_view.qtpl:175
			}
			//line views/templates/tables/table_view.qtpl:175
			qw422016.N().S(`
            `)
			//line views/templates/tables/table_view.qtpl:176
		}
		//line views/templates/tables/table_view.qtpl:176
		qw422016.N().S(`
        </div>
        `)
		//line views/templates/tables/table_view.qtpl:178
	}
	//line views/templates/tables/table_view.qtpl:178
	qw422016.N().S(`
    </div>

<!--/table-->

`)
//line views/templates/tables/table_view.qtpl:183
}

//line views/templates/tables/table_view.qtpl:183
func (query *QueryStruct) writerenderBodyTables(qq422016 qtio422016.Writer) {
	//line views/templates/tables/table_view.qtpl:183
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:183
	query.streamrenderBodyTables(qw422016)
	//line views/templates/tables/table_view.qtpl:183
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:183
}

//line views/templates/tables/table_view.qtpl:183
func (query *QueryStruct) renderBodyTables() string {
	//line views/templates/tables/table_view.qtpl:183
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:183
	query.writerenderBodyTables(qb422016)
	//line views/templates/tables/table_view.qtpl:183
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:183
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:183
	return qs422016
//line views/templates/tables/table_view.qtpl:183
}
