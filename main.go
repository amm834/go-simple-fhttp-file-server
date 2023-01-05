package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server is starting at port :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "Parse Form failed: %v", err)
		return
	}

	fmt.Fprintf(writer, "Post success\n")
	email := request.FormValue("email")
	password := request.FormValue("password")

	fmt.Fprintf(writer, "email: %s\n", email)
	fmt.Fprintf(writer, "password: %s\n", password)
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "404 | Not Found", http.StatusNotFound)
		return
	}

	if request.Method != "GET" {
		http.Error(writer, "405 | Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(writer, "hello!")
}
