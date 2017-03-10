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
<script src="/tables.js"></script>
<div class="table">
<div class='thead' >
    <div class='tr' >
        `)
	//line views/templates/tables/table_view.qtpl:21
	var figure, filterFields string
	if query.Order == "" {
		query.Order = "id"
	}

	//line views/templates/tables/table_view.qtpl:25
	qw422016.N().S(`
            `)
	//line views/templates/tables/table_view.qtpl:26
	for idx, fieldName := range query.columns {
		//line views/templates/tables/table_view.qtpl:26
		qw422016.N().S(`
                `)
		//line views/templates/tables/table_view.qtpl:28
		fieldStruct := query.fields[idx]
		key, titleFull, titleLabel := fieldStruct.COLUMN_NAME, fieldStruct.COLUMN_COMMENT, fieldStruct.COLUMN_COMMENT
		if fieldStruct.InputType == "" {
			fieldStruct.InputType = forms.StyleInput(fieldStruct.DATA_TYPE)
		}
		width, maxLen := forms.GetLengthFromType(fieldStruct.InputType)

		if (fieldStruct.DATA_TYPE == "varchar") && (fieldStruct.CHARACTER_MAXIMUM_LENGTH == 255) {
			width, maxLen = 200, 100
		}
		fieldStruct.CSSStyle = fmt.Sprintf("width:%dpx;", width)
		query.widthTable += width
		filterFields += inputFilterField(key, query.Tables[0].Name, fieldStruct)
		if len(titleFull) > maxLen {

			if pos := strings.LastIndex(titleFull[:maxLen+1], " "); pos > 0 {
				titleLabel = titleFull[:pos] + "..."
			} else {
				titleLabel = titleFull[:maxLen] + "..."
			}
		}

		if titleLabel == "" {
			titleLabel = fieldName
		}

		//line views/templates/tables/table_view.qtpl:53
		qw422016.N().S(`
                `)
		//line views/templates/tables/table_view.qtpl:54
		qw422016.N().S(`
                `)
		//line views/templates/tables/table_view.qtpl:55
		if figure != fieldStruct.Figure {
			//line views/templates/tables/table_view.qtpl:55
			qw422016.N().S(`
                    `)
			//line views/templates/tables/table_view.qtpl:56
			if figure > "" {
				//line views/templates/tables/table_view.qtpl:56
				qw422016.N().S(`
                        </div>
                    `)
				//line views/templates/tables/table_view.qtpl:58
			}
			//line views/templates/tables/table_view.qtpl:58
			qw422016.N().S(`
                    `)
			//line views/templates/tables/table_view.qtpl:60
			figure = fieldStruct.Figure

			//line views/templates/tables/table_view.qtpl:61
			qw422016.N().S(`
                    `)
			//line views/templates/tables/table_view.qtpl:62
			if figure > "" {
				//line views/templates/tables/table_view.qtpl:62
				qw422016.N().S(`
                            <div class='td-outer' style="outline: 1px inset gray;">
                                <div class='th' style="outline: 1px outset gray; padding-top: 1px;"> `)
				//line views/templates/tables/table_view.qtpl:64
				qw422016.E().S(figure)
				//line views/templates/tables/table_view.qtpl:64
				qw422016.N().S(`</div>
                    `)
				//line views/templates/tables/table_view.qtpl:65
			}
			//line views/templates/tables/table_view.qtpl:65
			qw422016.N().S(`
                `)
			//line views/templates/tables/table_view.qtpl:66
		}
		//line views/templates/tables/table_view.qtpl:66
		qw422016.N().S(`
                 <div title="`)
		//line views/templates/tables/table_view.qtpl:67
		qw422016.E().S(fieldName)
		//line views/templates/tables/table_view.qtpl:67
		qw422016.N().S(`" class="td`)
		//line views/templates/tables/table_view.qtpl:67
		if fieldName == query.Order {
			//line views/templates/tables/table_view.qtpl:67
			qw422016.N().S(` td-order`)
			//line views/templates/tables/table_view.qtpl:67
		}
		//line views/templates/tables/table_view.qtpl:67
		qw422016.N().S(`"
                 style="`)
		//line views/templates/tables/table_view.qtpl:68
		qw422016.E().S(fieldStruct.CSSStyle)
		//line views/templates/tables/table_view.qtpl:68
		qw422016.N().S(`">
                   <a href="`)
		//line views/templates/tables/table_view.qtpl:69
		qw422016.E().S(query.Href)
		//line views/templates/tables/table_view.qtpl:69
		qw422016.N().S(`/?order=`)
		//line views/templates/tables/table_view.qtpl:69
		qw422016.E().S(fieldName)
		//line views/templates/tables/table_view.qtpl:69
		qw422016.N().S(`" title="`)
		//line views/templates/tables/table_view.qtpl:69
		qw422016.E().S(titleFull)
		//line views/templates/tables/table_view.qtpl:69
		qw422016.N().S(`">`)
		//line views/templates/tables/table_view.qtpl:69
		qw422016.E().S(titleLabel)
		//line views/templates/tables/table_view.qtpl:69
		qw422016.N().S(`</a>
                </div>
            `)
		//line views/templates/tables/table_view.qtpl:71
	}
	//line views/templates/tables/table_view.qtpl:71
	qw422016.N().S(`
            `)
	//line views/templates/tables/table_view.qtpl:72
	if figure > "" {
		//line views/templates/tables/table_view.qtpl:72
		qw422016.N().S(`
                </div>
            `)
		//line views/templates/tables/table_view.qtpl:74
	}
	//line views/templates/tables/table_view.qtpl:74
	qw422016.N().S(`
    </div>
    <form name='fFilter' id='fFilter' class='form-simple tr thead' role='form' action='`)
	//line views/templates/tables/table_view.qtpl:76
	qw422016.E().S(query.Href)
	//line views/templates/tables/table_view.qtpl:76
	qw422016.N().S(`/?filter' method='post'
    target='content' style='width:auto;$style' onsubmit='return SaveObject( this );' enctype='multipart/form-data'
    oninput='return FormIsModified(event, this);' onabort="alert('fFilter');">
        <output name="State" ></output>
        <input name="`)
	//line views/templates/tables/table_view.qtpl:80
	qw422016.E().S(query.Tables[0].Name)
	//line views/templates/tables/table_view.qtpl:80
	qw422016.N().S(`" hidden />
        `)
	//line views/templates/tables/table_view.qtpl:81
	qw422016.N().S(filterFields)
	//line views/templates/tables/table_view.qtpl:81
	qw422016.N().S(`
    </form>
</div>
`)
//line views/templates/tables/table_view.qtpl:84
}

