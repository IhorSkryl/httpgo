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
	"../forms"
	"database/sql"
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
<link href="/css/tables.css" rel="stylesheet">

<table>
    <thead>
        <tr>
            `)
	//line views/templates/pages/table_view.qtpl:17
	for _, value := range ns.Rows {
		//line views/templates/pages/table_view.qtpl:17
		qw422016.N().S(`
                <td>`)
		//line views/templates/pages/table_view.qtpl:18
		if value.COLUMN_COMMENT > "" {
			//line views/templates/pages/table_view.qtpl:18
			qw422016.E().S(value.COLUMN_COMMENT)
			//line views/templates/pages/table_view.qtpl:18
		} else {
			//line views/templates/pages/table_view.qtpl:18
			qw422016.E().S(value.COLUMN_NAME)
			//line views/templates/pages/table_view.qtpl:18
		}
		//line views/templates/pages/table_view.qtpl:18
		qw422016.N().S(`</td>
            `)
		//line views/templates/pages/table_view.qtpl:19
	}
	//line views/templates/pages/table_view.qtpl:19
	qw422016.N().S(`
        </tr>
    </thead>
    <tbody>
        `)
	//line views/templates/pages/table_view.qtpl:24
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

	//line views/templates/pages/table_view.qtpl:38
	qw422016.N().S(`
        `)
	//line views/templates/pages/table_view.qtpl:39
	for rows.Next() {
		//line views/templates/pages/table_view.qtpl:39
		qw422016.N().S(`
            `)
		//line views/templates/pages/table_view.qtpl:41
		err = rows.Scan(row...)

		if err != nil {
			log.Println(err)
			continue
		}

		//line views/templates/pages/table_view.qtpl:48
		qw422016.N().S(`
        <tr>
            `)
		//line views/templates/pages/table_view.qtpl:50
		for idx, field := range rowField {
			//line views/templates/pages/table_view.qtpl:50
			qw422016.N().S(`
                `)
			//line views/templates/pages/table_view.qtpl:51
			if idx == 0 {
				//line views/templates/pages/table_view.qtpl:51
				qw422016.N().S(`
                    <td><a href="/admin/row/edit/?table=`)
				//line views/templates/pages/table_view.qtpl:52
				qw422016.E().S(tableName)
				//line views/templates/pages/table_view.qtpl:52
				qw422016.N().S(`&id=`)
				//line views/templates/pages/table_view.qtpl:52
				qw422016.E().S(field.String)
				//line views/templates/pages/table_view.qtpl:52
				qw422016.N().S(`" target="content">`)
				//line views/templates/pages/table_view.qtpl:52
				qw422016.E().S(field.String)
				//line views/templates/pages/table_view.qtpl:52
				qw422016.N().S(`</a>
                    </td>
                `)
				//line views/templates/pages/table_view.qtpl:54
			} else {
				//line views/templates/pages/table_view.qtpl:54
				qw422016.N().S(`
                    <td>`)
				//line views/templates/pages/table_view.qtpl:55
				if field.Valid {
					//line views/templates/pages/table_view.qtpl:55
					qw422016.E().S(field.String)
					//line views/templates/pages/table_view.qtpl:55
				}
				//line views/templates/pages/table_view.qtpl:55
				qw422016.N().S(`</td>
                `)
				//line views/templates/pages/table_view.qtpl:56
			}
			//line views/templates/pages/table_view.qtpl:56
			qw422016.N().S(`
            `)
			//line views/templates/pages/table_view.qtpl:57
		}
		//line views/templates/pages/table_view.qtpl:57
		qw422016.N().S(`

        </tr>
        `)
		//line views/templates/pages/table_view.qtpl:60
	}
	//line views/templates/pages/table_view.qtpl:60
	qw422016.N().S(`
    </tbody>

</table>

`)
//line views/templates/pages/table_view.qtpl:65
}

//line views/templates/pages/table_view.qtpl:65
func WriteShowTable(qq422016 qtio422016.Writer, tableName string, ns forms.FieldsTable, rows *sql.Rows) {
	//line views/templates/pages/table_view.qtpl:65
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/pages/table_view.qtpl:65
	StreamShowTable(qw422016, tableName, ns, rows)
	//line views/templates/pages/table_view.qtpl:65
	qt422016.ReleaseWriter(qw422016)
//line views/templates/pages/table_view.qtpl:65
}

//line views/templates/pages/table_view.qtpl:65
func ShowTable(tableName string, ns forms.FieldsTable, rows *sql.Rows) string {
	//line views/templates/pages/table_view.qtpl:65
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/pages/table_view.qtpl:65
	WriteShowTable(qb422016, tableName, ns, rows)
	//line views/templates/pages/table_view.qtpl:65
	qs422016 := string(qb422016.B)
	//line views/templates/pages/table_view.qtpl:65
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/pages/table_view.qtpl:65
	return qs422016
//line views/templates/pages/table_view.qtpl:65
}
