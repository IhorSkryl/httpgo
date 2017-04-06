// This file is automatically generated by qtc from "index.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line index.qtpl:1
package pages

//line index.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// All the text outside function templates is treated as comments,
// i.e. it is just ignored by quicktemplate compiler (`qtc`). It is for humans.

//line index.qtpl:3
import (
	"github.com/ruslanBik4/httpgo/views/templates/layouts"
)

// це главная страница.

//line index.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line index.qtpl:9
type ItemMenu struct {
	Link string
}
type IndexPageBody struct {
	Name    []byte
	Pass    []byte
	Content string
	Catalog []string
	TopMenu map[string]string
	Title   string
	Route   string
}

//line index.qtpl:22
func (body *IndexPageBody) StreamIndexHTML(qw422016 *qt422016.Writer) {
	//line index.qtpl:22
	qw422016.N().S(`
<body>
<div class="content-wrap">
    `)
	//line index.qtpl:25
	layouts.StreamHeaderHTML(qw422016, body.TopMenu)
	//line index.qtpl:25
	qw422016.N().S(`
<div id="container-fluid">
    <div class="row-fluid">
        <div class="sidebar-section">
            <div id="catalog_pane"  class="well sidebar-nav">
                `)
	//line index.qtpl:30
	for _, value := range body.Catalog {
		//line index.qtpl:30
		qw422016.N().S(`
                    `)
		//line index.qtpl:31
		qw422016.E().S(value)
		//line index.qtpl:31
		qw422016.N().S(`
                `)
		//line index.qtpl:32
	}
	//line index.qtpl:32
	qw422016.N().S(`

            </div>
        </div>
        <div class="content-section">
            <div id="content" rel="`)
	//line index.qtpl:37
	qw422016.E().S(body.Route)
	//line index.qtpl:37
	qw422016.N().S(`">`)
	//line index.qtpl:37
	qw422016.N().S(body.Content)
	//line index.qtpl:37
	qw422016.N().S(`</div>
        </div>
    </div>
</div>
</div>

    `)
	//line index.qtpl:43
	layouts.StreamFooterHTML(qw422016)
	//line index.qtpl:43
	qw422016.N().S(`

</body>
`)
//line index.qtpl:46
}

//line index.qtpl:46
func (body *IndexPageBody) WriteIndexHTML(qq422016 qtio422016.Writer) {
	//line index.qtpl:46
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line index.qtpl:46
	body.StreamIndexHTML(qw422016)
	//line index.qtpl:46
	qt422016.ReleaseWriter(qw422016)
//line index.qtpl:46
}

//line index.qtpl:46
func (body *IndexPageBody) IndexHTML() string {
	//line index.qtpl:46
	qb422016 := qt422016.AcquireByteBuffer()
	//line index.qtpl:46
	body.WriteIndexHTML(qb422016)
	//line index.qtpl:46
	qs422016 := string(qb422016.B)
	//line index.qtpl:46
	qt422016.ReleaseByteBuffer(qb422016)
	//line index.qtpl:46
	return qs422016
//line index.qtpl:46
}
