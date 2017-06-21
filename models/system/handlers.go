package system

import (
	"net/http"
	//"log"
	"fmt"
	"github.com/ruslanBik4/httpgo/views"
	//"runtime"
	"github.com/ruslanBik4/httpgo/models/logs"
)

type ErrNotLogin struct {
	Message string
}
func (err ErrNotLogin) Error() string{
	return err.Message
}
//Структура для ошибок базы данных
type ErrDb struct {
	Message string
}
//Функция для обработк структуры ошибок базы данных
func (err ErrDb) Error() string{
	return err.Message
}

type ErrNotPermission struct {
	Message string
}
func (err ErrNotPermission) Error() string{
	return err.Message
}

func Catch(w http.ResponseWriter, r *http.Request) {
	result := recover()

	switch err := result.(type) {
	case ErrNotLogin:
		fmt.Fprintf(w, "<title>%s</title>", "Для начала работы необходимо авторизоваться!" )
		views.RenderSignForm(w, r, "")
	case ErrNotPermission:
		views.RenderNoPermissionPage(w)
	case nil:
	case error:
		views.RenderInternalError(w, err)
		logs.ErrorStack()
	}
}

func WrapCatchHandler(fnc http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer Catch(w,r)
		fnc(w,r)
	})
}

