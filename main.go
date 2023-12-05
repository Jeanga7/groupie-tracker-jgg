package main

import (
	"fmt"
	"groupie-tracker/api/handlers"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/artists", handlers.GetArtists)

	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(port, router))
}
