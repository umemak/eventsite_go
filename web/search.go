package web

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/umemak/eventsite_go/model/user"
)

func GetSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	u, err := user.BuildFromContext(r.Context())
	if err != nil {
		log.Printf("user.BuildFromContext: %v", err)
	}
	var buf bytes.Buffer
	err = tpls["search.html"].Execute(&buf, struct {
		Header header
	}{
		Header: header{Title: "検索", User: u},
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("tpls.Execute: %v", err), http.StatusInternalServerError)
		return
	}
	buf.WriteTo(w)
}
