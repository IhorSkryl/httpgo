// Code generated by qtc from "head.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// All the text outside function templates is treated as comments,
// i.e. it is just ignored by quicktemplate compiler (`qtc`). It is for humans.
//
// head for index page.

//line views/templates/layouts/head.qtpl:5
package layouts

//line views/templates/layouts/head.qtpl:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/templates/layouts/head.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/layouts/head.qtpl:6
type HeadHTMLPage struct {
	Title      string
	Language   string
	Charset    string
	LinkStyles []string
	MetaTags   []string
	Scripts    []string
}

//line views/templates/layouts/head.qtpl:16
func (head *HeadHTMLPage) StreamHeadHTML(qw422016 *qt422016.Writer) {
//line views/templates/layouts/head.qtpl:16
	qw422016.N().S(`
<!DOCTYPE html>
<html lang="`)
//line views/templates/layouts/head.qtpl:18
	qw422016.E().S(head.Language)
//line views/templates/layouts/head.qtpl:18
	qw422016.N().S(`">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta artel='noSC'>
    <meta charset="`)
//line views/templates/layouts/head.qtpl:22
	qw422016.E().S(head.Charset)
//line views/templates/layouts/head.qtpl:22
	qw422016.N().S(`">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="">
    <meta name="author" content="httpgo">
    `)
//line views/templates/layouts/head.qtpl:26
	for i := 0; i < len(head.MetaTags); i++ {
//line views/templates/layouts/head.qtpl:26
		qw422016.N().S(`
        `)
//line views/templates/layouts/head.qtpl:27
		qw422016.E().S(head.MetaTags[i])
//line views/templates/layouts/head.qtpl:27
		qw422016.N().S(`
    `)
//line views/templates/layouts/head.qtpl:28
	}
//line views/templates/layouts/head.qtpl:28
	qw422016.N().S(`
    <title>`)
//line views/templates/layouts/head.qtpl:29
	qw422016.E().S(head.Title)
//line views/templates/layouts/head.qtpl:29
	qw422016.N().S(`</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/fancybox@3.0.1/dist/css/jquery.fancybox.css" type="text/css" media="screen">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/fancybox/3.5.7/jquery.fancybox.min.js" type="text/css" media="screen">
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/2.3.2/css/bootstrap-responsive.css">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/jScrollPane/2.0.23/style/jquery.jscrollpane.min.css">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/jquery.formstyler/1.7.8/jquery.formstyler.css">

    <!-- HTML5 shim, for IE6-8 support of HTML5 elements -->
    <!--[if lt IE 9]>
    <script src="http://html5shim.googlecode.com/svn/trunk/html5.js"></script>
    <![endif]-->

    <!-- Fav and touch icons -->
    <link rel="apple-touch-icon-precomposed" sizes="144x144" href="/bootstrap/ico/apple-touch-icon-144-precomposed.png">
    <link rel="apple-touch-icon-precomposed" sizes="114x114" href="/bootstrap/ico/apple-touch-icon-114-precomposed.png">
    <link rel="apple-touch-icon-precomposed" sizes="72x72" href="/bootstrap/ico/apple-touch-icon-72-precomposed.png">
    <link rel="apple-touch-icon-precomposed" href="/bootstrap/ico/apple-touch-icon-57-precomposed.png">
    <link rel="shortcut icon" href="/bootstrap/ico/favicon.png">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/jquery-datetimepicker/2.5.4/build/jquery.datetimepicker.min.css" type="text/css" media="screen">
    `)
//line views/templates/layouts/head.qtpl:49
	for i := 0; i < len(head.LinkStyles); i++ {
//line views/templates/layouts/head.qtpl:49
		qw422016.N().S(`
        `)
//line views/templates/layouts/head.qtpl:50
		qw422016.E().S(head.LinkStyles[i])
//line views/templates/layouts/head.qtpl:50
		qw422016.N().S(`
    `)
//line views/templates/layouts/head.qtpl:51
	}
//line views/templates/layouts/head.qtpl:51
	qw422016.N().S(`
    <script  src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.5.1/jquery.min.js" defer></script>
    <script  src="https://cdnjs.cloudflare.com/ajax/libs/jquery.form/4.3.0/jquery.form.min.js" defer></script>
    <script  src="/js/fancybox2/helpers/jquery.fancybox-buttons.js" defer></script>
    <script  src="https://cdnjs.cloudflare.com/ajax/libs/fancybox/3.5.7/jquery.fancybox.min.js" defer></script>
    <script  src="https://cdnjs.cloudflare.com/ajax/libs/jquery.nicescroll/3.6.8-fix/jquery.nicescroll.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/5.0.0-alpha1/js/bootstrap.min.js"></script>
    <!--[if lt IE 10]>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery-ajaxtransport-xdomainrequest/1.0.1/jquery.xdomainrequest.min.js"></script>
    <![endif]-->
    <script src="https://cdn.rawgit.com/Mikhus/canvas-gauges/gh-pages/download/latest/all/gauge.min.js"></script>
    <!-- Include Radiant Media Player JavaScript file in your <body> or <head> -->
    <script src="https://cdn.radiantmediatechs.com/rmp/v3/latest/js/rmp.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery-datetimepicker/2.5.4/build/jquery.datetimepicker.full.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery-mousewheel/3.1.13/jquery.mousewheel.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jScrollPane/2.0.23/script/jquery.jscrollpane.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/2.3.2/js/bootstrap.min.js"></script>
    <script  src="https://cdn.jsdelivr.net/jquery.formstyler/1.7.8/jquery.formstyler.min.js"></script>
    `)
//line views/templates/layouts/head.qtpl:69
	for i := 0; i < len(head.Scripts); i++ {
//line views/templates/layouts/head.qtpl:69
		qw422016.N().S(`
        `)
//line views/templates/layouts/head.qtpl:70
		qw422016.E().S(head.Scripts[i])
//line views/templates/layouts/head.qtpl:70
		qw422016.N().S(`
    `)
//line views/templates/layouts/head.qtpl:71
	}
//line views/templates/layouts/head.qtpl:71
	qw422016.N().S(`
    <link href="https://cdn.jsdelivr.net/jquery.suggestions/16.8/css/suggestions.css" type="text/css" rel="stylesheet" />
    <script src="https://cdn.jsdelivr.net/jquery.suggestions/16.8/js/jquery.suggestions.min.js" defer></script>

</head>

`)
//line views/templates/layouts/head.qtpl:77
}

