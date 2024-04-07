package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello")

}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseForm() error %v", err)
		return
	}
	fmt.Fprintf(w, "POST was successful")
	username := r.FormValue("user_name")
	email := r.FormValue("email")

	fmt.Fprintf(w, "Username %s\n", username)
	fmt.Fprintf(w, "Email %s\n", email)

}

func main() {
	fileReader := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileReader)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting the server at 8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
