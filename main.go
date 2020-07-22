package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8001",
	}
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	server.ListenAndServe()
}
