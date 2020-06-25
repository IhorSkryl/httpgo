// Code generated by qtc from "admin_index.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// All the text outside function templates is treated as comments,
// i.e. it is just ignored by quicktemplate compiler (`qtc`). It is for humans.

//line views/templates/pages/admin_index.qtpl:3
package pages

//line views/templates/pages/admin_index.qtpl:3
import (
	"github.com/ruslanBik4/httpgo/views/templates/layouts"
)

// Index page for administrator.

//line views/templates/pages/admin_index.qtpl:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/templates/pages/admin_index.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/pages/admin_index.qtpl:10
type AdminPageBody struct {
	Name    []byte
	Pass    []byte
	Content string
	Catalog map[string]*ItemMenu
	TopMenu map[string]string
	Title   string
	Head    *layouts.HeadHTMLPage
}

//line views/templates/pages/admin_index.qtpl:20
func (body *AdminPageBody) StreamShowAdminPage(qw422016 *qt422016.Writer, activePage string) {
//line views/templates/pages/admin_index.qtpl:20
	qw422016.N().S(`
`)
//line views/templates/pages/admin_index.qtpl:21
	body.Head.StreamHeadHTML(qw422016)
//line views/templates/pages/admin_index.qtpl:21
	qw422016.N().S(`
<body>
<div class="content-wrap">
<div id="container-fluid">

    <div class="row-fluid">
        <div class="sidebar-section">
            <div id="catalog_pane"  class="well sidebar-nav">
                `)
//line views/templates/pages/admin_index.qtpl:29
	for name, item := range body.Catalog {
//line views/templates/pages/admin_index.qtpl:29
		qw422016.N().S(`
                    <li><a href="`)
//line views/templates/pages/admin_index.qtpl:30
		qw422016.E().S(item.Link)
//line views/templates/pages/admin_index.qtpl:30
		qw422016.N().S(`">`)
//line views/templates/pages/admin_index.qtpl:30
		qw422016.E().S(name)
//line views/templates/pages/admin_index.qtpl:30
		qw422016.N().S(`</a></li>
                `)
//line views/templates/pages/admin_index.qtpl:31
	}
//line views/templates/pages/admin_index.qtpl:31
	qw422016.N().S(`

            </div>
        </div>
        <div id="content" rel="/admin/" class="content-section">`)
//line views/templates/pages/admin_index.qtpl:35
	qw422016.N().S(body.Content)
//line views/templates/pages/admin_index.qtpl:35
	qw422016.N().S(`</div>

    </div>
</div>
</div>

`)
//line views/templates/pages/admin_index.qtpl:41
	layouts.StreamFooterHTML(qw422016, "test@email.net", "8-000-00-00", "+380(00)000-00-00")
//line views/templates/pages/admin_index.qtpl:41
	qw422016.N().S(`
</body>
`)
//line views/templates/pages/admin_index.qtpl:43
}

//line views/templates/pages/admin_index.qtpl:43
func (body *AdminPageBody) WriteShowAdminPage(qq422016 qtio422016.Writer, activePage string) {
//line views/templates/pages/admin_index.qtpl:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/pages/admin_index.qtpl:43
	body.StreamShowAdminPage(qw422016, activePage)
//line views/templates/pages/admin_index.qtpl:43
	qt422016.ReleaseWriter(qw422016)
//line views/templates/pages/admin_index.qtpl:43
}

//line views/templates/pages/admin_index.qtpl:43
func (body *AdminPageBody) ShowAdminPage(activePage string) string {
//line views/templates/pages/admin_index.qtpl:43
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/pages/admin_index.qtpl:43
	body.WriteShowAdminPage(qb422016, activePage)
//line views/templates/pages/admin_index.qtpl:43
	qs422016 := string(qb422016.B)
//line views/templates/pages/admin_index.qtpl:43
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/pages/admin_index.qtpl:43
	return qs422016
//line views/templates/pages/admin_index.qtpl:43
}
