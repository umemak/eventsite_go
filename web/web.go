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
	"github.com/umemak/eventsite_go/sqlc"
)

const PB_URL = "http://pocketbase:8090"

var TokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

var tpls = map[string]*template.Template{}

type header struct {
	Title string
	User  sqlc.User
}

func init() {
	files, err := filepath.Glob(path.Join("web", "template", "*.html"))
	if err != nil {
		log.Fatalf("filepath.Glob: %v", err)
	}
	for _, file := range files {
		tpls[filepath.Base(file)] = template.Must(template.ParseFiles(file,
			path.Join("web", "template", "modules", "_header.html"),
			path.Join("web", "template", "modules", "_footer.html")))
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
	http.Redirect(w, r, "/", 302)
}

func makeToken(user sqlc.User) string {
	_, tokenString, _ := TokenAuth.Encode(map[string]interface{}{
		"id":   user.ID,
		"uid":  user.Uid,
		"name": user.Name,
	})
	return tokenString
}
