// This file is automatically generated by qtc from "anothersignupForm.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line anothersignupForm.qtpl:1
package forms

//line anothersignupForm.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// <!DOCTYPE html>

//line anothersignupForm.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line anothersignupForm.qtpl:2
func StreamAnotherSignUpForm(qw422016 *qt422016.Writer, placeholder string) {
	//line anothersignupForm.qtpl:2
	qw422016.N().S(`
<head>
    <link rel="import" href="/components/paper-input/paper-input-container.html">
    <link rel="import" href="/components/paper-input/paper-input-error.html">
    <link rel="import" href="/components/paper-input/paper-input.html">
    <link href="https://cdn.jsdelivr.net/jquery.suggestions/16.8/css/suggestions.css" type="text/css" rel="stylesheet" />
    <script type="text/javascript" src="https://cdn.jsdelivr.net/jquery.suggestions/16.8/js/jquery.suggestions.min.js"></script>
</head>
<form target="content" action="/admin/anothersignup/" method="post"  enctype="multipart/form-data"
      onsubmit="return saveForm(this, afterSignup);"  class="form-signin">
    <h2 class="form-signin-heading">Регистрация</h2>
    <select id="sex" name="sex" style="display: none;">
        <option value="0">Господин</option>
        <option value="1">Госпожа</option>
    </select>
    <label for="fullname">Введите Ваше ФИО</label>
    <input id="fullname" name="fullname" class="input-block-level" type="text" size="100" required placeholder="`)
	//line anothersignupForm.qtpl:18
	qw422016.E().S(placeholder)
	//line anothersignupForm.qtpl:18
	qw422016.N().S(`"/>
    <script type="text/javascript">
        $("#fullname").suggestions({
            serviceUrl: "https://suggestions.dadata.ru/suggestions/api/4_1/rs",
            token: "5bad4cdc891635fea16e676d89bea22bb88286a6",
            type: "NAME",
            count: 5,
            /* Вызывается, когда пользователь выбирает одну из подсказок */
            onSelect: signSuggestion,
        });
    </script>
    <label>E-mail, по которому мы зарегистрируем Вас в системе</label>

    <paper-input label="text input"></paper-input>
    <input type="email" name="login" class="input-block-level" required placeholder="email для регистрации в системе">
    <input type="submit" value="Зарегистрироваться.">
    <img class='loading' src='http://solution.allservice.in.ua/loading.gif' style='display:none;'>
    <progress value='0' max='100' hidden > </progress>
</form>

`)
//line anothersignupForm.qtpl:38
}

//line anothersignupForm.qtpl:38
func WriteAnotherSignUpForm(qq422016 qtio422016.Writer, placeholder string) {
	//line anothersignupForm.qtpl:38
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line anothersignupForm.qtpl:38
	StreamAnotherSignUpForm(qw422016, placeholder)
	//line anothersignupForm.qtpl:38
	qt422016.ReleaseWriter(qw422016)
//line anothersignupForm.qtpl:38
}

//line anothersignupForm.qtpl:38
func AnotherSignUpForm(placeholder string) string {
	//line anothersignupForm.qtpl:38
	qb422016 := qt422016.AcquireByteBuffer()
	//line anothersignupForm.qtpl:38
	WriteAnotherSignUpForm(qb422016, placeholder)
	//line anothersignupForm.qtpl:38
	qs422016 := string(qb422016.B)
	//line anothersignupForm.qtpl:38
	qt422016.ReleaseByteBuffer(qb422016)
	//line anothersignupForm.qtpl:38
	return qs422016
//line anothersignupForm.qtpl:38
}

//line anothersignupForm.qtpl:40
type AnotherMarshalRow struct {
	Msg string
	N   int
}

type AnotherPersonData struct {
	Id    int
	Login string
	Email string
	Sex   int
	Rows  []MarshalRow
}

// JSON marshaling

//line anothersignupForm.qtpl:56
func (d *AnotherPersonData) StreamJSON(qw422016 *qt422016.Writer) {
	//line anothersignupForm.qtpl:56
	qw422016.N().S(`{"login": "`)
	//line anothersignupForm.qtpl:58
	qw422016.E().S(d.Login)
	//line anothersignupForm.qtpl:58
	qw422016.N().S(`","email": "`)
	//line anothersignupForm.qtpl:59
	qw422016.E().S(d.Email)
	//line anothersignupForm.qtpl:59
	qw422016.N().S(`","sex": "`)
	//line anothersignupForm.qtpl:60
	if d.Sex == 0 {
		//line anothersignupForm.qtpl:60
		qw422016.N().S(`господин`)
		//line anothersignupForm.qtpl:60
	} else {
		//line anothersignupForm.qtpl:60
		qw422016.N().S(`госпожа`)
		//line anothersignupForm.qtpl:60
	}
	//line anothersignupForm.qtpl:60
	qw422016.N().S(`","Rows":[`)
	//line anothersignupForm.qtpl:62
	for i, r := range d.Rows {
		//line anothersignupForm.qtpl:62
		qw422016.N().S(`{"Msg":`)
		//line anothersignupForm.qtpl:64
		qw422016.N().Q(r.Msg)
		//line anothersignupForm.qtpl:64
		qw422016.N().S(`,"N":`)
		//line anothersignupForm.qtpl:65
		qw422016.N().D(r.N)
		//line anothersignupForm.qtpl:65
		qw422016.N().S(`}`)
		//line anothersignupForm.qtpl:67
		if i+1 < len(d.Rows) {
			//line anothersignupForm.qtpl:67
			qw422016.N().S(`,`)
			//line anothersignupForm.qtpl:67
		}
		//line anothersignupForm.qtpl:68
	}
	//line anothersignupForm.qtpl:68
	qw422016.N().S(`]}`)
//line anothersignupForm.qtpl:71
}

//line anothersignupForm.qtpl:71
func (d *AnotherPersonData) WriteJSON(qq422016 qtio422016.Writer) {
	//line anothersignupForm.qtpl:71
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line anothersignupForm.qtpl:71
	d.StreamJSON(qw422016)
	//line anothersignupForm.qtpl:71
	qt422016.ReleaseWriter(qw422016)
//line anothersignupForm.qtpl:71
}

//line anothersignupForm.qtpl:71
func (d *AnotherPersonData) JSON() string {
	//line anothersignupForm.qtpl:71
	qb422016 := qt422016.AcquireByteBuffer()
	//line anothersignupForm.qtpl:71
	d.WriteJSON(qb422016)
	//line anothersignupForm.qtpl:71
	qs422016 := string(qb422016.B)
	//line anothersignupForm.qtpl:71
	qt422016.ReleaseByteBuffer(qb422016)
	//line anothersignupForm.qtpl:71
	return qs422016
//line anothersignupForm.qtpl:71
}
