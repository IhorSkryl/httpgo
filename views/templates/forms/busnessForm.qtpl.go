// This file is automatically generated by qtc from "busnessForm.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line views/templates/forms/busnessForm.qtpl:1
package forms

//line views/templates/forms/busnessForm.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// All the text outside function templates is treated as comments,
// i.e. it is just ignored by quicktemplate compiler (`qtc`). It is for humans.
//
// .

//line views/templates/forms/busnessForm.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/forms/busnessForm.qtpl:5
func StreamDadataScript(qw422016 *qt422016.Writer) {
	//line views/templates/forms/busnessForm.qtpl:5
	qw422016.N().S(`
<head>
    <link href="https://cdn.jsdelivr.net/jquery.suggestions/16.8/css/suggestions.css" type="text/css" rel="stylesheet" />
    <script type="text/javascript" src="https://cdn.jsdelivr.net/jquery.suggestions/16.8/js/jquery.suggestions.min.js"></script>
</head>
<script>
    $("#name").suggestions({
        serviceUrl: "https://suggestions.dadata.ru/suggestions/api/4_1/rs",
        token: "5bad4cdc891635fea16e676d89bea22bb88286a6",
        type: "PARTY",
        count: 5,
        /* Вызывается, когда пользователь выбирает одну из подсказок */
        onSelect: busnessSuggestion,
    });
</script>
`)
//line views/templates/forms/busnessForm.qtpl:20
}

//line views/templates/forms/busnessForm.qtpl:20
func WriteDadataScript(qq422016 qtio422016.Writer) {
	//line views/templates/forms/busnessForm.qtpl:20
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/forms/busnessForm.qtpl:20
	StreamDadataScript(qw422016)
	//line views/templates/forms/busnessForm.qtpl:20
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/busnessForm.qtpl:20
}

//line views/templates/forms/busnessForm.qtpl:20
func DadataScript() string {
	//line views/templates/forms/busnessForm.qtpl:20
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/forms/busnessForm.qtpl:20
	WriteDadataScript(qb422016)
	//line views/templates/forms/busnessForm.qtpl:20
	qs422016 := string(qb422016.B)
	//line views/templates/forms/busnessForm.qtpl:20
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/forms/busnessForm.qtpl:20
	return qs422016
//line views/templates/forms/busnessForm.qtpl:20
}
