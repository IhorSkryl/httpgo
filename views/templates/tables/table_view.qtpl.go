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
<div class="table">
<div class='thead' >
    <div class='tr' >
        `)
	//line views/templates/tables/table_view.qtpl:20
	var figure, filterFields string
	if query.Order == "" {
		query.Order = "id"
	}

	//line views/templates/tables/table_view.qtpl:24
	qw422016.N().S(`
            `)
	//line views/templates/tables/table_view.qtpl:25
	for idx, fieldName := range query.columns {
		//line views/templates/tables/table_view.qtpl:25
		qw422016.N().S(`
                `)
		//line views/templates/tables/table_view.qtpl:27
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
		fieldStruct.CSSStyle = fmt.Sprintf("width:%dpx;", width)
		query.widthTable += width
		filterFields += fmt.Sprintf(`<div title="%s" class="td" style="%s"><input/></div>`, fieldName, fieldStruct.CSSStyle)

		//line views/templates/tables/table_view.qtpl:45
		qw422016.N().S(`
                `)
		//line views/templates/tables/table_view.qtpl:46
		qw422016.N().S(`
                `)
		//line views/templates/tables/table_view.qtpl:47
		if figure != fieldStruct.Figure {
			//line views/templates/tables/table_view.qtpl:47
			qw422016.N().S(`
                    `)
			//line views/templates/tables/table_view.qtpl:48
			if figure > "" {
				//line views/templates/tables/table_view.qtpl:48
				qw422016.N().S(`
                        </div>
                    `)
				//line views/templates/tables/table_view.qtpl:50
			}
			//line views/templates/tables/table_view.qtpl:50
			qw422016.N().S(`
                    `)
			//line views/templates/tables/table_view.qtpl:52
			figure = fieldStruct.Figure

			//line views/templates/tables/table_view.qtpl:53
			qw422016.N().S(`
                    `)
			//line views/templates/tables/table_view.qtpl:54
			if figure > "" {
				//line views/templates/tables/table_view.qtpl:54
				qw422016.N().S(`
                            <div class='td-outer' style="border: 1px inset gray;">
                                <div class='th' style="border-bottom: 1px outset gray; padding-top: 1px;"> `)
				//line views/templates/tables/table_view.qtpl:56
				qw422016.E().S(figure)
				//line views/templates/tables/table_view.qtpl:56
				qw422016.N().S(`</div>
                    `)
				//line views/templates/tables/table_view.qtpl:57
			}
			//line views/templates/tables/table_view.qtpl:57
			qw422016.N().S(`
                `)
			//line views/templates/tables/table_view.qtpl:58
		}
		//line views/templates/tables/table_view.qtpl:58
		qw422016.N().S(`
                 <div title="`)
		//line views/templates/tables/table_view.qtpl:59
		qw422016.E().S(fieldName)
		//line views/templates/tables/table_view.qtpl:59
		qw422016.N().S(`" class="td`)
		//line views/templates/tables/table_view.qtpl:59
		if fieldName == query.Order {
			//line views/templates/tables/table_view.qtpl:59
			qw422016.N().S(` td-order`)
			//line views/templates/tables/table_view.qtpl:59
		}
		//line views/templates/tables/table_view.qtpl:59
		qw422016.N().S(`"
                 style="`)
		//line views/templates/tables/table_view.qtpl:60
		qw422016.E().S(fieldStruct.CSSStyle)
		//line views/templates/tables/table_view.qtpl:60
		qw422016.N().S(`">
                   <a href="`)
		//line views/templates/tables/table_view.qtpl:61
		qw422016.E().S(query.Href)
		//line views/templates/tables/table_view.qtpl:61
		qw422016.N().S(`/?order=`)
		//line views/templates/tables/table_view.qtpl:61
		qw422016.E().S(fieldName)
		//line views/templates/tables/table_view.qtpl:61
		qw422016.N().S(`" title="`)
		//line views/templates/tables/table_view.qtpl:61
		qw422016.E().S(titleFull)
		//line views/templates/tables/table_view.qtpl:61
		qw422016.N().S(`">`)
		//line views/templates/tables/table_view.qtpl:61
		qw422016.E().S(titleLabel)
		//line views/templates/tables/table_view.qtpl:61
		qw422016.N().S(`</a>
                </div>
            `)
		//line views/templates/tables/table_view.qtpl:63
	}
	//line views/templates/tables/table_view.qtpl:63
	qw422016.N().S(`
            `)
	//line views/templates/tables/table_view.qtpl:64
	if figure > "" {
		//line views/templates/tables/table_view.qtpl:64
		qw422016.N().S(`
                </div>
            `)
		//line views/templates/tables/table_view.qtpl:66
	}
	//line views/templates/tables/table_view.qtpl:66
	qw422016.N().S(`
    </div>
    <form name='fFilter' id='fFilter' class='form-simple tr thead' role='form' action='`)
	//line views/templates/tables/table_view.qtpl:68
	qw422016.E().S(query.Href)
	//line views/templates/tables/table_view.qtpl:68
	qw422016.N().S(`/?filter'  method='post'  target='content' style='width:auto;$style' onsubmit='return SaveObject( this );' enctype='multipart/form-data' title='$title' oninput='return FormIsModified( event, this);'
    onabort="alert('fFilter');">
        <input name="`)
	//line views/templates/tables/table_view.qtpl:70
	qw422016.E().S(query.Tables[0].Name)
	//line views/templates/tables/table_view.qtpl:70
	qw422016.N().S(`" hidden />
        `)
	//line views/templates/tables/table_view.qtpl:71
	qw422016.N().S(filterFields)
	//line views/templates/tables/table_view.qtpl:71
	qw422016.N().S(`
    </form>
</div>
`)
//line views/templates/tables/table_view.qtpl:74
}

