package models

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Index struct {
	Relations []Relation `json:"index"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func AllArtists(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var artists []Artist
	err = json.Unmarshal(body, &artists)

	if err != nil {
		panic(err)
	}

	tmpl, err := template.ParseFiles("../web/templates/index.html")
	if err != nil {
		return
	}
	err = tmpl.Execute(w, artists)
	if err != nil {
		panic(err)
	}
}