//line views/templates/layouts/head.qtpl:77
func (head *HeadHTMLPage) WriteHeadHTML(qq422016 qtio422016.Writer) {
//line views/templates/layouts/head.qtpl:77
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/layouts/head.qtpl:77
	head.StreamHeadHTML(qw422016)
//line views/templates/layouts/head.qtpl:77
	qt422016.ReleaseWriter(qw422016)
//line views/templates/layouts/head.qtpl:77
}

//line views/templates/layouts/head.qtpl:77
func (head *HeadHTMLPage) HeadHTML() string {
//line views/templates/layouts/head.qtpl:77
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/layouts/head.qtpl:77
	head.WriteHeadHTML(qb422016)
//line views/templates/layouts/head.qtpl:77
	qs422016 := string(qb422016.B)
//line views/templates/layouts/head.qtpl:77
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/layouts/head.qtpl:77
	return qs422016
//line views/templates/layouts/head.qtpl:77
}

//line views/templates/layouts/head.qtpl:82
func StreamAdminHead(qw422016 *qt422016.Writer, title string) {
//line views/templates/layouts/head.qtpl:82
	qw422016.N().S(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="author" content="httpgo">
    <link rel="stylesheet" type="text/css" href="/bootstrap.min.css">
    <link rel="stylesheet" type="text/css" href="/bootstrap-theme.min.css">
    <link rel="stylesheet" type="text/css" href="/jquery.formstyler.css">
    <link href="https://fonts.googleapis.com/css?family=Fira+Sans" rel="stylesheet">
    <link href="https://fonts.googleapis.com/css?family=Open+Sans:300,400" rel="stylesheet">
    <link rel="stylesheet" type="text/css" href="/main.css">
    <title>`)
//line views/templates/layouts/head.qtpl:96
	qw422016.E().S(title)
//line views/templates/layouts/head.qtpl:96
	qw422016.N().S(`</title>
</head>

`)
//line views/templates/layouts/head.qtpl:99
}

//line views/templates/layouts/head.qtpl:99
func WriteAdminHead(qq422016 qtio422016.Writer, title string) {
//line views/templates/layouts/head.qtpl:99
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/layouts/head.qtpl:99
	StreamAdminHead(qw422016, title)
//line views/templates/layouts/head.qtpl:99
	qt422016.ReleaseWriter(qw422016)
//line views/templates/layouts/head.qtpl:99
}

//line views/templates/layouts/head.qtpl:99
func AdminHead(title string) string {
//line views/templates/layouts/head.qtpl:99
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/layouts/head.qtpl:99
	WriteAdminHead(qb422016, title)
//line views/templates/layouts/head.qtpl:99
	qs422016 := string(qb422016.B)
//line views/templates/layouts/head.qtpl:99
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/layouts/head.qtpl:99
	return qs422016
//line views/templates/layouts/head.qtpl:99
}