//line views/templates/tables/table_view.qtpl:84
func (query *QueryStruct) writerenderHeadTables(qq422016 qtio422016.Writer) {
	//line views/templates/tables/table_view.qtpl:84
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:84
	query.streamrenderHeadTables(qw422016)
	//line views/templates/tables/table_view.qtpl:84
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:84
}

//line views/templates/tables/table_view.qtpl:84
func (query *QueryStruct) renderHeadTables() string {
	//line views/templates/tables/table_view.qtpl:84
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:84
	query.writerenderHeadTables(qb422016)
	//line views/templates/tables/table_view.qtpl:84
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:84
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:84
	return qs422016
//line views/templates/tables/table_view.qtpl:84
}

//line views/templates/tables/table_view.qtpl:85
func (query *QueryStruct) StreamRenderTable(qw422016 *qt422016.Writer) {
	//line views/templates/tables/table_view.qtpl:85
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:87
	if err := query.beforeRender(); err != nil {
		return
	}

	//line views/templates/tables/table_view.qtpl:90
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:91
	query.streamrenderHeadTables(qw422016)
	//line views/templates/tables/table_view.qtpl:91
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:92
	query.streamrenderBodyTables(qw422016)
	//line views/templates/tables/table_view.qtpl:92
	qw422016.N().S(`
`)
//line views/templates/tables/table_view.qtpl:93
}

//line views/templates/tables/table_view.qtpl:93
func (query *QueryStruct) WriteRenderTable(qq422016 qtio422016.Writer) {
	//line views/templates/tables/table_view.qtpl:93
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:93
	query.StreamRenderTable(qw422016)
	//line views/templates/tables/table_view.qtpl:93
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:93
}

//line views/templates/tables/table_view.qtpl:93
func (query *QueryStruct) RenderTable() string {
	//line views/templates/tables/table_view.qtpl:93
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:93
	query.WriteRenderTable(qb422016)
	//line views/templates/tables/table_view.qtpl:93
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:93
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:93
	return qs422016
//line views/templates/tables/table_view.qtpl:93
}

