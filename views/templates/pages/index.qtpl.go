// This file is automatically generated by qtc from "index.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line views/templates/pages/index.qtpl:1
package pages

//line views/templates/pages/index.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// All the text outside function templates is treated as comments,
// i.e. it is just ignored by quicktemplate compiler (`qtc`). It is for humans.

//line views/templates/pages/index.qtpl:3
import (
	"github.com/ruslanBik4/httpgo/views/templates/layouts"
	"io"
)

// це главная страница.

//line views/templates/pages/index.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/pages/index.qtpl:10
type ItemMenu struct {
	Link string
}
type IndexPageBody struct {
	Name         []byte
	Pass         []byte
	Content      string
	ContentWrite func(w io.Writer)
	Buff         io.Writer
	Catalog      []string
	TopMenu      map[string]string
	Title        string
	Route        string
}

//line views/templates/pages/index.qtpl:25
func (body *IndexPageBody) StreamIndexHTML(qw422016 *qt422016.Writer) {
	//line views/templates/pages/index.qtpl:25
	qw422016.N().S(`
<body>
<div class="content-wrap">
    `)
	//line views/templates/pages/index.qtpl:28
	layouts.StreamHeaderHTML(qw422016, body.TopMenu)
	//line views/templates/pages/index.qtpl:28
	qw422016.N().S(`
<div id="container-fluid">
    <div class="row-fluid">
        <div class="sidebar-section">
            <div id="catalog_pane"  class="well sidebar-nav">
                    `)
	//line views/templates/pages/index.qtpl:33
	qw422016.N().V(body.Catalog)
	//line views/templates/pages/index.qtpl:33
	qw422016.N().S(`
            </div>
        </div>
        <div class="content-section">
            <div id="content" rel="`)
	//line views/templates/pages/index.qtpl:37
	qw422016.E().S(body.Route)
	//line views/templates/pages/index.qtpl:37
	qw422016.N().S(`">
                `)
	//line views/templates/pages/index.qtpl:38
	body.ContentWrite(body.Buff)

	//line views/templates/pages/index.qtpl:38
	qw422016.N().S(`
            </div>
        </div>
    </div>
</div>
</div>

    `)
	//line views/templates/pages/index.qtpl:45
	layouts.StreamFooterHTML(qw422016)
	//line views/templates/pages/index.qtpl:45
	qw422016.N().S(`

</body>
`)
//line views/templates/pages/index.qtpl:48
}

//line views/templates/pages/index.qtpl:48
func (body *IndexPageBody) WriteIndexHTML(qq422016 qtio422016.Writer) {
	//line views/templates/pages/index.qtpl:48
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/pages/index.qtpl:48
	body.StreamIndexHTML(qw422016)
	//line views/templates/pages/index.qtpl:48
	qt422016.ReleaseWriter(qw422016)
//line views/templates/pages/index.qtpl:48
}

//line views/templates/pages/index.qtpl:48
func (body *IndexPageBody) IndexHTML() string {
	//line views/templates/pages/index.qtpl:48
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/pages/index.qtpl:48
	body.WriteIndexHTML(qb422016)
	//line views/templates/pages/index.qtpl:48
	qs422016 := string(qb422016.B)
	//line views/templates/pages/index.qtpl:48
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/pages/index.qtpl:48
	return qs422016
//line views/templates/pages/index.qtpl:48
}
