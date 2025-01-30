package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// files' urls
	files := []string{
		"C:\\Users\\frogi\\GolandProjects\\game_market\\ui\\html\\base.tmpl.html",
		"C:\\Users\\frogi\\GolandProjects\\game_market\\ui\\html\\pages\\home.tmpl.html",
	}

	// template set
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func gameShow(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Showing game with id:%d", id)
}
func gameAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Adding a specific game"))
}
