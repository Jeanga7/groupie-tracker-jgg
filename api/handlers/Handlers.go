package handlers

import (
	"encoding/json"
	"groupie-tracker/api/models"
	"log"
	"net/http"
	"strconv"
)

func GetArtists(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:

		artistsAmount := r.FormValue("artists-amount")

		if artistsAmount == "" {
			http.Error(w, `"artists-amount" variable is required`, http.StatusBadRequest)
			return
		}

		amount, err := strconv.Atoi(artistsAmount)
		if err != nil {
			log.Println("Error during atoi conversion. Error:", err)
			http.Error(w, "Invalid artists-amount value", http.StatusBadRequest)
			return
		}

		dataArr := getArtistsData(amount)

		b, err := json.Marshal(dataArr)
		if err != nil {
			log.Println("Error during JSON marshalling. Error:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write(b)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("This function does not support " + r.Method + " method."))
	}
}

func getArtistsData(amount int) []models.Data {
	var dataArr []models.Data

	for i := 0; i < amount; i++ {
		dataArr = append(dataArr, models.Data{
			ArtistsID: cache.Artists[i].ID,
			Name:      cache.Artists[i].Name,
		})
	}

	return dataArr
}
