package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("index.html"))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/solve", solveHandler)

	http.ListenAndServe(":8080", nil)
}

func solveHandler(w http.ResponseWriter, r *http.Request) {
	// Implement the puzzle solving logic here
	// For now, just return a placeholder response
	w.Write([]byte("Puzzle solved!"))
}
