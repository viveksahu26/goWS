package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleFunc)
	fmt.Println("Server Listening on: 3300...")
	log.Fatal(http.ListenAndServe(":3300", nil))
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "<h1> Welcome %v to my awesome site! </h1>", name)
}
