package middleware

import (
	"fmt"
	"net/http"
	"text/template"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	t.Execute(w, nil)
}
