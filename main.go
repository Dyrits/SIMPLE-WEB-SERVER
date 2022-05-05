package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "Path is incorrect!", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(writer, "Method is not supported!", http.StatusNotFound)
		return
	}
	fmt.Fprintln(writer, "Hello!")
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/form" {
		http.Error(writer, "Path is incorrect!", http.StatusNotFound)
		return
	}
	if request.Method != "POST" {
		http.Error(writer, "Method is not supported!", http.StatusNotFound)
		return
	}
	if failure := request.ParseForm(); failure != nil {
		fmt.Fprintln(writer, "An error occurred while parsing the form.", failure)
		return
	}
	fmt.Fprintln(writer, "Request successful!")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintln(writer, "Name = "+name)
	fmt.Fprintln(writer, "Address = "+address)
}

func main() {
	server := http.FileServer(http.Dir("./static"))
	http.Handle("/", server)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Starting server on port 8080")
	if failure := http.ListenAndServe(":8080", nil); failure != nil {
		log.Fatal(failure)
	}
}
