// This file is automatically generated by qtc from "radio.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:1
package common

//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:2
func StreamRenderRadio(qw422016 *qt422016.Writer, InputType, name, class, id, value string) {
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:2
	qw422016.N().S(`

<label class="c-radio-label"><input name="`)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:4
	qw422016.E().S(name)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:4
	qw422016.N().S(`" value="`)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:4
	qw422016.E().S(value)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:4
	qw422016.N().S(`" id="`)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:4
	qw422016.E().S(id)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:4
	qw422016.N().S(`" class="`)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:4
	qw422016.E().S(class)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:4
	qw422016.N().S(`" type="radio"><span>`)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:4
	qw422016.E().S(value)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:4
	qw422016.N().S(`</span></label>

`)
//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:6
}

//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:6
func WriteRenderRadio(qq422016 qtio422016.Writer, InputType, name, class, id, value string) {
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:6
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:6
	StreamRenderRadio(qw422016, InputType, name, class, id, value)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:6
	qt422016.ReleaseWriter(qw422016)
//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:6
}

//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:6
func RenderRadio(InputType, name, class, id, value string) string {
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:6
	qb422016 := qt422016.AcquireByteBuffer()
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:6
	WriteRenderRadio(qb422016, InputType, name, class, id, value)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:6
	qs422016 := string(qb422016.B)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:6
	qt422016.ReleaseByteBuffer(qb422016)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:6
	return qs422016
//line src/github.com/ruslanBik4/httpgo/views/templates/common/radio.qtpl:6
}
