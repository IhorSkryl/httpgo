// This file is automatically generated by qtc from "checkbox.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:1
package common

//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:2
func StreamRenderCheckbox(qw422016 *qt422016.Writer, name string) {
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:2
	qw422016.N().S(`

<div class="c-app-checkbox">
    <input class="c-checkbox" type="checkbox" id="`)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:5
	qw422016.E().S(name)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:5
	qw422016.N().S(`" name="`)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:5
	qw422016.E().S(name)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:5
	qw422016.N().S(`">
    <label for="`)
	//line views/templates/common/checkbox.qtpl:6
	qw422016.E().S(name)
	//line views/templates/common/checkbox.qtpl:6
	qw422016.N().S(`"></label>
</div>

`)
//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:9
}

//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:9
func WriteRenderCheckbox(qq422016 qtio422016.Writer, name string) {
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:9
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:9
	StreamRenderCheckbox(qw422016, name)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:9
	qt422016.ReleaseWriter(qw422016)
//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:9
}

//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:9
func RenderCheckbox(name string) string {
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:9
	qb422016 := qt422016.AcquireByteBuffer()
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:9
	WriteRenderCheckbox(qb422016, name)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:9
	qs422016 := string(qb422016.B)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:9
	qt422016.ReleaseByteBuffer(qb422016)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:9
	return qs422016
//line src/github.com/ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:9
}