//line views/templates/tables/table_view.qtpl:94
func StreamShowTable(qw422016 *qt422016.Writer, tableName string, fields forms.FieldsTable, rows *sql.Rows) {
	//line views/templates/tables/table_view.qtpl:94
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:96
	var query QueryStruct
	query.Rows = rows
	query.Href = "/admin/table/" + tableName
	query.HrefEdit = "/admin/row/edit/?table=" + tableName + "&id="
	query.Tables = append(query.Tables, &fields)

	if err := query.beforeRender(); err != nil {
		return
	}

	//line views/templates/tables/table_view.qtpl:105
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:106
	query.streamrenderHeadTables(qw422016)
	//line views/templates/tables/table_view.qtpl:106
	qw422016.N().S(`
        `)
	//line views/templates/tables/table_view.qtpl:107
	query.streamrenderBodyTables(qw422016)
	//line views/templates/tables/table_view.qtpl:107
	qw422016.N().S(`
`)
//line views/templates/tables/table_view.qtpl:108
}

//line views/templates/tables/table_view.qtpl:108
func WriteShowTable(qq422016 qtio422016.Writer, tableName string, fields forms.FieldsTable, rows *sql.Rows) {
	//line views/templates/tables/table_view.qtpl:108
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:108
	StreamShowTable(qw422016, tableName, fields, rows)
	//line views/templates/tables/table_view.qtpl:108
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:108
}

//line views/templates/tables/table_view.qtpl:108
func ShowTable(tableName string, fields forms.FieldsTable, rows *sql.Rows) string {
	//line views/templates/tables/table_view.qtpl:108
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:108
	WriteShowTable(qb422016, tableName, fields, rows)
	//line views/templates/tables/table_view.qtpl:108
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:108
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:108
	return qs422016
//line views/templates/tables/table_view.qtpl:108
}

