package main

import (
	"log"
	"net/http"
)

func home(response http.ResponseWriter, request *http.Request) {

	if request.URL.Path != "/" {
		http.NotFound(response, request)
		return
	}

	response.Write([]byte("Hello from Snippetbox"))
}

func snippetView(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
