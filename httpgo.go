package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/ruslanBik4/httpgo/models/admin"
	"github.com/ruslanBik4/httpgo/models/api/v1"
	"github.com/ruslanBik4/httpgo/models/db"
	"github.com/ruslanBik4/httpgo/models/db/qb"
	_ "github.com/ruslanBik4/httpgo/models/docs"
	"github.com/ruslanBik4/httpgo/models/logs"
	"github.com/ruslanBik4/httpgo/models/server"
	"github.com/ruslanBik4/httpgo/models/services"
	"github.com/ruslanBik4/httpgo/models/system"
	"github.com/ruslanBik4/httpgo/models/users"
	"github.com/ruslanBik4/httpgo/views"
	"github.com/ruslanBik4/httpgo/views/fonts"
	"github.com/ruslanBik4/httpgo/views/templates/layouts"
	"github.com/ruslanBik4/httpgo/views/templates/pages"
	"io/ioutil"
	"models/config"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
	"github.com/ruslanBik4/httpgo/views/templates/json"
	"os/signal"
)

//go:generate qtc -dir=views/templates

const fpmSocket = "/var/run/php5-fpm.sock"

var (
	headerNameReplacer = strings.NewReplacer(" ", "_", "-", "_")
	// ErrIndexMissingSplit describes an index configuration error.
	//ErrIndexMissingSplit = errors.New("configured index file(s) must include split value")

	cacheMu sync.RWMutex
	cache   = map[string][]byte{}
	routes  = map[string]http.HandlerFunc{
		"/recache":       handlerRecache,
		"/update/":       handleUpdate,
		"/test/":         handleTest,
		"/fonts/":        fonts.HandleGetFont,
		"/query/":        db.HandlerDBQuery,
		"/menu/":         handlerMenu,
		"/show/forms/":   handlerForms,
		"/user/signup/":  users.HandlerSignUp,
		"/user/signin/":  users.HandlerSignIn,
		"/user/signout/": users.HandlerSignOut,
		"/user/active/":  users.HandlerActivateUser,
		"/user/profile/": users.HandlerProfile,
		//"/user/oauth/":    users.HandlerQauth2,
		"/user/GoogleCallback/": users.HandleGoogleCallback,
		"/components/":          handlerComponents,

		// TODO: API remove in single module
		"/api/v1/table/form/":   api.HandleFieldsJSON,
		"/api/v1/table/view/":   api.HandleTextRowJSON,
		"/api/v1/table/row/":    api.HandleRowJSON,
		"/api/v1/table/rows/":   api.HandleAllRowsJSON,
		"/api/v1/table/schema/": api.HandleSchema,
		"/api/v1/update/":       api.HandleUpdateServer,
		"/api/v1/restart/":      api.HandleRestartServer,
		"/api/v1/log/":          api.HandleLogServer,
		"/api/v1/photos/":       api.HandlePhotos,
		"/api/v1/video/":        api.HandleVideos,
		"/api/v1/photos/add/":   api.HandleAddPhoto,
		// short route
		"/api/table/form/":   api.HandleFieldsJSON,
		"/api/table/view/":   api.HandleTextRowJSON,
		"/api/table/row/":    api.HandleRowJSON,
		"/api/table/rows/":   api.HandleAllRowsJSON,
		"/api/table/schema/": api.HandleSchema,
		"/api/update/":       api.HandleUpdateServer,
		"/api/restart/":      api.HandleRestartServer,
		"/api/log/":          api.HandleLogServer,
		"/api/photos/":       api.HandlePhotos,
		"/api/video/":        api.HandleVideos,
		"/api/photos/add/":   api.HandleAddPhoto,
	}
)

func registerRoutes() {
	http.Handle("/", NewDefaultHandler())
	for route, fnc := range routes {
		http.HandleFunc(route, system.WrapCatchHandler(fnc))
	}
	admin.RegisterRoutes()
	config.RegisterRoutes()
}

// работа по умолчанию - кеширования общих файлов в частности, обработчики для php-fpm & php
type DefaultHandler struct {
	fpm       *system.FCGI
	php       *system.FCGI
	cache     []string
	whitelist []string
}

