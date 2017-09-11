// This file is automatically generated by qtc from "record.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line views/templates/forms/record.qtpl:1
package forms

//line views/templates/forms/record.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// All the text outside function templates is treated as comments,
// i.e. it is just ignored by quicktemplate compiler (`qtc`). It is for humans.
//
// просто показ одной записи таблицы.

//line views/templates/forms/record.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/forms/record.qtpl:5
func (fields *FieldsTable) StreamShowRecord(qw422016 *qt422016.Writer, id string) {
	//line views/templates/forms/record.qtpl:5
	qw422016.N().S(`
    <div>
        <h2>Запись №`)
	//line views/templates/forms/record.qtpl:7
	qw422016.N().S(id)
	//line views/templates/forms/record.qtpl:7
	qw422016.N().S(` в таблице `)
	//line views/templates/forms/record.qtpl:7
	qw422016.N().S(fields.Comment)
	//line views/templates/forms/record.qtpl:7
	qw422016.N().S(`</h2>
        `)
	//line views/templates/forms/record.qtpl:8
	for idx, field := range fields.Rows {
		//line views/templates/forms/record.qtpl:8
		qw422016.N().S(`
        <div id="divField`)
		//line views/templates/forms/record.qtpl:9
		qw422016.N().D(idx)
		//line views/templates/forms/record.qtpl:9
		qw422016.N().S(`" class="row `)
		//line views/templates/forms/record.qtpl:9
		qw422016.E().S(field.CSSClass)
		//line views/templates/forms/record.qtpl:9
		qw422016.N().S(`"
            `)
		//line views/templates/forms/record.qtpl:10
		if field.IsHidden {
			//line views/templates/forms/record.qtpl:10
			qw422016.N().S(` style="display:none" `)
			//line views/templates/forms/record.qtpl:10
		}
		//line views/templates/forms/record.qtpl:10
		qw422016.N().S(`
        >
            `)
		//line views/templates/forms/record.qtpl:12
		qw422016.E().S(field.Value)
		//line views/templates/forms/record.qtpl:12
		qw422016.N().S(`
        </div>
        `)
		//line views/templates/forms/record.qtpl:14
	}
	//line views/templates/forms/record.qtpl:14
	qw422016.N().S(`

    </div>
`)
//line views/templates/forms/record.qtpl:17
}

//line views/templates/forms/record.qtpl:17
func (fields *FieldsTable) WriteShowRecord(qq422016 qtio422016.Writer, id string) {
	//line views/templates/forms/record.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/forms/record.qtpl:17
	fields.StreamShowRecord(qw422016, id)
	//line views/templates/forms/record.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/record.qtpl:17
}

//line views/templates/forms/record.qtpl:17
func (fields *FieldsTable) ShowRecord(id string) string {
	//line views/templates/forms/record.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/forms/record.qtpl:17
	fields.WriteShowRecord(qb422016, id)
	//line views/templates/forms/record.qtpl:17
	qs422016 := string(qb422016.B)
	//line views/templates/forms/record.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/forms/record.qtpl:17
	return qs422016
//line views/templates/forms/record.qtpl:17
}
