package web

import (
	"log"
	"net/http"

	"github.com/umemak/eventsite_go/model/user"
)

func GetSignup(w http.ResponseWriter, r *http.Request) {
	u, err := user.BuildFromContext(r.Context())
	if err != nil {
		log.Printf("user.BuildFromContext: %v", err)
	}
	err = tpls["signup.html"].Execute(w, struct {
		Header  header
		Message string
	}{
		Header:  header{Title: "ユーザー登録", User: u},
		Message: "",
	})
	if err != nil {
		log.Fatalf("Execute: %v", err)
	}
}

func PostSignup(w http.ResponseWriter, r *http.Request) {
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
		err = tpls[html_ng].Execute(w, struct {
			Header  header
			Message string
		}{
			Header:  header{Title: "ユーザー登録", User: u},
			Message: err.Error(),
		})
		if err != nil {
			log.Fatalf("Execute: %v", err)
		}
		return
	}
	err = tpls[html_ok].Execute(w, struct {
		Header  header
		Message string
	}{
		Header:  header{Title: "ログイン", User: u},
		Message: "ユーザー作成成功。ログインしてください。",
	})
	if err != nil {
		log.Fatalf("Execute: %v", err)
	}
}
