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
		log.Fatalf("template error: %v", err)
	}
	_, err = user.Create(user.User{Name: "testuser"})
	if err != nil {
		log.Fatal(err)
	}
	users, err := user.Find()
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, struct {
		Users []user.User
	}{
		Users: users,
	})
}
