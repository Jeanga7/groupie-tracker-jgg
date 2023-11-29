package handlers

import (
	"fmt"
	"groupie-tracker/api/models"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage")
}


func HandleRequests() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/artists", models.AllArtists)
	fmt.Println("localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