func NewDefaultHandler() *DefaultHandler {
	handler := &DefaultHandler{
		fpm: system.NewFPM(fpmSocket),
		php: system.NewPHP(*f_web, fpmSocket),
		cache: []string{
			".svg", ".css", ".js", ".map", ".ico",
		},
		whitelist: []string{
			".jpg", ".jpeg", ".png", ".gif", ".ttf", ".pdf", ".json", ".htm", ".html", ".txt", ".mp4",
		},
	}
	// read from flags
	cacheExt := *f_cache
	p := strings.Index(cacheExt, ";")
	for p > 0 {

		handler.cache = append(handler.cache, cacheExt[:p])
		cacheExt = cacheExt[p:]
		p = strings.Index(cacheExt, ";")
	}
	return handler
}
func (h *DefaultHandler) toCache(ext string) bool {
	for _, name := range h.cache {
		if ext == name {
			return true
		}
	}
	return false
}
func (h *DefaultHandler) toServe(ext string) bool {
	for _, name := range h.whitelist {
		if ext == name {
			return true
		}
	}
	return false
}
func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	defer system.Catch(w, r)

	switch r.URL.Path {
	case "/":
		http.Redirect(w, r, "/customer/", http.StatusTemporaryRedirect)
		return

		//p := &pages.IndexPageBody{Title: "Главная страница"}
		////для авторизованного пользователя - сразу показать его данные на странице
		//p.Content = fmt.Sprintf("<script>afterLogin({login:'%d',sex:'0'})</script>", userID)
		//var menu db.MenuItems
		//menu.GetMenu("indexTop")
		//
		//p.TopMenu = make(map[string]string, len(menu.Items))
		//for _, item := range menu.Items {
		//	p.TopMenu[item.Title] = "/menu/" + item.Name + "/"
		//
		//}
		//views.RenderTemplate(w, r, "index", p)
		// спецвойска
	case "/polymer.html":
		http.ServeFile(w, r, filepath.Join(*f_static, "views/components/polymer/polymer.html"))
	case "/polymer-mini.html":
		http.ServeFile(w, r, filepath.Join(*f_static, "views/components/polymer/polymer-mini.html"))
	case "/polymer-micro.html":
		http.ServeFile(w, r, filepath.Join(*f_static, "views/components/polymer/polymer-micro.html"))
	case "/status", "/ping", "/pong":
		h.fpm.ServeHTTP(w, r)
	default:
		filename := strings.TrimLeft(r.URL.Path, "/")
		ext := filepath.Ext(filename)

		if strings.HasPrefix(ext, ".php") {
			h.php.ServeHTTP(w, r)
		} else if h.toCache(ext) {
			serveAndCache(filename, w, r)
		} else if h.toServe(ext) {
			http.ServeFile(w, r, filepath.Join(*f_web, filename))
		} else if fileName := filepath.Join(*f_web, filename); ext == "" {
			http.ServeFile(w, r, fileName)
		} else {
			h.php.ServeHTTP(w, r)
		}
	}
}
func handlerComponents(w http.ResponseWriter, r *http.Request) {

	filename := strings.TrimLeft(r.URL.Path, "/")

	http.ServeFile(w, r, filepath.Join(*f_static+"/views", filename))

}

// считываем файлы типа css/js ect в память и потом отдаем из нее
func setCache(path string, data []byte) {
	cacheMu.Lock()
	cache[path] = data
	cacheMu.Unlock()
}
func getCache(path string) ([]byte, bool) {
	cacheMu.RLock()
	data, ok := cache[path]
	cacheMu.RUnlock()
	return data, ok
}
func emptyCache() {
	cacheMu.RLock()
	cache = make(map[string][]byte, 0)
	cacheMu.RUnlock()

}
func serveAndCache(filename string, w http.ResponseWriter, r *http.Request) {
	keyName := path.Base(filename)

	data, ok := getCache(keyName)
	if !ok {
		data, err := ioutil.ReadFile(filepath.Join(*f_static, filename))
		if os.IsNotExist(err) {
			data, err = ioutil.ReadFile(filepath.Join(*f_web, filename))
		}
		if system.WriteError(w, err) {
			return
		}
		setCache(keyName, data)
	}
	http.ServeContent(w, r, filename, time.Time{}, bytes.NewReader(data))
}

