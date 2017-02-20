// This file is automatically generated by qtc from "head.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line views/templates/layouts/head.qtpl:1
package layouts

//line views/templates/layouts/head.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// All the text outside function templates is treated as comments,
// i.e. it is just ignored by quicktemplate compiler (`qtc`). It is for humans.
//
// head for index page.

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
    <meta name="author" content="">
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
    <link rel="stylesheet" href="http://solution.allservice.in.ua/js/fancybox2/jquery.fancybox.css" type="text/css" media="screen">
    <link rel="stylesheet" href="http://solution.allservice.in.ua/css/stog.css" type="text/css"  media="screen">
    <link rel="stylesheet" href="http://solution.allservice.in.ua/css/menu.css" type="text/css" media="screen">
    <link rel="stylesheet" type="text/css" href="http://solution.allservice.in.ua/js/fancybox2/helpers/jquery.fancybox-buttons.css" media="screen">
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
    <link rel="stylesheet" href="https://cdnjs.com/libraries/twitter-bootstrap/2.3.2">
    <link rel="stylesheet" href="/jquery.formstyler.css">
    <link rel="stylesheet" href="/fonts.css">
    <link rel="stylesheet" href="/travel.css"">
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
        <link rel="stylesheet" href="/jquery.datetimepicker.css" type="text/css" media="screen">
    `)
	//line views/templates/layouts/head.qtpl:51
	for i := 0; i < len(head.LinkStyles); i++ {
		//line views/templates/layouts/head.qtpl:51
		qw422016.N().S(`
        `)
		//line views/templates/layouts/head.qtpl:52
		qw422016.E().S(head.LinkStyles[i])
		//line views/templates/layouts/head.qtpl:52
		qw422016.N().S(`
    `)
		//line views/templates/layouts/head.qtpl:53
	}
	//line views/templates/layouts/head.qtpl:53
	qw422016.N().S(`
    <!--[if IE]><script type="text/javascript" src="https://code.jquery.com/jquery-1.12.4.min.js"></script><![endif]-->
    <!--[if !IE]--><script type="text/javascript" src="https://code.jquery.com/jquery-3.1.1.min.js"></script><!--[endif]-->
    <script  src="http://yandex.st/jquery/fancybox/2.1.4/jquery.fancybox.min.js"></script>
    <script  src="http://solution.allservice.in.ua/js/fancybox2/helpers/jquery.fancybox-buttons.js"></script>
    <script  src="http://yandex.st/jquery/form/3.14/jquery.form.min.js"></script>
    <script  src="http://solution.allservice.in.ua/js/stog.js"></script>
    <script src="/js/bootstrap.min.js"></script>
    <!--[if lt IE 10]>
    <script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/jquery-ajaxtransport-xdomainrequest/1.0.1/jquery.xdomainrequest.min.js"></script>
    <![endif]-->
    <script src="//cdn.rawgit.com/Mikhus/canvas-gauges/gh-pages/download/latest/all/gauge.min.js"></script>
    <!-- Include Radiant Media Player JavaScript file in your <body> or <head> -->
    <script src="https://cdn.radiantmediatechs.com/rmp/v3/latest/js/rmp.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/2.3.2/js/bootstrap.min.js"></script>
    <script  src="/js/system.js"></script>
    <script  src="/js/drag_drop.js"></script>
    <script  src="/js/forms.js"></script>
    <script  src="/js/jquery.formstyler.min.js"></script>
    <script  src="/js/forView.js"></script>
    `)
	//line views/templates/layouts/head.qtpl:73
	for i := 0; i < len(head.Scripts); i++ {
		//line views/templates/layouts/head.qtpl:73
		qw422016.N().S(`
        `)
		//line views/templates/layouts/head.qtpl:74
		qw422016.E().S(head.Scripts[i])
		//line views/templates/layouts/head.qtpl:74
		qw422016.N().S(`
    `)
		//line views/templates/layouts/head.qtpl:75
	}
	//line views/templates/layouts/head.qtpl:75
	qw422016.N().S(`
    <link href="https://cdn.jsdelivr.net/jquery.suggestions/16.8/css/suggestions.css" type="text/css" rel="stylesheet" />
    <script type="text/javascript" src="https://cdn.jsdelivr.net/jquery.suggestions/16.8/js/jquery.suggestions.min.js"></script>

</head>
`)
//line views/templates/layouts/head.qtpl:80
}

//line views/templates/layouts/head.qtpl:80
func (head *HeadHTMLPage) WriteHeadHTML(qq422016 qtio422016.Writer) {
	//line views/templates/layouts/head.qtpl:80
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/layouts/head.qtpl:80
	head.StreamHeadHTML(qw422016)
	//line views/templates/layouts/head.qtpl:80
	qt422016.ReleaseWriter(qw422016)
//line views/templates/layouts/head.qtpl:80
}

//line views/templates/layouts/head.qtpl:80
func (head *HeadHTMLPage) HeadHTML() string {
	//line views/templates/layouts/head.qtpl:80
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/layouts/head.qtpl:80
	head.WriteHeadHTML(qb422016)
	//line views/templates/layouts/head.qtpl:80
	qs422016 := string(qb422016.B)
	//line views/templates/layouts/head.qtpl:80
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/layouts/head.qtpl:80
	return qs422016
//line views/templates/layouts/head.qtpl:80
}