//line views/templates/tables/table_view.qtpl:74
func (query *QueryStruct) writerenderHeadTables(qq422016 qtio422016.Writer) {
	//line views/templates/tables/table_view.qtpl:74
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:74
	query.streamrenderHeadTables(qw422016)
	//line views/templates/tables/table_view.qtpl:74
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:74
}

//line views/templates/tables/table_view.qtpl:74
func (query *QueryStruct) renderHeadTables() string {
	//line views/templates/tables/table_view.qtpl:74
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:74
	query.writerenderHeadTables(qb422016)
	//line views/templates/tables/table_view.qtpl:74
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:74
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:74
	return qs422016
//line views/templates/tables/table_view.qtpl:74
}

//line views/templates/tables/table_view.qtpl:75
func (query *QueryStruct) StreamRenderTable(qw422016 *qt422016.Writer) {
	//line views/templates/tables/table_view.qtpl:75
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:77
	if err := query.beforeRender(); err != nil {
		return
	}

	//line views/templates/tables/table_view.qtpl:80
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:81
	query.streamrenderHeadTables(qw422016)
	//line views/templates/tables/table_view.qtpl:81
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:82
	query.streamrenderBodyTables(qw422016)
	//line views/templates/tables/table_view.qtpl:82
	qw422016.N().S(`
`)
//line views/templates/tables/table_view.qtpl:83
}

//line views/templates/tables/table_view.qtpl:83
func (query *QueryStruct) WriteRenderTable(qq422016 qtio422016.Writer) {
	//line views/templates/tables/table_view.qtpl:83
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:83
	query.StreamRenderTable(qw422016)
	//line views/templates/tables/table_view.qtpl:83
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:83
}

//line views/templates/tables/table_view.qtpl:83
func (query *QueryStruct) RenderTable() string {
	//line views/templates/tables/table_view.qtpl:83
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:83
	query.WriteRenderTable(qb422016)
	//line views/templates/tables/table_view.qtpl:83
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:83
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:83
	return qs422016
//line views/templates/tables/table_view.qtpl:83
}

