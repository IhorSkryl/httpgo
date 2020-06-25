// Code generated by qtc from "footer.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/templates/layouts/footer.qtpl:1
package layouts

//line views/templates/layouts/footer.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/templates/layouts/footer.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/layouts/footer.qtpl:1
func StreamFooterHTML(qw422016 *qt422016.Writer, email, phone1, phone2 string) {
//line views/templates/layouts/footer.qtpl:1
	qw422016.N().S(`
<footer class="main-footer">
		<div class="footer-inner-wrap">
            <div class="footer-block">
                <h5>Защищенные системы</h5>
                <img src="/images/bitmap.png" alt="" class="footer-img">
                <img src="/images/bitmap1.png" alt="" class="footer-img">
            </div>
            <div class="footer-block">
            	<div class="footer-phones">
            		<div class="first-phone clearfix">
            			<span class="phone-numbers">`)
//line views/templates/layouts/footer.qtpl:12
	qw422016.E().S(phone1)
//line views/templates/layouts/footer.qtpl:12
	qw422016.N().S(` &ensp;- &ensp;free on country</span>
            		</div>
            		<div class="second-phone clearfix">
            			<span class="phone-numbers2">`)
//line views/templates/layouts/footer.qtpl:15
	qw422016.E().S(phone2)
//line views/templates/layouts/footer.qtpl:15
	qw422016.N().S(`</span>
            			<span class="phone-descr">&ensp;&ensp; - &ensp;согласно тарифам</span>
            			<div class="perenos">Вашего оператора</div>
            		</div>
            		<div class="mail">
            			<a href="mailto:ingo@82RU.ru" class="footer-mail">`)
//line views/templates/layouts/footer.qtpl:20
	qw422016.E().S(email)
//line views/templates/layouts/footer.qtpl:20
	qw422016.N().S(`</a>
            		</div>
            	</div>
            	<div class="footer-social">
            		<img src="/images/instagram.svg" alt="" class="social">
            		<img src="/images/facebook.svg" alt="" class="social">
            		<img src="/images/vk.svg" alt="" class="social">
            	</div>
            </div>
            <div class="footer-block">
            	<div class="footer-mnu">
            		<ul class="main-list">
            			<li><a href="#">Главная</a></li>
            			<li><a href="#">Отели</a></li>
            			<li><a href="#">Транспорт</a></li>
            			<li><a href="#">Отдых</a></li>
            			<li><a href="#">Путеводитель</a></li>
            			<li><a href="#">Блог</a></li>
            			<li><a href="#">Партнерам</a></li>
            		</ul>
            		<ul class="second-list">
            				<li><a href="#">F.A.Q.</a></li>
            				<li><a href="#">О нас</a></li>
            			</ul>
            	</div>
            </div>
       </div>
	</footer>
`)
//line views/templates/layouts/footer.qtpl:48
}

//line views/templates/layouts/footer.qtpl:48
func WriteFooterHTML(qq422016 qtio422016.Writer, email, phone1, phone2 string) {
//line views/templates/layouts/footer.qtpl:48
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/layouts/footer.qtpl:48
	StreamFooterHTML(qw422016, email, phone1, phone2)
//line views/templates/layouts/footer.qtpl:48
	qt422016.ReleaseWriter(qw422016)
//line views/templates/layouts/footer.qtpl:48
}

//line views/templates/layouts/footer.qtpl:48
func FooterHTML(email, phone1, phone2 string) string {
//line views/templates/layouts/footer.qtpl:48
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/layouts/footer.qtpl:48
	WriteFooterHTML(qb422016, email, phone1, phone2)
//line views/templates/layouts/footer.qtpl:48
	qs422016 := string(qb422016.B)
//line views/templates/layouts/footer.qtpl:48
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/layouts/footer.qtpl:48
	return qs422016
//line views/templates/layouts/footer.qtpl:48
}
