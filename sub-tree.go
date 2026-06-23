package main

import (
	"fmt"
	"net/http"
)

func SubHandler(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/v1/ping":
		w.WriteHeader(200)
		w.Write([]byte("pong"))

	case "/v1/greet":
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
}
