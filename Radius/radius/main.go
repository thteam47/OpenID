package main

import (
	"net/http"
	"text/template"
)

type Login struct {
	Username string
	Password string
}

func main() {
	tmp := template.Must(template.ParseFiles("/radius/template.html"))
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmp.Execute(rw, nil)
			return
		}
		tmp.Execute(rw, struct{ succes bool }{true})
	})
	http.ListenAndServe(":8080", nil)
}