//line views/templates/tables/table_view.qtpl:109
func (query *QueryStruct) streamrenderBodyTables(qw422016 *qt422016.Writer) {
	//line views/templates/tables/table_view.qtpl:109
	qw422016.N().S(`
<div id="table_body" >
        `)
	//line views/templates/tables/table_view.qtpl:111
	for query.Rows.Next() {
		//line views/templates/tables/table_view.qtpl:111
		qw422016.N().S(`
            `)
		//line views/templates/tables/table_view.qtpl:113
		if err := query.Rows.Scan(query.row...); err != nil {
			log.Println(err)
			continue
		}

		required, tablePrefix, titleLabel := "", "", ""

		//line views/templates/tables/table_view.qtpl:119
		qw422016.N().S(`
        <div class='tr' style='$style_div'>
            `)
		//line views/templates/tables/table_view.qtpl:121
		for idx, rawField := range query.row {
			//line views/templates/tables/table_view.qtpl:121
			qw422016.N().S(`
                `)
			//line views/templates/tables/table_view.qtpl:123
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

			//line views/templates/tables/table_view.qtpl:143
			qw422016.N().S(`
                <div name="`)
			//line views/templates/tables/table_view.qtpl:144
			qw422016.E().S(key)
			//line views/templates/tables/table_view.qtpl:144
			qw422016.N().S(`" class="td `)
			//line views/templates/tables/table_view.qtpl:144
			qw422016.N().S(required)
			//line views/templates/tables/table_view.qtpl:144
			qw422016.N().S(` field-`)
			//line views/templates/tables/table_view.qtpl:144
			qw422016.E().S(nameInput)
			//line views/templates/tables/table_view.qtpl:144
			qw422016.N().S(` `)
			//line views/templates/tables/table_view.qtpl:144
			qw422016.E().S(fieldStruct.CSSClass)
			//line views/templates/tables/table_view.qtpl:144
			qw422016.N().S(`"
                         style="`)
			//line views/templates/tables/table_view.qtpl:145
			qw422016.E().S(fieldStruct.CSSStyle)
			//line views/templates/tables/table_view.qtpl:145
			qw422016.N().S(`;margin:auto; height:auto;">
                `)
			//line views/templates/tables/table_view.qtpl:146
			if idx == 0 {
				//line views/templates/tables/table_view.qtpl:146
				qw422016.N().S(`
                    <a href="`)
				//line views/templates/tables/table_view.qtpl:147
				qw422016.N().S(query.HrefEdit)
				//line views/templates/tables/table_view.qtpl:147
				qw422016.E().S(val)
				//line views/templates/tables/table_view.qtpl:147
				qw422016.N().S(`" target="content">`)
				//line views/templates/tables/table_view.qtpl:147
				qw422016.E().S(val)
				//line views/templates/tables/table_view.qtpl:147
				qw422016.N().S(`</a>
                    `)
				//line views/templates/tables/table_view.qtpl:149
				tablePrefix = fields.Name
				fields.ID, _ = strconv.Atoi(val)

				//line views/templates/tables/table_view.qtpl:151
				qw422016.N().S(`
                `)
				//line views/templates/tables/table_view.qtpl:152
			} else if fieldStruct.Html > "" {
				//line views/templates/tables/table_view.qtpl:152
				qw422016.N().S(`
                    `)
				//line views/templates/tables/table_view.qtpl:153
				qw422016.N().S(fieldStruct.Html)
				//line views/templates/tables/table_view.qtpl:153
				qw422016.N().S(`
                `)
				//line views/templates/tables/table_view.qtpl:154
			} else {
				//line views/templates/tables/table_view.qtpl:154
				qw422016.N().S(`
                `)
				//line views/templates/tables/table_view.qtpl:157
				qw422016.N().S(`
                        `)
				//line views/templates/tables/table_view.qtpl:158
				if key == "parent_id" {
					//line views/templates/tables/table_view.qtpl:158
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:159
					fieldStruct.StreamRenderParentSelect(qw422016, fields.Name, key, val, titleLabel, required, events, dataJson)
					//line views/templates/tables/table_view.qtpl:159
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:160
				} else if strings.HasPrefix(key, "id_") {
					//line views/templates/tables/table_view.qtpl:160
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:161
					fieldStruct.StreamRenderForeignSelect(qw422016, tablePrefix, key, val, titleLabel, required, events, dataJson)
					//line views/templates/tables/table_view.qtpl:161
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:162
				} else if strings.HasPrefix(key, "setid_") || strings.HasPrefix(key, "nodeid_") {
					//line views/templates/tables/table_view.qtpl:162
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:163
					fieldStruct.StreamRenderMultiSelect(qw422016, fields, tablePrefix, key, val, "См. ", required)
					//line views/templates/tables/table_view.qtpl:163
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:164
				} else if strings.HasPrefix(key, "tableid_") {
					//line views/templates/tables/table_view.qtpl:164
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:165
					fieldStruct.StreamRenderTable(qw422016, fields, tablePrefix, key, val, "Табл", required)
					//line views/templates/tables/table_view.qtpl:165
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:166
				} else {
					//line views/templates/tables/table_view.qtpl:166
					qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:167
					switch fieldStruct.DATA_TYPE {
					//line views/templates/tables/table_view.qtpl:168
					case "tinyint":
						//line views/templates/tables/table_view.qtpl:168
						qw422016.N().S(`
                                    `)
						//line views/templates/tables/table_view.qtpl:170
						checked := ""
						if val == "1" {
							checked = "checked"
						}

						//line views/templates/tables/table_view.qtpl:174
						qw422016.N().S(`
                                    `)
						//line views/templates/tables/table_view.qtpl:175
						forms.StreamRenderCheckBox(qw422016, nameInput, "1", titleLabel, 1, checked, required, events, dataJson)
						//line views/templates/tables/table_view.qtpl:175
						qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:176
					case "enum":
						//line views/templates/tables/table_view.qtpl:176
						qw422016.N().S(`
                                `)
						//line views/templates/tables/table_view.qtpl:177
						t := fieldStruct.RenderEnum(nameInput, val, required, events, dataJson)

						//line views/templates/tables/table_view.qtpl:177
						qw422016.N().S(`
                                `)
						//line views/templates/tables/table_view.qtpl:178
						qw422016.N().S(t)
						//line views/templates/tables/table_view.qtpl:178
						qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:179
					case "set":
						//line views/templates/tables/table_view.qtpl:179
						qw422016.N().S(`
                                `)
						//line views/templates/tables/table_view.qtpl:180
						t := fieldStruct.RenderSet(nameInput, val, required, events, dataJson)

						//line views/templates/tables/table_view.qtpl:180
						qw422016.N().S(`
                                `)
						//line views/templates/tables/table_view.qtpl:181
						qw422016.N().S(t)
						//line views/templates/tables/table_view.qtpl:181
						qw422016.N().S(`
                            `)
					//line views/templates/tables/table_view.qtpl:182
					case "text":
						//line views/templates/tables/table_view.qtpl:182
						qw422016.N().S(`
                                <p id="`)
						//line views/templates/tables/table_view.qtpl:183
						qw422016.E().S(key)
						//line views/templates/tables/table_view.qtpl:183
						qw422016.N().S(`" name="`)
						//line views/templates/tables/table_view.qtpl:183
						qw422016.E().S(nameInput)
						//line views/templates/tables/table_view.qtpl:183
						qw422016.N().S(`" class="controls"  `)
						//line views/templates/tables/table_view.qtpl:183
						qw422016.N().S(events)
						//line views/templates/tables/table_view.qtpl:183
						qw422016.N().S(` `)
						//line views/templates/tables/table_view.qtpl:183
						qw422016.N().S(dataJson)
						//line views/templates/tables/table_view.qtpl:183
						qw422016.N().S(`>
                                    `)
						//line views/templates/tables/table_view.qtpl:184
						qw422016.N().S(val)
						//line views/templates/tables/table_view.qtpl:184
						qw422016.N().S(`
                                </p>
                            `)
					//line views/templates/tables/table_view.qtpl:186
					case "blob":
						//line views/templates/tables/table_view.qtpl:186
						qw422016.N().S(`
                                <image src="/images/expand_hover.png" />
                            `)
					//line views/templates/tables/table_view.qtpl:188
					default:
						//line views/templates/tables/table_view.qtpl:188
						qw422016.N().S(`
                                `)
						//line views/templates/tables/table_view.qtpl:189
						if fieldStruct.LinkTD > "" {
							//line views/templates/tables/table_view.qtpl:189
							qw422016.N().S(`
                                 <a href="`)
							//line views/templates/tables/table_view.qtpl:190
							qw422016.N().S(fieldStruct.LinkTD)
							//line views/templates/tables/table_view.qtpl:190
							qw422016.N().S(`?id=`)
							//line views/templates/tables/table_view.qtpl:190
							qw422016.N().D(fields.ID)
							//line views/templates/tables/table_view.qtpl:190
							qw422016.N().S(`" target="_blank">`)
							//line views/templates/tables/table_view.qtpl:190
							qw422016.E().S(val)
							//line views/templates/tables/table_view.qtpl:190
							qw422016.N().S(`</a>
                                 `)
							//line views/templates/tables/table_view.qtpl:191
						} else {
							//line views/templates/tables/table_view.qtpl:191
							qw422016.N().S(`
                                    `)
							//line views/templates/tables/table_view.qtpl:192
							qw422016.E().S(val)
							//line views/templates/tables/table_view.qtpl:192
							qw422016.N().S(`
                                 `)
							//line views/templates/tables/table_view.qtpl:193
						}
						//line views/templates/tables/table_view.qtpl:193
						qw422016.N().S(`
                            `)
						//line views/templates/tables/table_view.qtpl:194
					}
					//line views/templates/tables/table_view.qtpl:194
					qw422016.N().S(`
                        `)
					//line views/templates/tables/table_view.qtpl:195
				}
				//line views/templates/tables/table_view.qtpl:195
				qw422016.N().S(`
                `)
				//line views/templates/tables/table_view.qtpl:196
			}
			//line views/templates/tables/table_view.qtpl:196
			qw422016.N().S(`
                </div> <!-- field -->
            `)
			//line views/templates/tables/table_view.qtpl:198
		}
		//line views/templates/tables/table_view.qtpl:198
		qw422016.N().S(`
            `)
		//line views/templates/tables/table_view.qtpl:199
		for _, fieldStruct := range query.PostFields {
			//line views/templates/tables/table_view.qtpl:199
			qw422016.N().S(`
                <div name="`)
			//line views/templates/tables/table_view.qtpl:200
			qw422016.E().S(fieldStruct.COLUMN_NAME)
			//line views/templates/tables/table_view.qtpl:200
			qw422016.N().S(`" class="td `)
			//line views/templates/tables/table_view.qtpl:200
			qw422016.E().S(fieldStruct.CSSClass)
			//line views/templates/tables/table_view.qtpl:200
			qw422016.N().S(`"
                         style="`)
			//line views/templates/tables/table_view.qtpl:201
			qw422016.E().S(fieldStruct.CSSStyle)
			//line views/templates/tables/table_view.qtpl:201
			qw422016.N().S(`;margin:auto; height:auto;" >
                    `)
			//line views/templates/tables/table_view.qtpl:202
			qw422016.N().S(fieldStruct.Html)
			//line views/templates/tables/table_view.qtpl:202
			qw422016.N().S(`
                </div> <!-- field -->
            `)
			//line views/templates/tables/table_view.qtpl:204
		}
		//line views/templates/tables/table_view.qtpl:204
		qw422016.N().S(`
        </div> <!-- tr -->
        `)
		//line views/templates/tables/table_view.qtpl:206
	}
	//line views/templates/tables/table_view.qtpl:206
	qw422016.N().S(`
    </div>
</div>
`)
//line views/templates/tables/table_view.qtpl:209
}

