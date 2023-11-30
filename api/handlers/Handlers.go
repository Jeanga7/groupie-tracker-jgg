package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/api/models"
	"html/template"
	"io"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage")
}

func HandleRequests() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/artists", AllArtists)
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
func AllArtists(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.URL.Path)
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var artists []models.Artist
	err = json.Unmarshal(body, &artists)

	if err != nil {
		panic(err)
	}

	tmpl, err := template.ParseFiles("./templates/base.html", "./templates/artists.html")
	if err != nil {
		fmt.Println("lecture de page")
	}
	err = tmpl.ExecuteTemplate(w, "content", artists)
	if err != nil {
		panic(err)
	}
}
