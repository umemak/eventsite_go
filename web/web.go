package web

import (
	"log"
	"net/http"
	"text/template"

	"github.com/umemak/eventsite_go/model/user"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/template/index.html")
	if err != nil {
		log.Fatalf("template.ParseFiles: %v", err)
	}
	users, err := user.List()
	if err != nil {
		log.Fatalf("user.List: %v", err)
	}
	t.Execute(w, struct {
		Users []user.User
	}{
		Users: users,
	})
}

func PostRoot(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/template/index.html")
	if err != nil {
		log.Fatalf("template.ParseFiles: %v", err)
	}
	err = r.ParseForm()
	if err != nil {
		log.Fatalf("ParseForm: %v", err)
	}
	_, err = user.Create(user.User{Name: r.PostFormValue("name")})
	if err != nil {
		log.Fatalf("user.Create: %v", err)
	}
	users, err := user.List()
	if err != nil {
		log.Fatalf("user.List: %v", err)
	}
	t.Execute(w, struct {
		Users []user.User
	}{
		Users: users,
	})
}
