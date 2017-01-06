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
)

//line views/templates/pages/table_view.qtpl:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/pages/table_view.qtpl:11
func StreamShowTable(qw422016 *qt422016.Writer, tableName string, ns forms.FieldsTable, rows *sql.Rows) {
	//line views/templates/pages/table_view.qtpl:11
	qw422016.N().S(`
<link href="/tables.css" rel="stylesheet">

<table>
    <thead>
        <tr>
            `)
	//line views/templates/pages/table_view.qtpl:17
	for _, field := range ns.Rows {
		//line views/templates/pages/table_view.qtpl:17
		qw422016.N().S(`
                <td>`)
		//line views/templates/pages/table_view.qtpl:19
		titleFull, titleLabel, _, _ := field.GetColumnTitles()
		if len(titleLabel) > 22 {
			titleLabel = titleLabel[22:] + "..."
		}

		//line views/templates/pages/table_view.qtpl:23
		qw422016.N().S(`
                    <a href="#" title="`)
		//line views/templates/pages/table_view.qtpl:24
		qw422016.N().S(titleFull)
		//line views/templates/pages/table_view.qtpl:24
		qw422016.N().S(`">`)
		//line views/templates/pages/table_view.qtpl:24
		qw422016.N().S(titleLabel)
		//line views/templates/pages/table_view.qtpl:24
		qw422016.N().S(`</a>
                </td>
            `)
		//line views/templates/pages/table_view.qtpl:26
	}
	//line views/templates/pages/table_view.qtpl:26
	qw422016.N().S(`
        </tr>
    </thead>
    <tbody>
        `)
	//line views/templates/pages/table_view.qtpl:31
	columns, err := rows.Columns()
	if err != nil {
		log.Println(err)
	}

	var row []interface{}

	rowField := make([]*sql.NullString, len(columns))

	for idx, _ := range columns {

		rowField[idx] = new(sql.NullString)
		row = append(row, rowField[idx])
	}

	//line views/templates/pages/table_view.qtpl:45
	qw422016.N().S(`
        `)
	//line views/templates/pages/table_view.qtpl:46
	for rows.Next() {
		//line views/templates/pages/table_view.qtpl:46
		qw422016.N().S(`
            `)
		//line views/templates/pages/table_view.qtpl:48
		err = rows.Scan(row...)

		if err != nil {
			log.Println(err)
			continue
		}

		//line views/templates/pages/table_view.qtpl:55
		qw422016.N().S(`
        <tr>
            `)
		//line views/templates/pages/table_view.qtpl:57
		for idx, field := range rowField {
			//line views/templates/pages/table_view.qtpl:57
			qw422016.N().S(`
                `)
			//line views/templates/pages/table_view.qtpl:58
			if idx == 0 {
				//line views/templates/pages/table_view.qtpl:58
				qw422016.N().S(`
                    <td><a href="/admin/row/edit/?table=`)
				//line views/templates/pages/table_view.qtpl:59
				qw422016.E().S(tableName)
				//line views/templates/pages/table_view.qtpl:59
				qw422016.N().S(`&id=`)
				//line views/templates/pages/table_view.qtpl:59
				qw422016.E().S(field.String)
				//line views/templates/pages/table_view.qtpl:59
				qw422016.N().S(`" target="content">`)
				//line views/templates/pages/table_view.qtpl:59
				qw422016.E().S(field.String)
				//line views/templates/pages/table_view.qtpl:59
				qw422016.N().S(`</a>
                    </td>
                `)
				//line views/templates/pages/table_view.qtpl:61
			} else {
				//line views/templates/pages/table_view.qtpl:61
				qw422016.N().S(`
                    <td>`)
				//line views/templates/pages/table_view.qtpl:62
				if field.Valid {
					//line views/templates/pages/table_view.qtpl:62
					qw422016.E().S(field.String)
					//line views/templates/pages/table_view.qtpl:62
				}
				//line views/templates/pages/table_view.qtpl:62
				qw422016.N().S(`</td>
                `)
				//line views/templates/pages/table_view.qtpl:63
			}
			//line views/templates/pages/table_view.qtpl:63
			qw422016.N().S(`
            `)
			//line views/templates/pages/table_view.qtpl:64
		}
		//line views/templates/pages/table_view.qtpl:64
		qw422016.N().S(`

        </tr>
        `)
		//line views/templates/pages/table_view.qtpl:67
	}
	//line views/templates/pages/table_view.qtpl:67
	qw422016.N().S(`
    </tbody>

</table>

`)
//line views/templates/pages/table_view.qtpl:72
}

//line views/templates/pages/table_view.qtpl:72
func WriteShowTable(qq422016 qtio422016.Writer, tableName string, ns forms.FieldsTable, rows *sql.Rows) {
	//line views/templates/pages/table_view.qtpl:72
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/pages/table_view.qtpl:72
	StreamShowTable(qw422016, tableName, ns, rows)
	//line views/templates/pages/table_view.qtpl:72
	qt422016.ReleaseWriter(qw422016)
//line views/templates/pages/table_view.qtpl:72
}

//line views/templates/pages/table_view.qtpl:72
func ShowTable(tableName string, ns forms.FieldsTable, rows *sql.Rows) string {
	//line views/templates/pages/table_view.qtpl:72
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/pages/table_view.qtpl:72
	WriteShowTable(qb422016, tableName, ns, rows)
	//line views/templates/pages/table_view.qtpl:72
	qs422016 := string(qb422016.B)
	//line views/templates/pages/table_view.qtpl:72
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/pages/table_view.qtpl:72
	return qs422016
//line views/templates/pages/table_view.qtpl:72
}
