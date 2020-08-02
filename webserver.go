package main

import (
	"fmt"
	"net/http"
	"os"
	"simplehttp/api"
)

func port() string {
	port := os.Getenv("WEBSERVER_PORT")
	if port == "" {
		return ":8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Webserver running at port %v", port())
}

func echo(w http.ResponseWriter, r *http.Request) {
	msgs := r.URL.Query()
	msg := "Default Message Echoed"
	if tempmsg, ok := msgs["message"]; ok {
		msg = tempmsg[0]
	}
	fmt.Fprintf(w, msg)
}

func main() {
	fmt.Println("Webserver running....")
	fmt.Printf("Port: %v", port())

	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)
	http.HandleFunc("/api/books", api.BooksHandler)
	http.HandleFunc("/api/book/", api.BookHandler)
	http.ListenAndServe(port(), nil)
}