//line views/templates/tables/table_view.qtpl:84
func StreamShowTable(qw422016 *qt422016.Writer, tableName string, fields forms.FieldsTable, rows *sql.Rows) {
	//line views/templates/tables/table_view.qtpl:84
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:86
	var query QueryStruct
	query.Rows = rows
	query.Href = "/admin/table/" + tableName
	query.HrefEdit = "/admin/row/edit/?table=" + tableName + "&id="
	query.Tables = append(query.Tables, &fields)

	if err := query.beforeRender(); err != nil {
		return
	}

	//line views/templates/tables/table_view.qtpl:95
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:96
	query.streamrenderHeadTables(qw422016)
	//line views/templates/tables/table_view.qtpl:96
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:97
	query.streamrenderBodyTables(qw422016)
	//line views/templates/tables/table_view.qtpl:97
	qw422016.N().S(`
`)
//line views/templates/tables/table_view.qtpl:98
}

//line views/templates/tables/table_view.qtpl:98
func WriteShowTable(qq422016 qtio422016.Writer, tableName string, fields forms.FieldsTable, rows *sql.Rows) {
	//line views/templates/tables/table_view.qtpl:98
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:98
	StreamShowTable(qw422016, tableName, fields, rows)
	//line views/templates/tables/table_view.qtpl:98
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:98
}

//line views/templates/tables/table_view.qtpl:98
func ShowTable(tableName string, fields forms.FieldsTable, rows *sql.Rows) string {
	//line views/templates/tables/table_view.qtpl:98
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:98
	WriteShowTable(qb422016, tableName, fields, rows)
	//line views/templates/tables/table_view.qtpl:98
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:98
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:98
	return qs422016
//line views/templates/tables/table_view.qtpl:98
}

