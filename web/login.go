package web

import (
	"log"
	"net/http"

	"github.com/umemak/eventsite_go/model/user"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {
	err := tpls["login.html"].Execute(w, struct {
		Message string
	}{
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
			Message string
		}{
			Message: err.Error(),
		})
		if err != nil {
			log.Fatalf("Execute: %v", err)
		}
		return
	}
	user, err := user.GetByUID(u.UID)
	if err != nil {
		log.Printf("user.GetByUID: %v", err)
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
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    makeToken(user),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", 301)
}
