package users

import (
	"github.com/ruslanBik4/httpgo/views"
	"github.com/ruslanBik4/httpgo/views/templates/layouts"
	"github.com/ruslanBik4/httpgo/views/templates/forms"
	"github.com/ruslanBik4/httpgo/views/templates/mails"
	"github.com/ruslanBik4/httpgo/models/db"
	"net/http"
	"fmt"
	"strconv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
	"log"
	"io/ioutil"
	//"gopkg.in/gomail.v2"
	"net/mail"
	"crypto/rand"
	"encoding/base64"
	"hash/crc32"
	"github.com/gorilla/sessions"
	"gopkg.in/gomail.v2"
	"github.com/ruslanBik4/httpgo/models/system"
)
const nameSession = "PHPSESSID"
const NOT_AUTHORIZE = "Нет данных об авторизации!"
var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "",
		ClientID:     os.Getenv("googlekey"),
		ClientSecret: os.Getenv("googlesecret"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	oauthStateString = "random"
	Store = sessions.NewFilesystemStore("/var/lib/php/session",[]byte("travel.com.ua"))

)
func SetSessionPath(f_session string) {
	Store = sessions.NewFilesystemStore(f_session,[]byte("travel.com.ua"))
}
//func HandlerQauth2(w http.ResponseWriter, r *http.Request) {
//
//
//	googleOauthConfig.RedirectURL = r.Host +  "/user/GoogleCallback/"
//	url := googleOauthConfig.AuthCodeURL(oauthStateString)
//	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
//	//var ctx context.Context = appengine.NewContext(r)
//	//client := &http.Client{
//	//	Transport: &oauth2.Transport{
//	//		Source: google.AppEngineTokenSource(ctx, "scope"),
//	//		Base:   &urlfetch.Transport{Context: ctx},
//	//	},
//	//}
//	//resp, _ := client.Get("...")
//	//w.Write(resp.Body)
//}
//Эти callback было бы неплохо регистрировать в одну общую библиотеку для авторизации
func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Printf("Code exchange failed with '%v'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	fmt.Fprintf(w, "Content: %s\n", contents)
}
type UserRecord struct {
	Id int
	Name string
	Sex int
}
var greetings = [] string {"господин", "госпожа"}

func GetSession(r *http.Request, name string) *sessions.Session {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, err := Store.Get(r, name)
	if err != nil {
		log.Println(err)
		return nil
	}
	return session
}
func IsLogin(r *http.Request) string {
	session := GetSession(r, nameSession)
	if session == nil {
		panic(http.ErrNotSupported)
	}
	if userID, ok := session.Values["id"]; ok {

		return strconv.Itoa(userID.(int))
	} else {
		panic(system.ErrNotLogin{Message:"not login user!"})
	}

	return ""
}
func deleteCurrentUser(w http.ResponseWriter, r *http.Request) error {
	session := GetSession(r, nameSession )
	delete(session.Values, "id")
	delete(session.Values, "email")
	return session.Save(r, w)

}
func HandlerProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	session := GetSession(r, nameSession )
	email, ok := session.Values["email"]
	if !ok {
		http.Redirect(w,r, "/show/forms/?name=signin", http.StatusSeeOther)
		return
	}
	rows := db.DoQuery("select id, fullname, sex from users where login=?", email )

	var row UserRecord

	defer rows.Close()
	for rows.Next() {

		err := rows.Scan(&row.Id, &row.Name, &row.Sex)

		if err != nil {
			log.Println(err)
			continue
		}
	}

	p := &layouts.MenuOwnerBody{ Title: greetings[row.Sex] + " " + row.Name, TopMenu: make(map[string] *layouts.ItemMenu, 0)}

	var menu db.MenuItems

	menu.GetMenu("menuOwner")

	for _, item := range menu.Items {
		p.TopMenu[item.Title] = &layouts.ItemMenu{ Link: "/menu/" + item.Name + "/"  }

	}
	fmt.Fprint(w, p.MenuOwner() )
}
func HandlerSignIn(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	email := r.FormValue("login")
	password := r.FormValue("password")

	if (email == "") || (password == "") {
		panic(&system.ErrNotLogin{Message:"Not enoug login parameters!"})
	}

	err, userId, userName := CheckUserCredentials(email, password)

	if err != nil {
		panic(&system.ErrNotLogin{Message:"Wrong email or password"})
	}

	// session save BEFORE write page
	SaveSession(w, r, userId, email)

	p := &forms.PersonData{ Id: userId, Login: userName, Email: email }
	fmt.Fprint(w, p.JSON())
}
func HandlerSignOut(w http.ResponseWriter, r *http.Request) {

	if err := deleteCurrentUser(w, r); err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, "<title>%s</title>", "Для начала работы необходимо авторизоваться!" )
	views.RenderSignForm(w, r, "")
}

