package fonts

import (
	"net/http"
	"strings"
)
var 	PathWeb string
func GetPath(path string) {
	PathWeb = path
}
func contains(array [] string, str string) bool {
	for _, value := range array {
		if strings.Contains(value, str) {
			return true
		}
	}

	return false
}
func HandleGetFont(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Content-Type", "mime/type; ttf")

	//log.Println(PathWeb + r.URL.Path)
	ext := ".ttf"
	if browser:= r.Header["User-Agent"]; contains(browser, "Safari") {
		ext = ".woff"
	}
	http.ServeFile(w, r, PathWeb + r.URL.Path + ext)
}
