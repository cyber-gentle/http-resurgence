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
	mainMux := http.NewServeMux()

	apiMux := http.NewServeMux()

	mainMux.HandleFunc("/status", statusHandler)
	mainMux.HandleFunc("/echo", EchoChamberHandler)
	mainMux.HandleFunc("/form", FormDecoderHandler)
	mainMux.HandleFunc("/headers", HeaderDetectiveHandler)
	mainMux.HandleFunc("/method-inspector", MethodInspectorHandler)

	apiMux.HandleFunc("/v1/ping", SubHandler)
	apiMux.HandleFunc("/v1/greet", SubHandler)

	mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))
	log.Fatal(http.ListenAndServe(":8080", mainMux))
}
