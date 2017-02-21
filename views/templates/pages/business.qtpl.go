// This file is automatically generated by qtc from "business.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line business.qtpl:1
package pages

//line business.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line business.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line business.qtpl:1
func StreamPutBusinessHead(qw422016 *qt422016.Writer) {
	//line business.qtpl:1
	qw422016.N().S(`
    <div class="main-form-wrap">
    <h1 class="business-title">Регистрация партнера на сервисе</h1>
    <h4 class="business-title-second">Добавьте свой бизнес для быстрого роста продаж</h4>
`)
//line business.qtpl:5
}

//line business.qtpl:5
func WritePutBusinessHead(qq422016 qtio422016.Writer) {
	//line business.qtpl:5
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line business.qtpl:5
	StreamPutBusinessHead(qw422016)
	//line business.qtpl:5
	qt422016.ReleaseWriter(qw422016)
//line business.qtpl:5
}

//line business.qtpl:5
func PutBusinessHead() string {
	//line business.qtpl:5
	qb422016 := qt422016.AcquireByteBuffer()
	//line business.qtpl:5
	WritePutBusinessHead(qb422016)
	//line business.qtpl:5
	qs422016 := string(qb422016.B)
	//line business.qtpl:5
	qt422016.ReleaseByteBuffer(qb422016)
	//line business.qtpl:5
	return qs422016
//line business.qtpl:5
}

//line business.qtpl:7
func StreamPutBusinessEnd(qw422016 *qt422016.Writer) {
	//line business.qtpl:7
	qw422016.N().S(`
</div>
`)
//line business.qtpl:9
}

//line business.qtpl:9
func WritePutBusinessEnd(qq422016 qtio422016.Writer) {
	//line business.qtpl:9
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line business.qtpl:9
	StreamPutBusinessEnd(qw422016)
	//line business.qtpl:9
	qt422016.ReleaseWriter(qw422016)
//line business.qtpl:9
}

//line business.qtpl:9
func PutBusinessEnd() string {
	//line business.qtpl:9
	qb422016 := qt422016.AcquireByteBuffer()
	//line business.qtpl:9
	WritePutBusinessEnd(qb422016)
	//line business.qtpl:9
	qs422016 := string(qb422016.B)
	//line business.qtpl:9
	qt422016.ReleaseByteBuffer(qb422016)
	//line business.qtpl:9
	return qs422016
//line business.qtpl:9
}
