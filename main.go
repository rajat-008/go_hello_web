package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var messages = make([]string, 0, 10)

type Page struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var a = 0

func main() {
	log.SetOutput(os.Stdout)
	log.Println("server started")
	router := mux.NewRouter()
	router.Use(
		loggingMiddleware,
		nothingMiddleware,
	)
	router.StrictSlash(true)
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/add_msg", addMsgHandler)
	router.HandleFunc("/get_msg", getMsgHandler)
	http.ListenAndServe(":8000", router)
}
