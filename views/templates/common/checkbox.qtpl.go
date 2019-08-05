// This file is automatically generated by qtc from "checkbox.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:1
package common

//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// {#
//     attr[0] - name input, id input
//     attr[1] - data-form id form
// #}
//

//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:7
func StreamRenderCheckbox(qw422016 *qt422016.Writer, attr ...string) {
	//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:7
	qw422016.N().S(`

    `)
	//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:10
	dataForm := ""
	if len(attr) > 1 {
		dataForm = "${ Variables.paramsFormChildren }=" + attr[1] + "-${ data.idForm } ${ Variables.paramsJSONIdData }=${ data.idForm }"
	}

	//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:14
	qw422016.N().S(`

    <div class="c-app-checkbox">
        <label>
            <input class="c-checkbox" type="checkbox" id="`)
	//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:18
	qw422016.E().S(attr[0])
	//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:18
	qw422016.N().S(`" name="`)
	//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:18
	qw422016.E().S(attr[0])
	//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:18
	qw422016.N().S(`" `)
	//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:18
	qw422016.E().S(dataForm)
	//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:18
	qw422016.N().S(` value="1">
            <span data-set-text></span>
        </label>
    </div>

`)
//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:23
}

//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:23
func WriteRenderCheckbox(qq422016 qtio422016.Writer, attr ...string) {
	//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:23
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:23
	StreamRenderCheckbox(qw422016, attr...)
	//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:23
	qt422016.ReleaseWriter(qw422016)
//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:23
}

//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:23
func RenderCheckbox(attr ...string) string {
	//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:23
	qb422016 := qt422016.AcquireByteBuffer()
	//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:23
	WriteRenderCheckbox(qb422016, attr...)
	//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:23
	qs422016 := string(qb422016.B)
	//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:23
	qt422016.ReleaseByteBuffer(qb422016)
	//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:23
	return qs422016
//line ../../ruslanBik4/httpgo/views/templates/common/checkbox.qtpl:23
}
