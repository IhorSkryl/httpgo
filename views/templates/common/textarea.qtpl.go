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

    <label class="c-app-textarea">
        <textarea class="c-textarea" id="`)
	//line views/templates/common/textarea.qtpl:5
	qw422016.E().S(name)
	//line views/templates/common/textarea.qtpl:5
	qw422016.N().S(`" name="`)
	//line views/templates/common/textarea.qtpl:5
	qw422016.E().S(name)
	//line views/templates/common/textarea.qtpl:5
	qw422016.N().S(`"></textarea>
        <span data-set-text></span>
    </label>

`)
//line views/templates/common/textarea.qtpl:9
}

//line views/templates/common/textarea.qtpl:9
func WriteRenderTextarea(qq422016 qtio422016.Writer, name string) {
	//line views/templates/common/textarea.qtpl:9
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/common/textarea.qtpl:9
	StreamRenderTextarea(qw422016, name)
	//line views/templates/common/textarea.qtpl:9
	qt422016.ReleaseWriter(qw422016)
//line views/templates/common/textarea.qtpl:9
}

//line views/templates/common/textarea.qtpl:9
func RenderTextarea(name string) string {
	//line views/templates/common/textarea.qtpl:9
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/common/textarea.qtpl:9
	WriteRenderTextarea(qb422016, name)
	//line views/templates/common/textarea.qtpl:9
	qs422016 := string(qb422016.B)
	//line views/templates/common/textarea.qtpl:9
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/common/textarea.qtpl:9
	return qs422016
//line views/templates/common/textarea.qtpl:9
}
