package main

import (
	"fmt"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Stranger"
		fmt.Fprintf(w, "Greetings, %s!\n", name)
	} else {
		fmt.Fprintf(w, "Greetings %s!\n", name)
	}

}
