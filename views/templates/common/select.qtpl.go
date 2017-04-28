// This file is automatically generated by qtc from "select.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:1
package common

//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:2
func StreamRenderSelect(qw422016 *qt422016.Writer, name string) {
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:2
	qw422016.N().S(`

    <select class='c-select' id='`)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:4
	qw422016.E().S(name)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:4
	qw422016.N().S(`' name='`)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:4
	qw422016.E().S(name)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:4
	qw422016.N().S(`'>
        <option value=""></option>
    </select>

`)
//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:8
}

//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:8
func WriteRenderSelect(qq422016 qtio422016.Writer, name string) {
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:8
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:8
	StreamRenderSelect(qw422016, name)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:8
	qt422016.ReleaseWriter(qw422016)
//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:8
}

//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:8
func RenderSelect(name string) string {
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:8
	qb422016 := qt422016.AcquireByteBuffer()
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:8
	WriteRenderSelect(qb422016, name)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:8
	qs422016 := string(qb422016.B)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:8
	qt422016.ReleaseByteBuffer(qb422016)
	//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:8
	return qs422016
//line src/github.com/ruslanBik4/httpgo/views/templates/common/select.qtpl:8
}