func sockCatch() {
	err := recover()
	logs.ErrorLog(err.(error))
}
const _24K = (1 << 10) * 24

func handleTest(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(_24K)
	id := r.FormValue("id")
	qBuilder := qb.Create("b.id=?", "", "")

	qBuilder.AddTable("b", "business")

	qBuilder.JoinTable("cl", "currency_list", "INNER JOIN", " ON b.id_currency_list=cl.id").AddFields(map[string]string{
		"currency_title": "title",
	})
	qBuilder.JoinTable("c", "country_list", "INNER JOIN", " ON b.id_country_list=c.id").AddFields(map[string]string{
		"country_bank_title": "title",
	})
	qBuilder.AddArg(id)

	if r.FormValue("batch") == "1" {
		arrJSON, err := qBuilder.SelectToMultidimension()
		if err != nil {
			views.RenderInternalError(w, err)
			return
		}
		views.RenderArrayJSON(w, arrJSON)

		return
	}

	w.Write([]byte("{"))

	err := qBuilder.SelectRunFunc(func(fields []*qb.QBField) error {
		for idx, field := range fields {
			if idx > 0 {
				w.Write( []byte (",") )
			}

			w.Write([]byte(`"` + fields[idx].Alias + `":`))
			if field.Value == nil {
				w.Write([]byte("null"))
			} else {
				json.WriteElement(w, field.GetNativeValue(true) )
			}
		}

		return nil
	})
	if err != nil {
		views.RenderInternalError(w, err)
		return
	}
	views.WriteJSONHeaders(w)
	w.Write([]byte("}"))
	return

	qBuilder = qb.Create("hs.id_hotels=?", "", "")

	qBuilder.AddTable("hs", "hotels_services").AddField("", "id_services_list AS id_services_list").AddField("", "id_hotels")
	qBuilder.Join("sl", "services_list", "ON (sl.id = hs.id_services_list)").AddField("", "id_services_category_list")

	//qBuilder.Union("SELECT sl.id AS id_services_list,  0 AS id_hotels, sl.id_services_category_list FROM services_list AS sl")

	qBuilder.AddArg(r.FormValue("id_hotels"))
	arrJSON, err := qBuilder.SelectToMultidimension()

	if err != nil {
		logs.ErrorLog(err)
		return
	}

	views.RenderArrayJSON(w, arrJSON)
	return
	//qBuilder := qb.Create("", "", "")

	logs.DebugLog(r)
	r.ParseMultipartForm(_24K)
	for _, headers := range r.MultipartForm.File {
		for _, header := range headers {
			var err interface{}
			inFile, _ := header.Open()

			err = services.Send("photos", "save", header.Filename, inFile)
			if err != nil {
				switch err.(type) {
				case services.ErrServiceNotCorrectOperation:
					logs.ErrorLog(err.(error))
				}
				w.Write([]byte(err.(error).Error()))

			} else {
				w.Write([]byte("Succesfull"))
			}
		}
	}

}

func handleUpdate(w http.ResponseWriter, r *http.Request) {

}

