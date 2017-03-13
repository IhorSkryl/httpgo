package views

//go:generate /Users/rus/go/bin/qtc -dir=views/templates

import (
	"github.com/ruslanBik4/httpgo/views/templates/forms"
	"fmt"
	"net/http"
	"github.com/ruslanBik4/httpgo/models/db"
	"github.com/ruslanBik4/httpgo/views/templates/layouts"
	"github.com/ruslanBik4/httpgo/views/templates/pages"
	"github.com/ruslanBik4/httpgo/views/templates/json"
	"github.com/ruslanBik4/httpgo/models/users"
)

//noinspection GoInvalidConstType
var HEADERS = map[string] string {
	"Content-Type": "text/html; charset=utf-8",
	"author":	"uStudio",
}
func WriteHeaders(w http.ResponseWriter) {
	// выдаем стандартные заголовки страницы
	for key, value := range HEADERS {
		w.Header().Set(key, value)
	}
}
func IsAJAXRequest(r *http.Request) bool {
	return len(r.Header["X-Requested-With"]) > 0
}
func RenderAnyPage(w http.ResponseWriter, r *http.Request, strContent string) {
	if IsAJAXRequest(r) {
		fmt.Fprint(w, strContent)
	} else {
		p := &pages.IndexPageBody{ Content: strContent }
		RenderTemplate(w, r, "index", p)
	}
}
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmplName string, Content interface{} ) error {
	var menu db.MenuItems

	WriteHeaders(w)

	headPage  := &layouts.HeadHTMLPage{
		Charset: "charset=utf-8",
		Language: "ru",
		Title: "Заголовок новой страницы",
	}

	switch tmplName {
	case "index":
		var p *pages.IndexPageBody = Content.(*pages.IndexPageBody)


		p.TopMenu = make( map[string] string, 0)

		menu.GetMenu("indexTop")

		for _, item := range menu.Items {
			p.TopMenu[item.Title] = "/menu/" + item.Name + "/"

		}
		if p.Content == "" {

			//p.Title   = "Авторизация"
			if userID, ok := users.IsLogin(r); ok {
				p.Content = fmt.Sprintf("<script>afterLogin({login:'%d',sex:'0'})</script>", userID)
			} else {
				p.Content = forms.SigninForm("", "Введите пароль") + forms.ShowForm("введите фамилию имя отчество")
			}
			headPage.Title = "введите Ваше имя отчество фамилию"
			p.Route = "/"
		}

		fmt.Fprint(w, headPage.HeadHTML())
		fmt.Fprint(w, p.IndexHTML())
	case "signinForm":
		var p *pages.IndexPageBody = Content.(*pages.IndexPageBody)
		RenderAnyPage(w, r, forms.SigninForm(p.Title, "Введите пароль, полученный по почте") )

	case "adminPage":
		var p *pages.AdminPageBody = Content.(*pages.AdminPageBody)

		p.TopMenu = make( map[string] string, 0)
		menu.GetMenu("indexTop")

		for _, item := range menu.Items {
			p.TopMenu[item.Title] = "/menu/" + item.Name + "/"

		}
		fmt.Fprint(w, headPage.HeadHTML())
		fmt.Fprint(w, p.ShowAdminPage(""))
	default:
		fmt.Fprint(w, "no rendering with page %s with data %v", tmplName, Content)
	}
	return nil
}

func RenderAnyForm(w http.ResponseWriter, r *http.Request, Title string, fields forms.FieldsTable,
			Inputs map[string] []string, head, foot string ) error  {

	WriteHeaders(w)

	if Inputs != nil {
		head += layouts.DadataHead()
		foot += layouts.DadataScript(Inputs)
	}
	RenderAnyPage(w, r, head + layouts.PutHeadForm() + fields.ShowAnyForm("/admin/exec/", Title) + layouts.PutEndForm() + foot )

	return nil

}
func RenderAnyJSON(w http.ResponseWriter, arrJSON map[string] interface {}) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, json.WriteAnyJSON(arrJSON) )
}
