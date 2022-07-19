package web

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/umemak/eventsite_go/model/user"
)

func GetSignup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	u, err := user.BuildFromContext(r.Context())
	if err != nil {
		log.Printf("user.BuildFromContext: %v", err)
	}
	var buf bytes.Buffer
	err = tpls["signup.html"].Execute(&buf, struct {
		Header  header
		Message string
	}{
		Header:  header{Title: "ユーザー登録", User: u},
		Message: "",
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("tpls.Execute: %v", err), http.StatusInternalServerError)
		return
	}
	buf.WriteTo(w)
}

func PostSignup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	html_ok := "login.html"
	html_ng := "signup.html"
	u, err := user.BuildFromContext(r.Context())
	if err != nil {
		log.Printf("user.BuildFromContext: %v", err)
	}
	err = r.ParseForm()
	if err != nil {
		log.Fatalf("ParseForm: %v", err)
	}
	_, err = user.Create(
		r.PostFormValue("email"), r.PostFormValue("password"),
		r.PostFormValue("password"), r.PostFormValue("name"),
	)
	if err != nil {
		var buf bytes.Buffer
		err = tpls[html_ng].Execute(&buf, struct {
			Header  header
			Message string
		}{
			Header:  header{Title: "ユーザー登録", User: u},
			Message: err.Error(),
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("tpls.Execute: %v", err), http.StatusInternalServerError)
			return
		}
		buf.WriteTo(w)
		return
	}
	var buf bytes.Buffer
	err = tpls[html_ok].Execute(&buf, struct {
		Header  header
		Message string
	}{
		Header:  header{Title: "ログイン", User: u},
		Message: "ユーザー作成成功。ログインしてください。",
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("tpls.Execute: %v", err), http.StatusInternalServerError)
		return
	}
	buf.WriteTo(w)
}
