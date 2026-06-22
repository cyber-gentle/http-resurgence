package main

import (
	"fmt"
	"io"
	"net/http"
)

func EchoChamberHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, string(body))
}

