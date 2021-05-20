package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var b = a + 2

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "login")
}

func getMsgHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/msgs.html")
	m := map[string]interface{}{
		"messages": messages,
		"hello":    "helo",
	}
	t.Execute(w, m)
}

func addMsgHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.FormValue("msg")
	messages = append(messages, msg)
	http.Redirect(w, r, "/get_msg/", 303)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.URL)
	t, _ := template.ParseFiles("templates/index.html")
	page := Page{Title: "title", Body: "body"}
	d := map[string]interface{}{
		"page":  page,
		"title": page.Title,
		"body":  "body",
	}
	t.Execute(w, d)
}
