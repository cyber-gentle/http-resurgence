package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "code parameter is required", http.StatusBadRequest)
	}
	codeStatus, err := strconv.Atoi(code)
	if err != nil {
		http.Error(w, "code must be a valid integer", http.StatusBadRequest)
	}

	w.WriteHeader(codeStatus)

	if codeStatus < 100 || codeStatus > 599 {
		http.Error(w, "code must be a valid HTTP status code (100–599)", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Responding with status %s\n", code)
}

// if the header isn't set up before writing to response response,
// automatically the header is set to statusOk.