func SaveSession(w http.ResponseWriter, r *http.Request, id int, email string) {
	session := sessions.NewSession(Store, nameSession)
	session.Options = &sessions.Options{Path: "/", HttpOnly: true, MaxAge: int(3600)}
	session.Values["id"] = id
	session.Values["email"] = email
	if err := session.Save(r, w); err != nil {
		log.Println(err)
	}
}
// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
func GeneratePassword(email string) (string, error) {

	log.Println(email)
	return GenerateRandomString(16)

}
func HashPassword(password string) interface{} {
	// crypto password
	crc32q := crc32.MakeTable(0xD5828281)
return 	crc32.Checksum([]byte(password), crc32q)
}
func HandlerSignUp(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32000)

	var args [] interface{}
	sql, comma, values := "insert into users (", "", ") values ("

	for key, val := range r.MultipartForm.Value {
		args = append(args, val[0])
		sql += comma + key
		values += comma + "?"
		comma = ","
	}
	email := r.MultipartForm.Value["login"][0]
	password, err := GeneratePassword(email)
	if err != nil {
		log.Println(err)
	}
	sql += comma + "hash"
	values += comma + "?"

	args = append(args, HashPassword(password) )
	lastInsertId, err := db.DoInsert(sql + values + ")", args... )
	if err != nil {

		fmt.Fprintf(w, "%v", err)
		return
	}
	w.Header().Set("Content-Type", "text/json; charset=utf-8")

	mRow := forms.MarshalRow{Msg: "Append row", N: lastInsertId}
	sex, _ := strconv.Atoi(r.MultipartForm.Value["sex"][0])

	if _, err := mail.ParseAddress(email); err !=nil {
		log.Println(err)
		fmt.Fprintf(w, "Что-то неверное с вашей почтой, не смогу отослать письмо! %v", err)
		return
	}
	p := &forms.PersonData{ Id: lastInsertId, Login: r.MultipartForm.Value["fullname"][0], Sex: sex,
		Rows: []forms.MarshalRow{mRow}, Email: email }
	fmt.Fprint(w, p.JSON())

	go SendMail(email, password)
}
func SendMail(email, password string)  {

	m := gomail.NewMessage()
	m.SetHeader("From", "ruslan-bik@yandex.ru")
	m.SetHeader("To", email )
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Регистрация на travel.com.ua!")
	m.SetBody("text/html", mails.InviteEmail(email, password) )
	m.Attach("/home/travel/bootstrap/ico/favicon.png")

	d := gomail.NewDialer("smtp.yandex.ru", 587, "ruslan-bik", "FalconF99")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
	}
}
func HandlerActivateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm();

	if r.FormValue("email") == "" {
		log.Println("activate user not has email")
	}
	if result, err := db.DoUpdate("update users set active=1 where login=?", r.Form["email"][0]); err != nil {
		log.Println(err)
	} else {
		fmt.Fprint(w, result)

	}
}

func CheckUserCredentials(login string, password string) (error, int, string) {

	rows, err := db.DoSelect("select id, fullname, sex from users where login=? and hash=?", login, HashPassword(password) )
	if err != nil {
		log.Println(err)
		return err, 0, ""
	}
	defer rows.Close()
	var row UserRecord

	for rows.Next() {

		err := rows.Scan(&row.Id, &row.Name, &row.Sex)

		if err != nil {
			log.Println(err)
			continue
		}

		return nil, row.Id, row.Name
	}

	return &system.ErrNotLogin{Message:"Wrong email or password"}, 0, ""
}


