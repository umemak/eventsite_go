package web

import (
	"context"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"text/template"

	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/umemak/eventsite_go/model/user"
)

const PB_URL = "http://pocketbase:8090"

var TokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

var tpls = map[string]*template.Template{}

func init() {
	files, err := filepath.Glob(path.Join("web", "template", "*.html"))
	if err != nil {
		log.Fatalf("filepath.Glob: %v", err)
	}
	for _, file := range files {
		tpls[filepath.Base(file)] = template.Must(template.ParseFiles(file))
	}
}

func isLogin(ctx context.Context) bool {
	token, _, err := jwtauth.FromContext(ctx)
	if err != nil {
		return false
	}
	if token == nil || jwt.Validate(token) != nil {
		return false
	}
	return true
}

func GetLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		HttpOnly: true,
	})
	err := tpls["index.html"].Execute(w, nil)
	if err != nil {
		log.Fatalf("Execute: %v", err)
	}
}

func makeToken(user *user.User) string {
	if user == nil {
		return ""
	}
	_, tokenString, _ := TokenAuth.Encode(map[string]interface{}{
		"id":   user.ID,
		"uid":  user.UID,
		"name": user.Name,
	})
	return tokenString
}
