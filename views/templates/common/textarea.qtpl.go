// This file is automatically generated by qtc from "textarea.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line views/templates/common/textarea.qtpl:1
package common

//line views/templates/common/textarea.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/templates/common/textarea.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/common/textarea.qtpl:2
func StreamRenderTextarea(qw422016 *qt422016.Writer, name string) {
	//line views/templates/common/textarea.qtpl:2
	qw422016.N().S(`

    <textarea class="c-textarea" id="`)
	//line views/templates/common/textarea.qtpl:4
	qw422016.E().S(name)
	//line views/templates/common/textarea.qtpl:4
	qw422016.N().S(`" name="`)
	//line views/templates/common/textarea.qtpl:4
	qw422016.E().S(name)
	//line views/templates/common/textarea.qtpl:4
	qw422016.N().S(`"></textarea>

`)
//line views/templates/common/textarea.qtpl:6
}

//line views/templates/common/textarea.qtpl:6
func WriteRenderTextarea(qq422016 qtio422016.Writer, name string) {
	//line views/templates/common/textarea.qtpl:6
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/common/textarea.qtpl:6
	StreamRenderTextarea(qw422016, name)
	//line views/templates/common/textarea.qtpl:6
	qt422016.ReleaseWriter(qw422016)
//line views/templates/common/textarea.qtpl:6
}

//line views/templates/common/textarea.qtpl:6
func RenderTextarea(name string) string {
	//line views/templates/common/textarea.qtpl:6
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/common/textarea.qtpl:6
	WriteRenderTextarea(qb422016, name)
	//line views/templates/common/textarea.qtpl:6
	qs422016 := string(qb422016.B)
	//line views/templates/common/textarea.qtpl:6
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/common/textarea.qtpl:6
	return qs422016
//line views/templates/common/textarea.qtpl:6
}
