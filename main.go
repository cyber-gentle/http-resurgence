// 1. using mainMux without the trailing slash will not
// lead to the right page cause no main path that leads
// to that path, just like no parent name that can link to the child.

// 2. http.StripPrefix trims homepage url path to avoid double writing e.g
// /api/api//v1/ping

package main

import (
	"log"
	"net/http"
)

func main() {
	Server()
}

func Server() {
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/echo", EchoChamberHandler)
	http.HandleFunc("/form", FormDecoderHandler)
	http.HandleFunc("/headers", HeaderDetectiveHandler)
	http.HandleFunc("/method-inspector", MethodInspectorHandler)

	apiMux := http.NewServeMux()

	apiMux.HandleFunc("/v1/ping", pingHandler)
	apiMux.HandleFunc("/v1/greet", greetHandler)

	mainMux := http.NewServeMux()
	mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))
	log.Fatal(http.ListenAndServe(":8080", mainMux))
}