func handlerForms(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, r, r.FormValue("name")+"Form", &pages.IndexPageBody{Title: r.FormValue("email")})
}
func isAJAXRequest(r *http.Request) bool {
	return len(r.Header["X-Requested-With"]) > 0
}
func handlerMenu(w http.ResponseWriter, r *http.Request) {

	userID := users.IsLogin(r)
	resultId, err := strconv.Atoi(userID)
	if err != nil || !admin.GetUserPermissionForPageByUserId(resultId, r.URL.Path, "View") {
		views.RenderNoPermissionPage(w)
		return
	}
	var menu db.MenuItems

	idx := strings.LastIndex(r.URL.Path, "menu/") + 5
	idMenu := r.URL.Path[idx : len(r.URL.Path)-1]

	//отдаем полный список меню для фронтового фреймворка
	if idMenu == "all" {
		qBuilder := qb.CreateEmpty()
		qBuilder.AddTable("", "menu_items")

		if arrJSON, err := qBuilder.SelectToMultidimension(); err != nil {
			logs.ErrorLog(err.(error))
		} else {
			views.RenderArrayJSON(w, arrJSON)
		}
		return
	}

	var catalog, content string
	// отрисовка меню страницы
	if menu.GetMenu(idMenu) > 0 {

		p := &layouts.MenuOwnerBody{Title: idMenu, TopMenu: make(map[string]*layouts.ItemMenu, 0)}

		for _, item := range menu.Items {
			p.TopMenu[item.Title] = &layouts.ItemMenu{Link: "/menu/" + item.Name + "/"}

		}

		// return into parent menu if he occurent
		if menu.Self.ParentID > 0 {
			p.TopMenu["< на уровень выше"] = &layouts.ItemMenu{Link: fmt.Sprintf("/menu/%d/", menu.Self.ParentID)}
		}
		catalog = p.MenuOwner()
	}
	//для отрисовки контента страницы
	if menu.Self.Link > "" {
		content = fmt.Sprintf("<div class='autoload' data-href='%s'></div>", menu.Self.Link)
	}
	views.RenderAnyPage(w, r, catalog+content)
}

// считываю части из папки
func cacheWalk(path string, info os.FileInfo, err error) error {
	if (err != nil) || ((info != nil) && info.IsDir()) {
		//log.Println(err, info)
		return nil
	}
	ext := filepath.Ext(path)
	switch ext {
	case ".php":
		return nil
	}

	keyName := filepath.Base(path)
	if _, ok := getCache(keyName); !ok {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			logs.ErrorLog(err)
			return err
		}
		setCache(keyName, data)
		//log.Println(keyName)
	}
	return nil
}
func cacheFiles() {
	filepath.Walk(filepath.Join(*f_static, "js"), cacheWalk)
	filepath.Walk(filepath.Join(*f_static, "css"), cacheWalk)

	cachePath := *f_chePath
	p := strings.Index(cachePath, ";")
	for p > 0 {

		filepath.Walk(filepath.Join(*f_web, cachePath[:p]), cacheWalk)
		cachePath = cachePath[p+1:]
		p = strings.Index(cachePath, ";")
	}
	filepath.Walk(filepath.Join(*f_web, cachePath), cacheWalk)
}

// rereads files to cache directive
func handlerRecache(w http.ResponseWriter, r *http.Request) {

	emptyCache()
	cacheFiles()
	w.Write([]byte("recache succesfull!"))
}

var (
	f_port    = flag.String("port", ":80", "host address to listen on")
	f_static  = flag.String("path", "./", "path to static files")
	f_web     = flag.String("web", "/home/travel/web/", "path to web files")
	f_session = flag.String("sessionPath", "/var/lib/php/session", "path to store sessions data")
	f_cache   = flag.String("cacheFileExt", `.eot;.ttf;.woff;.woff2;.otf;`, "file extensions for caching HTTPGO")
	f_chePath = flag.String("cachePath", "css;js;fonts;images", "path to cached files")
	F_test = flag.Bool("user8", false, "test mode")
)

func init() {
	flag.Parse()
	ServerConfig := server.GetServerConfig()
	if err := ServerConfig.Init(f_static, f_web, f_session); err != nil {
		logs.ErrorLog(err)
	}

	MongoConfig := server.GetMongodConfig()
	if err := MongoConfig.Init(f_static, f_web, f_session); err != nil {
		logs.ErrorLog(err)
	}
	services.InitServices()
}
func main() {
	users.SetSessionPath(*f_session)
	go cacheFiles()

	fonts.GetPath(f_web)

	registerRoutes()

	logs.StatusLog("Server starting")
	logs.StatusLog("Static files found in ", *f_web)
	logs.StatusLog("System files found in " + *f_static)

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go listenOnShutdown(ch)

	defer func() {
		logs.StatusLog("Server correct shutdown")
	}()
	logs.Fatal(http.ListenAndServe(*f_port, nil))

}
func listenOnShutdown(ch <- chan os.Signal) {
	//var signShut os.Signal
	signShut := <- ch
	logs.StatusLog(signShut.String())
}
