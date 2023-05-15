package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(response, request)
		return
	}

	response.Write([]byte("Hello from Snippetbox"))
}

func snippetView(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(response, request)
		return
	}

	fmt.Fprintf(response, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		response.Header().Set("Allow", http.MethodPost)
		http.Error(response, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	response.Write([]byte("Create a new snippet..."))
}
