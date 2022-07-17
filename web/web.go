package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/go-chi/jwtauth/v5"
	"github.com/koron/go-dproxy"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/umemak/eventsite_go/model/event"
	"github.com/umemak/eventsite_go/model/eventUser"
	"github.com/umemak/eventsite_go/model/user"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/template/index.html")
	if err != nil {
		log.Fatalf("template.ParseFiles: %v", err)
	}

	isLogin := true
	token, _, err := jwtauth.FromContext(r.Context())
	if err != nil {
		isLogin = false
	}
	if token == nil || jwt.Validate(token) != nil {
		isLogin = false
	}
	name := ""
	if isLogin {
		_, claims, _ := jwtauth.FromContext(r.Context())
		name = claims["name"].(string)
	}

	events, err := event.List()
	if err != nil {
		log.Fatalf("event.List: %v", err)
	}
	t.Execute(w, struct {
		Events  []event.Event
		IsLogin bool
		Name    string
	}{
		Events:  events,
		IsLogin: isLogin,
		Name:    name,
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

func GetLogin(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/template/login.html")
	if err != nil {
		log.Fatalf("template.ParseFiles: %v", err)
	}
	t.Execute(w, nil)
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	t_ok, err := template.ParseFiles("web/template/index.html")
	if err != nil {
		log.Fatalf("template.ParseFiles: %v", err)
	}
	t_ng, err := template.ParseFiles("web/template/login.html")
	if err != nil {
		log.Fatalf("template.ParseFiles: %v", err)
	}
	err = r.ParseForm()
	if err != nil {
		log.Fatalf("ParseForm: %v", err)
	}
	url := "http://pocketbase:8090/api/users/auth-via-email"
	jsonString := fmt.Sprintf(`{"email": "%s", "password": "%s"}`,
		r.PostFormValue("email"), r.PostFormValue("password"))
	fmt.Printf("%+v\n", jsonString)
	req, err := http.NewRequest("POST", url, strings.NewReader(jsonString))
	if err != nil {
		log.Printf("NewRequest: %v", err)
		t_ng.Execute(w, nil)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Do: %v", err)
		t_ng.Execute(w, nil)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ReadAll: %v", err)
		t_ng.Execute(w, nil)
		return
	}
	var obj any
	json.Unmarshal(body, &obj)
	st := resp.StatusCode
	if st != http.StatusOK {
		message, err := dproxy.New(obj).M("message").String()
		if err != nil {
			log.Printf("dproxy.New(obj).M(\"message\").String(): %v", err)
			t_ng.Execute(w, nil)
			return
		}
		t_ng.Execute(w, struct {
			Message string
		}{
			Message: message,
		})
		return
	}

	id, err := dproxy.New(obj).M("user").M("id").String()
	if err != nil {
		log.Printf("dproxy.New(obj).M(\"user\").M(\"id\").String(): %v", err)
		t_ng.Execute(w, nil)
		return
	}
	user, err := user.GetByUID(id)
	if err != nil {
		log.Printf("user.GetByUID: %v", err)
		t_ng.Execute(w, nil)
		return
	}
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"name": user.Name})
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	events, err := event.List()
	if err != nil {
		log.Printf("event.List: %v", err)
		t_ng.Execute(w, nil)
		return
	}
	t_ok.Execute(w, struct {
		Events  []event.Event
		IsLogin bool
		Name    string
	}{
		Events:  events,
		IsLogin: true,
		Name:    user.Name,
	})
}

func GetSignup(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/template/signup.html")
	if err != nil {
		log.Fatalf("template.ParseFiles: %v", err)
	}
	t.Execute(w, struct {
		Message string
	}{
		Message: "",
	})
}

func PostSignup(w http.ResponseWriter, r *http.Request) {
	t_ok, err := template.ParseFiles("web/template/login.html")
	if err != nil {
		log.Fatalf("template.ParseFiles: %v", err)
	}
	t_ng, err := template.ParseFiles("web/template/signup.html")
	if err != nil {
		log.Fatalf("template.ParseFiles: %v", err)
	}
	err = r.ParseForm()
	if err != nil {
		log.Fatalf("ParseForm: %v", err)
	}
	url := "http://pocketbase:8090/api/users"
	jsonString := fmt.Sprintf(`{"email": "%s", "password": "%s", "passwordConfirm": "%s"}`,
		r.PostFormValue("email"), r.PostFormValue("password"), r.PostFormValue("password"))
	fmt.Printf("%+v\n", jsonString)
	req, err := http.NewRequest("POST", url, strings.NewReader(jsonString))
	if err != nil {
		log.Printf("NewRequest: %v", err)
		t_ng.Execute(w, nil)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Do: %v", err)
		t_ng.Execute(w, nil)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ReadAll: %v", err)
		t_ng.Execute(w, nil)
		return
	}
	var obj any
	json.Unmarshal(body, &obj)
	st := resp.StatusCode
	if st != http.StatusOK {
		message, err := dproxy.New(obj).M("message").String()
		if err != nil {
			t_ng.Execute(w, nil)
			return
		}
		t_ng.Execute(w, struct {
			Message string
		}{
			Message: message,
		})
		return
	}

	id, err := dproxy.New(obj).M("id").String()
	if err != nil {
		t_ng.Execute(w, nil)
		return
	}
	name := r.PostFormValue("name")
	_, err = user.Create(user.User{UID: id, Name: name})
	if err != nil {
		log.Fatalf("user.Create: %v", err)
	}
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"name": name})
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	t_ok.Execute(w, struct {
		User user.User
	}{
		User: user.User{UID: id},
	})
}

func GetLogout(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/template/index.html")
	if err != nil {
		log.Fatalf("template.ParseFiles: %v", err)
	}
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	t.Execute(w, nil)
}

func GetSearch(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/template/search.html")
	if err != nil {
		log.Fatalf("template.ParseFiles: %v", err)
	}
	t.Execute(w, nil)
}

func GetEventCreate(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/template/event_create.html")
	if err != nil {
		log.Fatalf("template.ParseFiles: %v", err)
	}
	t.Execute(w, nil)
}

func GetEventDetail(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/template/event_detail.html")
	if err != nil {
		log.Fatalf("template.ParseFiles: %v", err)
	}
	values := r.URL.Query()
	id, err := strconv.ParseInt(values.Get("id"), 10, 64)
	e, err := event.Find(id)
	if err != nil {
		log.Fatalf("event.Find: %v", err)
	}
	eu, err := eventUser.FindByEvent(id)
	if err != nil {
		log.Fatalf("event.Find: %v", err)
	}
	t.Execute(w, struct {
		Event      event.Event
		EventUsers []eventUser.EventUser
	}{
		Event:      *e,
		EventUsers: eu,
	})
	t.Execute(w, nil)
}
