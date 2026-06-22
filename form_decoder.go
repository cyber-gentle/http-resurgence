package main

import (
	"fmt"
	"net/http"
)

func FormDecoderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not aloowed", http.StatusMethodNotAllowed)
	}
	//r.ParseForm() + r.Form.Get()
	username := r.FormValue("username")
	language := r.FormValue("language")

	if username == "" && language == "" {
		http.Error(w,
			"username is required\nlanguage is required\n",
			http.StatusBadRequest)
	} else if language == "" {
		http.Error(w, "language is required\n", http.StatusBadRequest)
	} else if username == "" {
		http.Error(w, "username is required\n", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Hello %s. you are coding in %s\n", username, language)
}