//line views/templates/tables/table_view.qtpl:99
func (query *QueryStruct) streamrenderBodyTables(qw422016 *qt422016.Writer) {
	//line views/templates/tables/table_view.qtpl:99
	qw422016.N().S(`
<div id="table_body" >
        `)
	//line views/templates/tables/table_view.qtpl:101
	for query.Rows.Next() {
		//line views/templates/tables/table_view.qtpl:101
		qw422016.N().S(`
            `)
		//line views/templates/tables/table_view.qtpl:103
		if err := query.Rows.Scan(query.row...); err != nil {
			log.Println(err)
			continue
		}

		required, tablePrefix, titleLabel := "", "", ""

		//line views/templates/tables/table_view.qtpl:109
		qw422016.N().S(`
        <div class='tr' style='$style_div'>
            `)
		//line views/templates/tables/table_view.qtpl:111
		for idx, rawField := range query.row {
			//line views/templates/tables/table_view.qtpl:111
			qw422016.N().S(`
                `)
			//line views/templates/tables/table_view.qtpl:113
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

			//line views/templates/tables/table_view.qtpl:133
			qw422016.N().S(`
                `)
			//line views/templates/tables/table_view.qtpl:134
			if idx == 0 {
				//line views/templates/tables/table_view.qtpl:134
				qw422016.N().S(`
                    <div name="`)
				//line views/templates/tables/table_view.qtpl:135
				qw422016.E().S(key)
				//line views/templates/tables/table_view.qtpl:135
				qw422016.N().S(`" class="td" style="`)
				//line views/templates/tables/table_view.qtpl:135
				qw422016.E().S(fieldStruct.CSSStyle)
				//line views/templates/tables/table_view.qtpl:135
				qw422016.N().S(`">
                        <a href="`)
				//line views/templates/tables/table_view.qtpl:136
				qw422016.N().S(query.HrefEdit)
				//line views/templates/tables/table_view.qtpl:136
				qw422016.E().S(val)
				//line views/templates/tables/table_view.qtpl:136
				qw422016.N().S(`" target="content">`)
				//line views/templates/tables/table_view.qtpl:136
				qw422016.E().S(val)
				//line views/templates/tables/table_view.qtpl:136
				qw422016.N().S(`</a>
                    </div>
                    `)
				//line views/templates/tables/table_view.qtpl:139
				tablePrefix = fields.Name
				fields.ID, _ = strconv.Atoi(val)

				//line views/templates/tables/table_view.qtpl:141
				qw422016.N().S(`
                `)
				//line views/templates/tables/table_view.qtpl:142
			} else {
				//line views/templates/tables/table_view.qtpl:142
				qw422016.N().S(`
                `)
				//line views/templates/tables/table_view.qtpl:145
				qw422016.N().S(`
                    <div name="`)
				//line views/templates/tables/table_view.qtpl:146
				qw422016.E().S(key)
				//line views/templates/tables/table_view.qtpl:146
				qw422016.N().S(`" class="td `)
				//line views/templates/tables/table_view.qtpl:146
				qw422016.N().S(required)
				//line views/templates/tables/table_view.qtpl:146
				qw422016.N().S(` field-`)
				//line views/templates/tables/table_view.qtpl:146
				qw422016.E().S(nameInput)
				//line views/templates/tables/table_view.qtpl:146
				qw422016.N().S(` `)
				//line views/templates/tables/table_view.qtpl:146
				qw422016.E().S(fieldStruct.CSSClass)
				//line views/templates/tables/table_view.qtpl:146
				qw422016.N().S(`"
                         style="`)
				//line views/templates/tables/table_view.qtpl:147
				qw422016.E().S(fieldStruct.CSSStyle)
				//line views/templates/tables/table_view.qtpl:147
				qw422016.N().S(`;margin:auto; height:auto;">
                        `)
				//line views/templates/tables/table_view.qtpl:148
				if key == "parent_id" {
					//line views/templates/tables/table_view.qtpl:148
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:149
					fieldStruct.StreamRenderParentSelect(qw422016, fields.Name, key, val, titleLabel, required, events, dataJson)
					//line views/templates/tables/table_view.qtpl:149
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:150
				} else if strings.HasPrefix(key, "id_") {
					//line views/templates/tables/table_view.qtpl:150
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:151
					fieldStruct.StreamRenderForeignSelect(qw422016, tablePrefix, key, val, titleLabel, required, events, dataJson)
					//line views/templates/tables/table_view.qtpl:151
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:152
				} else if strings.HasPrefix(key, "setid_") || strings.HasPrefix(key, "nodeid_") {
					//line views/templates/tables/table_view.qtpl:152
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:153
					fieldStruct.StreamRenderMultiSelect(qw422016, fields, tablePrefix, key, val, "См. ", required)
					//line views/templates/tables/table_view.qtpl:153
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:154
				} else if strings.HasPrefix(key, "tableid_") {
					//line views/templates/tables/table_view.qtpl:154
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:155
					fieldStruct.StreamRenderTable(qw422016, fields, tablePrefix, key, val, "Табл", required)
					//line views/templates/tables/table_view.qtpl:155
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:156
				} else {
					//line views/templates/tables/table_view.qtpl:156
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:157
					switch fieldStruct.DATA_TYPE {
					//line views/templates/tables/table_view.qtpl:158
					case "tinyint":
						//line views/templates/tables/table_view.qtpl:158
						qw422016.N().S(`
                                    `)
						//line views/templates/tables/table_view.qtpl:160
						checked := ""
						if val == "1" {
							checked = "checked"
						}

						//line views/templates/tables/table_view.qtpl:164
						qw422016.N().S(`
                                    `)
						//line views/templates/tables/table_view.qtpl:165
						forms.StreamRenderCheckBox(qw422016, nameInput, "1", titleLabel, 1, checked, required, events, dataJson)
						//line views/templates/tables/table_view.qtpl:165
						qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:166
					case "enum":
						//line views/templates/tables/table_view.qtpl:166
						qw422016.N().S(`
                                `)
						//line views/templates/tables/table_view.qtpl:167
						t := fieldStruct.RenderEnum(nameInput, val, required, events, dataJson)

						//line views/templates/tables/table_view.qtpl:167
						qw422016.N().S(`
                                `)
						//line views/templates/tables/table_view.qtpl:168
						qw422016.N().S(t)
						//line views/templates/tables/table_view.qtpl:168
						qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:169
					case "set":
						//line views/templates/tables/table_view.qtpl:169
						qw422016.N().S(`
                                `)
						//line views/templates/tables/table_view.qtpl:170
						t := fieldStruct.RenderSet(nameInput, val, required, events, dataJson)

						//line views/templates/tables/table_view.qtpl:170
						qw422016.N().S(`
                                `)
						//line views/templates/tables/table_view.qtpl:171
						qw422016.N().S(t)
						//line views/templates/tables/table_view.qtpl:171
						qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:172
					case "text":
						//line views/templates/tables/table_view.qtpl:172
						qw422016.N().S(`
                                <textarea id="`)
						//line views/templates/tables/table_view.qtpl:173
						qw422016.E().S(key)
						//line views/templates/tables/table_view.qtpl:173
						qw422016.N().S(`" name="`)
						//line views/templates/tables/table_view.qtpl:173
						qw422016.E().S(nameInput)
						//line views/templates/tables/table_view.qtpl:173
						qw422016.N().S(`" class="controls"  `)
						//line views/templates/tables/table_view.qtpl:173
						qw422016.N().S(events)
						//line views/templates/tables/table_view.qtpl:173
						qw422016.N().S(` `)
						//line views/templates/tables/table_view.qtpl:173
						qw422016.N().S(dataJson)
						//line views/templates/tables/table_view.qtpl:173
						qw422016.N().S(`>
                                    `)
						//line views/templates/tables/table_view.qtpl:174
						qw422016.N().S(val)
						//line views/templates/tables/table_view.qtpl:174
						qw422016.N().S(`
                                </textarea>
                            `)
					//line views/templates/tables/table_view.qtpl:176
					case "blob":
						//line views/templates/tables/table_view.qtpl:176
						qw422016.N().S(`
                                <image src="/images/expand_hover.png" />
                            `)
					//line views/templates/tables/table_view.qtpl:178
					default:
						//line views/templates/tables/table_view.qtpl:178
						qw422016.N().S(`
                                `)
						//line views/templates/tables/table_view.qtpl:179
						if fieldStruct.LinkTD > "" {
							//line views/templates/tables/table_view.qtpl:179
							qw422016.N().S(`
                                 <a href="`)
							//line views/templates/tables/table_view.qtpl:180
							qw422016.N().S(fieldStruct.LinkTD)
							//line views/templates/tables/table_view.qtpl:180
							qw422016.N().S(`?id=`)
							//line views/templates/tables/table_view.qtpl:180
							qw422016.N().D(fields.ID)
							//line views/templates/tables/table_view.qtpl:180
							qw422016.N().S(`" target="_blank"> !! </a>
                                 `)
							//line views/templates/tables/table_view.qtpl:181
						}
						//line views/templates/tables/table_view.qtpl:181
						qw422016.N().S(`
                                `)
						//line views/templates/tables/table_view.qtpl:182
						qw422016.E().S(val)
						//line views/templates/tables/table_view.qtpl:182
						qw422016.N().S(`
                            `)
						//line views/templates/tables/table_view.qtpl:183
					}
					//line views/templates/tables/table_view.qtpl:183
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:184
				}
				//line views/templates/tables/table_view.qtpl:184
				qw422016.N().S(`
                    </div>
                `)
				//line views/templates/tables/table_view.qtpl:186
			}
			//line views/templates/tables/table_view.qtpl:186
			qw422016.N().S(`
            `)
			//line views/templates/tables/table_view.qtpl:187
		}
		//line views/templates/tables/table_view.qtpl:187
		qw422016.N().S(`
        </div>
        `)
		//line views/templates/tables/table_view.qtpl:189
	}
	//line views/templates/tables/table_view.qtpl:189
	qw422016.N().S(`
    </div>
</div>
<!--/table-->

`)
//line views/templates/tables/table_view.qtpl:194
}

//line views/templates/tables/table_view.qtpl:194
func (query *QueryStruct) writerenderBodyTables(qq422016 qtio422016.Writer) {
	//line views/templates/tables/table_view.qtpl:194
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:194
	query.streamrenderBodyTables(qw422016)
	//line views/templates/tables/table_view.qtpl:194
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:194
}

//line views/templates/tables/table_view.qtpl:194
func (query *QueryStruct) renderBodyTables() string {
	//line views/templates/tables/table_view.qtpl:194
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:194
	query.writerenderBodyTables(qb422016)
	//line views/templates/tables/table_view.qtpl:194
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:194
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:194
	return qs422016
//line views/templates/tables/table_view.qtpl:194
}
