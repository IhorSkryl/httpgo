// This file is automatically generated by qtc from "table_view.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line views/templates/pages/table_view.qtpl:1
package pages

//line views/templates/pages/table_view.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// All the text outside function templates is treated as comments,
// i.e. it is just ignored by quicktemplate compiler (`qtc`). It is for humans.
//
// показ табличных данных.

//line views/templates/pages/table_view.qtpl:5
import (
	"database/sql"
	"github.com/ruslanBik4/httpgo/views/templates/forms"
	"log"
	"strings"
)

//line views/templates/pages/table_view.qtpl:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/pages/table_view.qtpl:12
func StreamShowTable(qw422016 *qt422016.Writer, tableName string, fields *forms.FieldsTable, rows *sql.Rows) {
	//line views/templates/pages/table_view.qtpl:12
	qw422016.N().S(`
        `)
	//line views/templates/pages/table_view.qtpl:14
	columns, err := rows.Columns()
	if err != nil {
		log.Println(err)
	}

	var row []interface{}

	// mfields может не соответствовать набору столбцов, потому завязываем на имеющиеся, прочие - игнорируем
	for _, fieldName := range columns {
		var field interface{}

		if fieldStruct := fields.FindField(fieldName); fieldStruct != nil {
			field = fieldStruct
		} else {
			field = new(sql.NullString)
		}
		row = append(row, field)
	}

	//line views/templates/pages/table_view.qtpl:32
	qw422016.N().S(`
<link href="/tables.css" rel="stylesheet">

<table class="table table-striped table-bordered table-hover table-condensed">
    <thead>
        <tr>
            `)
	//line views/templates/pages/table_view.qtpl:38
	for idx, fieldName := range columns {
		//line views/templates/pages/table_view.qtpl:38
		qw422016.N().S(`
                <td>`)
		//line views/templates/pages/table_view.qtpl:40
		fieldStruct := fields.FindField(fieldName)
		titleFull, titleLabel, _, _, _ := fieldStruct.GetColumnTitles()

		//line views/templates/pages/table_view.qtpl:43
		qw422016.N().S(`
                    <a href="#`)
		//line views/templates/pages/table_view.qtpl:44
		qw422016.N().D(idx)
		//line views/templates/pages/table_view.qtpl:44
		qw422016.N().S(`" title="`)
		//line views/templates/pages/table_view.qtpl:44
		qw422016.E().S(titleFull)
		//line views/templates/pages/table_view.qtpl:44
		qw422016.N().S(`">`)
		//line views/templates/pages/table_view.qtpl:44
		qw422016.E().S(titleLabel)
		//line views/templates/pages/table_view.qtpl:44
		qw422016.N().S(`</a>
                </td>
            `)
		//line views/templates/pages/table_view.qtpl:46
	}
	//line views/templates/pages/table_view.qtpl:46
	qw422016.N().S(`
        </tr>
    </thead>
    <tbody>
        `)
	//line views/templates/pages/table_view.qtpl:50
	for rows.Next() {
		//line views/templates/pages/table_view.qtpl:50
		qw422016.N().S(`
            `)
		//line views/templates/pages/table_view.qtpl:52
		err = rows.Scan(row...)

		if err != nil {
			log.Println(err)
			continue
		}

		//line views/templates/pages/table_view.qtpl:59
		qw422016.N().S(`
        <tr>
            `)
		//line views/templates/pages/table_view.qtpl:61
		for idx, rawField := range row {
			//line views/templates/pages/table_view.qtpl:61
			qw422016.N().S(`
                `)
			//line views/templates/pages/table_view.qtpl:63
			var fieldStruct *forms.FieldStructure

			switch rawField.(type) {
			case *forms.FieldStructure:
				fieldStruct = rawField.(*forms.FieldStructure)
			default:
				continue
			}
			key := fieldStruct.COLUMN_NAME
			val := fieldStruct.Value
			required, tablePrefix, titleLabel := "", "", ""

			//line views/templates/pages/table_view.qtpl:74
			qw422016.N().S(`
                `)
			//line views/templates/pages/table_view.qtpl:75
			if idx == 0 {
				//line views/templates/pages/table_view.qtpl:75
				qw422016.N().S(`
                    <td><a href="/admin/row/edit/?table=`)
				//line views/templates/pages/table_view.qtpl:76
				qw422016.E().S(tableName)
				//line views/templates/pages/table_view.qtpl:76
				qw422016.N().S(`&id=`)
				//line views/templates/pages/table_view.qtpl:76
				qw422016.E().S(val)
				//line views/templates/pages/table_view.qtpl:76
				qw422016.N().S(`" target="content">`)
				//line views/templates/pages/table_view.qtpl:76
				qw422016.E().S(val)
				//line views/templates/pages/table_view.qtpl:76
				qw422016.N().S(`</a>
                    </td>
                `)
				//line views/templates/pages/table_view.qtpl:78
			} else {
				//line views/templates/pages/table_view.qtpl:78
				qw422016.N().S(`
                    <td>
                        `)
				//line views/templates/pages/table_view.qtpl:80
				if key == "parent_id" {
					//line views/templates/pages/table_view.qtpl:80
					qw422016.N().S(`
                            `)
					//line views/templates/pages/table_view.qtpl:81
					fieldStruct.StreamRenderParentSelect(qw422016, fields.Name, key, val, titleLabel, required)
					//line views/templates/pages/table_view.qtpl:81
					qw422016.N().S(`
                        `)
					//line views/templates/pages/table_view.qtpl:82
				} else if strings.HasPrefix(key, "id_") {
					//line views/templates/pages/table_view.qtpl:82
					qw422016.N().S(`
                            `)
					//line views/templates/pages/table_view.qtpl:83
					fieldStruct.StreamRenderForeignSelect(qw422016, tablePrefix, key, val, titleLabel, required)
					//line views/templates/pages/table_view.qtpl:83
					qw422016.N().S(`
                        `)
					//line views/templates/pages/table_view.qtpl:84
				} else if strings.HasPrefix(key, "setid_") {
					//line views/templates/pages/table_view.qtpl:84
					qw422016.N().S(`
                            `)
					//line views/templates/pages/table_view.qtpl:85
					fieldStruct.StreamRenderMultiSelect(qw422016, fields, tablePrefix, key, val, "См. ", required)
					//line views/templates/pages/table_view.qtpl:85
					qw422016.N().S(`
                        `)
					//line views/templates/pages/table_view.qtpl:86
				} else {
					//line views/templates/pages/table_view.qtpl:86
					qw422016.N().S(`
                            `)
					//line views/templates/pages/table_view.qtpl:87
					qw422016.E().S(val)
					//line views/templates/pages/table_view.qtpl:87
					qw422016.N().S(`
                        `)
					//line views/templates/pages/table_view.qtpl:88
				}
				//line views/templates/pages/table_view.qtpl:88
				qw422016.N().S(`
                    </td>
                `)
				//line views/templates/pages/table_view.qtpl:90
			}
			//line views/templates/pages/table_view.qtpl:90
			qw422016.N().S(`
            `)
			//line views/templates/pages/table_view.qtpl:91
		}
		//line views/templates/pages/table_view.qtpl:91
		qw422016.N().S(`

        </tr>
        `)
		//line views/templates/pages/table_view.qtpl:94
	}
	//line views/templates/pages/table_view.qtpl:94
	qw422016.N().S(`
    </tbody>

</table>

`)
//line views/templates/pages/table_view.qtpl:99
}

//line views/templates/pages/table_view.qtpl:99
func WriteShowTable(qq422016 qtio422016.Writer, tableName string, fields *forms.FieldsTable, rows *sql.Rows) {
	//line views/templates/pages/table_view.qtpl:99
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/pages/table_view.qtpl:99
	StreamShowTable(qw422016, tableName, fields, rows)
	//line views/templates/pages/table_view.qtpl:99
	qt422016.ReleaseWriter(qw422016)
//line views/templates/pages/table_view.qtpl:99
}

//line views/templates/pages/table_view.qtpl:99
func ShowTable(tableName string, fields *forms.FieldsTable, rows *sql.Rows) string {
	//line views/templates/pages/table_view.qtpl:99
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/pages/table_view.qtpl:99
	WriteShowTable(qb422016, tableName, fields, rows)
	//line views/templates/pages/table_view.qtpl:99
	qs422016 := string(qb422016.B)
	//line views/templates/pages/table_view.qtpl:99
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/pages/table_view.qtpl:99
	return qs422016
//line views/templates/pages/table_view.qtpl:99
}
