package handlers

import (
	"fmt"
	"groupie-tracker/pkg/api/models"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage")
}

func HandleRequests() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/artists", models.AllArtists)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