//line views/templates/tables/table_view.qtpl:209
func (query *QueryStruct) writerenderBodyTables(qq422016 qtio422016.Writer) {
	//line views/templates/tables/table_view.qtpl:209
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:209
	query.streamrenderBodyTables(qw422016)
	//line views/templates/tables/table_view.qtpl:209
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:209
}

//line views/templates/tables/table_view.qtpl:209
func (query *QueryStruct) renderBodyTables() string {
	//line views/templates/tables/table_view.qtpl:209
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:209
	query.writerenderBodyTables(qb422016)
	//line views/templates/tables/table_view.qtpl:209
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:209
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:209
	return qs422016
//line views/templates/tables/table_view.qtpl:209
}

//line views/templates/tables/table_view.qtpl:210
func streamrenderFilterSelect(qw422016 *qt422016.Writer, key string, fieldStruct *forms.FieldStructure) {
	//line views/templates/tables/table_view.qtpl:210
	qw422016.N().S(`
    <select id="`)
	//line views/templates/tables/table_view.qtpl:211
	qw422016.E().S(key)
	//line views/templates/tables/table_view.qtpl:211
	qw422016.N().S(`" name="`)
	//line views/templates/tables/table_view.qtpl:211
	qw422016.E().S(key)
	//line views/templates/tables/table_view.qtpl:211
	qw422016.N().S(`" onchange="return FilterIsModified( event, this);">
        <option selected>-</option>
        `)
	//line views/templates/tables/table_view.qtpl:213
	qw422016.N().S(fieldStruct.Html)
	//line views/templates/tables/table_view.qtpl:213
	qw422016.N().S(`
    </select>
`)
//line views/templates/tables/table_view.qtpl:215
}

