package main

import (
	"net/http"
	"os"
)

func dataHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello je suis Lucas"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/getData", dataHandler)
	http.ListenAndServe(":"+port, nil)
}
