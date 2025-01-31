package main

import (
	"log"
	"net/http"
)

func main() {
	// creating new router
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("../../ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/game/show", gameShow)
	mux.HandleFunc("/game/add", gameAdd)

	// listening on a web-server
	log.Print("Starting web server on a localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