//line views/templates/tables/table_view.qtpl:215
func writerenderFilterSelect(qq422016 qtio422016.Writer, key string, fieldStruct *forms.FieldStructure) {
	//line views/templates/tables/table_view.qtpl:215
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:215
	streamrenderFilterSelect(qw422016, key, fieldStruct)
	//line views/templates/tables/table_view.qtpl:215
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:215
}

//line views/templates/tables/table_view.qtpl:215
func renderFilterSelect(key string, fieldStruct *forms.FieldStructure) string {
	//line views/templates/tables/table_view.qtpl:215
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:215
	writerenderFilterSelect(qb422016, key, fieldStruct)
	//line views/templates/tables/table_view.qtpl:215
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:215
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:215
	return qs422016
//line views/templates/tables/table_view.qtpl:215
}

//line views/templates/tables/table_view.qtpl:216
func streaminputFilterField(qw422016 *qt422016.Writer, key, nameTable string, fieldStruct *forms.FieldStructure) {
	//line views/templates/tables/table_view.qtpl:216
	qw422016.N().S(`
 <div title="`)
	//line views/templates/tables/table_view.qtpl:217
	qw422016.E().S(key)
	//line views/templates/tables/table_view.qtpl:217
	qw422016.N().S(`" class="td" style="`)
	//line views/templates/tables/table_view.qtpl:217
	qw422016.N().S(fieldStruct.CSSStyle)
	//line views/templates/tables/table_view.qtpl:217
	qw422016.N().S(`">
                        `)
	//line views/templates/tables/table_view.qtpl:218
	if key == "parent_id" {
		//line views/templates/tables/table_view.qtpl:218
		qw422016.N().S(`
                                `)
		//line views/templates/tables/table_view.qtpl:219
		fieldStruct.GetOptions(nameTable, "")

		//line views/templates/tables/table_view.qtpl:219
		qw422016.N().S(`
                                `)
		//line views/templates/tables/table_view.qtpl:220
		streamrenderFilterSelect(qw422016, key, fieldStruct)
		//line views/templates/tables/table_view.qtpl:220
		qw422016.N().S(`
                        `)
		//line views/templates/tables/table_view.qtpl:221
	} else if strings.HasPrefix(key, "id_") {
		//line views/templates/tables/table_view.qtpl:221
		qw422016.N().S(`
                                `)
		//line views/templates/tables/table_view.qtpl:222
		fieldStruct.GetOptions(key[3:], "")

		//line views/templates/tables/table_view.qtpl:222
		qw422016.N().S(`
                                `)
		//line views/templates/tables/table_view.qtpl:223
		streamrenderFilterSelect(qw422016, key, fieldStruct)
		//line views/templates/tables/table_view.qtpl:223
		qw422016.N().S(`
                        `)
		//line views/templates/tables/table_view.qtpl:224
	} else if strings.HasPrefix(key, "setid_") || strings.HasPrefix(key, "nodeid_") {
		//line views/templates/tables/table_view.qtpl:224
		qw422016.N().S(`
                            `)
		//line views/templates/tables/table_view.qtpl:225
		fieldStruct.StreamRenderMultiSelect(qw422016, nil, nameTable, key, "", "См. ", "")
		//line views/templates/tables/table_view.qtpl:225
		qw422016.N().S(`
                        `)
		//line views/templates/tables/table_view.qtpl:226
	} else if strings.HasPrefix(key, "tableid_") {
		//line views/templates/tables/table_view.qtpl:226
		qw422016.N().S(`
                            `)
		//line views/templates/tables/table_view.qtpl:227
		fieldStruct.StreamRenderTable(qw422016, nil, nameTable, key, "", "Табл", "")
		//line views/templates/tables/table_view.qtpl:227
		qw422016.N().S(`
                        `)
		//line views/templates/tables/table_view.qtpl:228
	} else {
		//line views/templates/tables/table_view.qtpl:228
		qw422016.N().S(`
                            `)
		//line views/templates/tables/table_view.qtpl:229
		switch fieldStruct.DATA_TYPE {
		//line views/templates/tables/table_view.qtpl:230
		case "tinyint":
			//line views/templates/tables/table_view.qtpl:230
			qw422016.N().S(`
                                    `)
			//line views/templates/tables/table_view.qtpl:231
			forms.StreamRenderCheckBox(qw422016, key, "1", "", 1, "", "", "", "")
			//line views/templates/tables/table_view.qtpl:231
			qw422016.N().S(`
                            `)
		//line views/templates/tables/table_view.qtpl:232
		case "enum":
			//line views/templates/tables/table_view.qtpl:232
			qw422016.N().S(`
                                `)
			//line views/templates/tables/table_view.qtpl:233
			t := fieldStruct.RenderEnum(key, "", "", "", "")

			//line views/templates/tables/table_view.qtpl:233
			qw422016.N().S(`
                                `)
			//line views/templates/tables/table_view.qtpl:234
			qw422016.N().S(t)
			//line views/templates/tables/table_view.qtpl:234
			qw422016.N().S(`
                            `)
		//line views/templates/tables/table_view.qtpl:235
		case "set":
			//line views/templates/tables/table_view.qtpl:235
			qw422016.N().S(`
                                `)
			//line views/templates/tables/table_view.qtpl:236
			t := fieldStruct.RenderSet(key, "", "", "", "")

			//line views/templates/tables/table_view.qtpl:236
			qw422016.N().S(`
                                `)
			//line views/templates/tables/table_view.qtpl:237
			qw422016.N().S(t)
			//line views/templates/tables/table_view.qtpl:237
			qw422016.N().S(`
               `)
		//line views/templates/tables/table_view.qtpl:238
		default:
			//line views/templates/tables/table_view.qtpl:238
			qw422016.N().S(`
                        <input id="`)
			//line views/templates/tables/table_view.qtpl:239
			qw422016.E().S(key)
			//line views/templates/tables/table_view.qtpl:239
			qw422016.N().S(`" name="`)
			//line views/templates/tables/table_view.qtpl:239
			qw422016.E().S(key)
			//line views/templates/tables/table_view.qtpl:239
			qw422016.N().S(`" type=

                        `)
			//line views/templates/tables/table_view.qtpl:241
			if fieldStruct.InputType > "" {
				//line views/templates/tables/table_view.qtpl:241
				qw422016.N().S(`
                            "`)
				//line views/templates/tables/table_view.qtpl:242
				qw422016.E().S(fieldStruct.InputType)
				//line views/templates/tables/table_view.qtpl:242
				qw422016.N().S(`"
                        `)
				//line views/templates/tables/table_view.qtpl:243
			} else if fieldStruct.DATA_TYPE == "int" || fieldStruct.DATA_TYPE == "double" {
				//line views/templates/tables/table_view.qtpl:243
				qw422016.N().S(`
                            "number" `)
				//line views/templates/tables/table_view.qtpl:244
				if strings.Contains(fieldStruct.COLUMN_TYPE, "unsigned") {
					//line views/templates/tables/table_view.qtpl:244
					qw422016.N().S(`min="0"`)
					//line views/templates/tables/table_view.qtpl:244
				}
				//line views/templates/tables/table_view.qtpl:244
				qw422016.N().S(`
                        `)
				//line views/templates/tables/table_view.qtpl:245
			} else if fieldStruct.DATA_TYPE == "date" {
				//line views/templates/tables/table_view.qtpl:245
				qw422016.N().S(`
                            "date"  `)
				//line views/templates/tables/table_view.qtpl:246
				fieldStruct.StreamRenderDateAttributtes(qw422016)
				//line views/templates/tables/table_view.qtpl:246
				qw422016.N().S(`
                        `)
				//line views/templates/tables/table_view.qtpl:247
			} else if fieldStruct.DATA_TYPE == "datetime" {
				//line views/templates/tables/table_view.qtpl:247
				qw422016.N().S(`
                            "datetime" `)
				//line views/templates/tables/table_view.qtpl:248
				fieldStruct.StreamRenderDateAttributtes(qw422016)
				//line views/templates/tables/table_view.qtpl:248
				qw422016.N().S(`
                        `)
				//line views/templates/tables/table_view.qtpl:249
			} else if strings.Contains(key, "email") {
				//line views/templates/tables/table_view.qtpl:249
				qw422016.N().S(`
                            "email"
                        `)
				//line views/templates/tables/table_view.qtpl:251
			} else {
				//line views/templates/tables/table_view.qtpl:251
				qw422016.N().S(`
                            "text"
                               `)
				//line views/templates/tables/table_view.qtpl:253
				if fieldStruct.CHARACTER_MAXIMUM_LENGTH > 0 {
					//line views/templates/tables/table_view.qtpl:253
					qw422016.N().S(`
                                    maxlength="`)
					//line views/templates/tables/table_view.qtpl:254
					qw422016.N().D(fieldStruct.CHARACTER_MAXIMUM_LENGTH)
					//line views/templates/tables/table_view.qtpl:254
					qw422016.N().S(`"
                               `)
					//line views/templates/tables/table_view.qtpl:255
				}
				//line views/templates/tables/table_view.qtpl:255
				qw422016.N().S(`
                        `)
				//line views/templates/tables/table_view.qtpl:256
			}
			//line views/templates/tables/table_view.qtpl:256
			qw422016.N().S(`
                        onkeyup="return FilterIsModified( event, this);"
                        />
               `)
			//line views/templates/tables/table_view.qtpl:259
		}
		//line views/templates/tables/table_view.qtpl:259
		qw422016.N().S(`
        `)
		//line views/templates/tables/table_view.qtpl:260
	}
	//line views/templates/tables/table_view.qtpl:260
	qw422016.N().S(`
  </div>
`)
//line views/templates/tables/table_view.qtpl:262
}

//line views/templates/tables/table_view.qtpl:262
func writeinputFilterField(qq422016 qtio422016.Writer, key, nameTable string, fieldStruct *forms.FieldStructure) {
	//line views/templates/tables/table_view.qtpl:262
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/tables/table_view.qtpl:262
	streaminputFilterField(qw422016, key, nameTable, fieldStruct)
	//line views/templates/tables/table_view.qtpl:262
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_view.qtpl:262
}

//line views/templates/tables/table_view.qtpl:262
func inputFilterField(key, nameTable string, fieldStruct *forms.FieldStructure) string {
	//line views/templates/tables/table_view.qtpl:262
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/tables/table_view.qtpl:262
	writeinputFilterField(qb422016, key, nameTable, fieldStruct)
	//line views/templates/tables/table_view.qtpl:262
	qs422016 := string(qb422016.B)
	//line views/templates/tables/table_view.qtpl:262
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/tables/table_view.qtpl:262
	return qs422016
//line views/templates/tables/table_view.qtpl:262
}
