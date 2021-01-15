package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	var html string
	html = "<h1>Universidad de las Fuerzas Armadas ESPE</h1>"
	html += "<h4>Nombre: Ariel Altamirano</h4><h4>NRC: 3138</h4>"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w,html)
}