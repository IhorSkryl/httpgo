// This file is automatically generated by qtc from "signinForm.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line views/templates/forms/signinForm.qtpl:1
package forms

//line views/templates/forms/signinForm.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// All the text outside function templates is treated as comments,
// i.e. it is just ignored by quicktemplate compiler (`qtc`). It is for humans.
//
// форма авторизации на сайте.

//line views/templates/forms/signinForm.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/forms/signinForm.qtpl:6
type SignForm struct {
	Email, Password string
}

//line views/templates/forms/signinForm.qtpl:10
func (sf *SignForm) StreamSigninForm(qw422016 *qt422016.Writer) {
	//line views/templates/forms/signinForm.qtpl:10
	qw422016.N().S(`
<div class="main-form-wrap">
    <form target="content" action="/user/signin/" class="form-signin" onsubmit="return saveForm(this, afterLogin);" >
        <h2 class="form-signin-heading">Авторизация</h2>
        <input type="email" name="login" class="input-block-level" placeholder="Email, указанный при регистрации" value="`)
	//line views/templates/forms/signinForm.qtpl:14
	qw422016.E().S(sf.Email)
	//line views/templates/forms/signinForm.qtpl:14
	qw422016.N().S(`">
        <input type="password" name="password" class="input-block-level" placeholder="`)
	//line views/templates/forms/signinForm.qtpl:15
	qw422016.E().S(sf.Password)
	//line views/templates/forms/signinForm.qtpl:15
	qw422016.N().S(`">
        <label class="checkbox">
         <input type="checkbox" name="remember" value="remember-me"> Запомнить меня в системе
        </label>
       <button class="main-btn" type="submit">Войти</button>
    </form>
</div>
`)
//line views/templates/forms/signinForm.qtpl:22
}

//line views/templates/forms/signinForm.qtpl:22
func (sf *SignForm) WriteSigninForm(qq422016 qtio422016.Writer) {
	//line views/templates/forms/signinForm.qtpl:22
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/forms/signinForm.qtpl:22
	sf.StreamSigninForm(qw422016)
	//line views/templates/forms/signinForm.qtpl:22
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/signinForm.qtpl:22
}

//line views/templates/forms/signinForm.qtpl:22
func (sf *SignForm) SigninForm() string {
	//line views/templates/forms/signinForm.qtpl:22
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/forms/signinForm.qtpl:22
	sf.WriteSigninForm(qb422016)
	//line views/templates/forms/signinForm.qtpl:22
	qs422016 := string(qb422016.B)
	//line views/templates/forms/signinForm.qtpl:22
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/forms/signinForm.qtpl:22
	return qs422016
//line views/templates/forms/signinForm.qtpl:22
}
