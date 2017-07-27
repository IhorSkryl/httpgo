// This file is automatically generated by qtc from "system.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line views/templates/system/system.qtpl:1
package system

//line views/templates/system/system.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// All the text outside function templates is treated as comments,
// i.e. it is just ignored by quicktemplate compiler (`qtc`). It is for humans.
//
// предназначен для оформления выдачи сообщений управления веб-сервером.
// route - имя пути, на который надо сделать запрос и полученный ответ показать на странице

//line views/templates/system/system.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/system/system.qtpl:6
func StreamAddRescanJS(qw422016 *qt422016.Writer, route string) {
	//line views/templates/system/system.qtpl:6
	qw422016.N().S(`
    <script>
    route string
    </script>
    <div id=addString >
    </div>
`)
//line views/templates/system/system.qtpl:12
}

//line views/templates/system/system.qtpl:12
func WriteAddRescanJS(qq422016 qtio422016.Writer, route string) {
	//line views/templates/system/system.qtpl:12
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/system/system.qtpl:12
	StreamAddRescanJS(qw422016, route)
	//line views/templates/system/system.qtpl:12
	qt422016.ReleaseWriter(qw422016)
//line views/templates/system/system.qtpl:12
}

//line views/templates/system/system.qtpl:12
func AddRescanJS(route string) string {
	//line views/templates/system/system.qtpl:12
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/system/system.qtpl:12
	WriteAddRescanJS(qb422016, route)
	//line views/templates/system/system.qtpl:12
	qs422016 := string(qb422016.B)
	//line views/templates/system/system.qtpl:12
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/system/system.qtpl:12
	return qs422016
//line views/templates/system/system.qtpl:12
}
