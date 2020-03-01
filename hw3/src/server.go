package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	port := "8080"
	router.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	log.Printf("start listen on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
