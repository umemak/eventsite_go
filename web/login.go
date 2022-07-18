package web

import (
	"log"
	"net/http"

	"github.com/umemak/eventsite_go/model/user"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {
	err := tpls["login.html"].Execute(w, struct {
		Header  header
		Message string
	}{
		Header:  header{Title: "ログイン"},
		Message: "",
	})
	if err != nil {
		log.Fatalf("Execute: %v", err)
	}
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	html_ng := "login.html"
	err := r.ParseForm()
	if err != nil {
		log.Fatalf("ParseForm: %v", err)
	}
	u, err := user.AuthViaEmail(r.PostFormValue("email"), r.PostFormValue("password"))
	if err != nil {
		log.Printf("AuthViaEmail: %v", err)
		err = tpls[html_ng].Execute(w, struct {
			Header  header
			Message string
		}{
			Header:  header{Title: "ログイン"},
			Message: err.Error(),
		})
		if err != nil {
			log.Fatalf("Execute: %v", err)
		}
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    makeToken(u),
		HttpOnly: true,
	})
	http.Redirect(w, r, "/", 302)
}
