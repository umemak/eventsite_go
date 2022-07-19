package web

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/umemak/eventsite_go/model/user"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	u, err := user.BuildFromContext(r.Context())
	if err != nil {
		log.Printf("user.BuildFromContext: %v", err)
	}
	var buf bytes.Buffer
	err = tpls["login.html"].Execute(&buf, struct {
		Header  header
		Message string
	}{
		Header:  header{Title: "ログイン", User: u},
		Message: "",
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("tpls.Execute: %v", err), http.StatusInternalServerError)
		return
	}
	buf.WriteTo(w)
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	html_ng := "login.html"
	uh, err := user.BuildFromContext(r.Context())
	if err != nil {
		log.Printf("user.BuildFromContext: %v", err)
	}
	err = r.ParseForm()
	if err != nil {
		log.Fatalf("ParseForm: %v", err)
	}
	u, err := user.AuthViaEmail(r.PostFormValue("email"), r.PostFormValue("password"))
	if err != nil {
		var buf bytes.Buffer
		err = tpls[html_ng].Execute(&buf, struct {
			Header  header
			Message string
		}{
			Header:  header{Title: "ログイン", User: uh},
			Message: err.Error(),
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("tpls.Execute: %v", err), http.StatusInternalServerError)
			return
		}
		buf.WriteTo(w)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    makeToken(u),
		HttpOnly: true,
	})
	http.Redirect(w, r, "/", 302)
}
