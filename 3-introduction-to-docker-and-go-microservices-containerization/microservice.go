package main

import (
	"net/http"
	"os"

	"github.com/avisiedo/go-cloud-native/api"
)

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func main() {
	http.HandleFunc("/", api.HelloHandleFunc)
	http.HandleFunc("/api/echo", api.EchoHandleFunc)
	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.HandleFunc("/api/books/", api.BookHandleFunc)
	http.ListenAndServe(port(), nil)
}
