// This file is automatically generated by qtc from "input.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line views/templates/common/input.qtpl:1
package common

//line views/templates/common/input.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// {#
//     attr[0] - name input, id input
//     attr[1] - data-form id form
// #}
//

//line views/templates/common/input.qtpl:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/common/input.qtpl:7
func StreamRenderInput(qw422016 *qt422016.Writer, attr ...string) {
	//line views/templates/common/input.qtpl:7
	qw422016.N().S(`

    `)
	//line views/templates/common/input.qtpl:10
	dataForm := ""
	if len(attr) > 1 {
		dataForm = "${ Variables.paramsFormChildren }=" + attr[1] + "-${ data.idForm } ${ Variables.paramsJSONIdData }=${ data.idForm }"
	}

	//line views/templates/common/input.qtpl:14
	qw422016.N().S(`

    <label class="c-app-input">
        <input class="c-input" type="text" id="`)
	//line views/templates/common/input.qtpl:17
	qw422016.E().S(attr[0])
	//line views/templates/common/input.qtpl:17
	qw422016.N().S(`" name="`)
	//line views/templates/common/input.qtpl:17
	qw422016.E().S(attr[0])
	//line views/templates/common/input.qtpl:17
	qw422016.N().S(`" `)
	//line views/templates/common/input.qtpl:17
	qw422016.E().S(dataForm)
	//line views/templates/common/input.qtpl:17
	qw422016.N().S(`>
        <span data-set-text></span>
    </label>

`)
//line views/templates/common/input.qtpl:21
}

//line views/templates/common/input.qtpl:21
func WriteRenderInput(qq422016 qtio422016.Writer, attr ...string) {
	//line views/templates/common/input.qtpl:21
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/common/input.qtpl:21
	StreamRenderInput(qw422016, attr...)
	//line views/templates/common/input.qtpl:21
	qt422016.ReleaseWriter(qw422016)
//line views/templates/common/input.qtpl:21
}

//line views/templates/common/input.qtpl:21
func RenderInput(attr ...string) string {
	//line views/templates/common/input.qtpl:21
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/common/input.qtpl:21
	WriteRenderInput(qb422016, attr...)
	//line views/templates/common/input.qtpl:21
	qs422016 := string(qb422016.B)
	//line views/templates/common/input.qtpl:21
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/common/input.qtpl:21
	return qs422016
//line views/templates/common/input.qtpl:21
}
