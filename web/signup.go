package web

import (
	"log"
	"net/http"

	"github.com/umemak/eventsite_go/model/user"
)

func GetSignup(w http.ResponseWriter, r *http.Request) {
	err := tpls["signup.html"].Execute(w, struct {
		Message string
	}{
		Message: "",
	})
	if err != nil {
		log.Fatalf("Execute: %v", err)
	}
}

func PostSignup(w http.ResponseWriter, r *http.Request) {
	html_ok := "login.html"
	html_ng := "signup.html"
	err := r.ParseForm()
	if err != nil {
		log.Fatalf("ParseForm: %v", err)
	}
	u, err := user.Create(
		r.PostFormValue("email"), r.PostFormValue("password"),
		r.PostFormValue("password"), r.PostFormValue("name"),
	)
	if err != nil {
		err = tpls[html_ng].Execute(w, struct {
			Message string
		}{
			Message: err.Error(),
		})
		if err != nil {
			log.Fatalf("Execute: %v", err)
		}
		return
	}
	err = tpls[html_ok].Execute(w, struct {
		User *user.User
	}{
		User: u,
	})
	if err != nil {
		log.Fatalf("Execute: %v", err)
	}
}
