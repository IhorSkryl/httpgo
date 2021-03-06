// Code generated by qtc from "formLayouts.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// All the text outside function templates is treated as comments,
// i.e. it is just ignored by quicktemplate compiler (`qtc`). It is for humans.
//
// Обязательніе подключения для форм редактирования.

//line views/templates/layouts/formLayouts.qtpl:5
package layouts

//line views/templates/layouts/formLayouts.qtpl:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/templates/layouts/formLayouts.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/layouts/formLayouts.qtpl:5
func StreamPutHeadForm(qw422016 *qt422016.Writer) {
//line views/templates/layouts/formLayouts.qtpl:5
	qw422016.N().S(`
    <link rel="stylesheet" href="/fields.css" type="text/css"  media="screen">
`)
//line views/templates/layouts/formLayouts.qtpl:7
}

//line views/templates/layouts/formLayouts.qtpl:7
func WritePutHeadForm(qq422016 qtio422016.Writer) {
//line views/templates/layouts/formLayouts.qtpl:7
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/layouts/formLayouts.qtpl:7
	StreamPutHeadForm(qw422016)
//line views/templates/layouts/formLayouts.qtpl:7
	qt422016.ReleaseWriter(qw422016)
//line views/templates/layouts/formLayouts.qtpl:7
}

//line views/templates/layouts/formLayouts.qtpl:7
func PutHeadForm() string {
//line views/templates/layouts/formLayouts.qtpl:7
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/layouts/formLayouts.qtpl:7
	WritePutHeadForm(qb422016)
//line views/templates/layouts/formLayouts.qtpl:7
	qs422016 := string(qb422016.B)
//line views/templates/layouts/formLayouts.qtpl:7
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/layouts/formLayouts.qtpl:7
	return qs422016
//line views/templates/layouts/formLayouts.qtpl:7
}

//line views/templates/layouts/formLayouts.qtpl:9
func StreamPutEndForm(qw422016 *qt422016.Writer) {
//line views/templates/layouts/formLayouts.qtpl:9
	qw422016.N().S(`
`)
//line views/templates/layouts/formLayouts.qtpl:10
}

//line views/templates/layouts/formLayouts.qtpl:10
func WritePutEndForm(qq422016 qtio422016.Writer) {
//line views/templates/layouts/formLayouts.qtpl:10
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/layouts/formLayouts.qtpl:10
	StreamPutEndForm(qw422016)
//line views/templates/layouts/formLayouts.qtpl:10
	qt422016.ReleaseWriter(qw422016)
//line views/templates/layouts/formLayouts.qtpl:10
}

//line views/templates/layouts/formLayouts.qtpl:10
func PutEndForm() string {
//line views/templates/layouts/formLayouts.qtpl:10
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/layouts/formLayouts.qtpl:10
	WritePutEndForm(qb422016)
//line views/templates/layouts/formLayouts.qtpl:10
	qs422016 := string(qb422016.B)
//line views/templates/layouts/formLayouts.qtpl:10
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/layouts/formLayouts.qtpl:10
	return qs422016
//line views/templates/layouts/formLayouts.qtpl:10
}
