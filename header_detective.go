package main

import (
	"fmt"
	"net/http"
)

func HeaderDetectiveHandler(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("X-Custom-Token") == "" {
		http.Error(w, "X-Custom-Token header is missing", http.StatusBadRequest)
	}

	fmt.Fprintf(w, "Token received: %s\n", r.Header.Get("X-Custom-Token"))

	if r.Header.Get("Content-Type") == "" {
		fmt.Fprint(w, "Content-Type not provided")
	}

	fmt.Fprint(w, r.Header.Get("Content-Type"))

}
