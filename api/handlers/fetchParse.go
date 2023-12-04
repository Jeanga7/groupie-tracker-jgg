package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/api/models"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

var url = "https://groupietrackers.herokuapp.com/api"
var cache models.Cache

func Fetch(url string, pointer interface{}) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(responseData, &pointer)

}

func Parse() {
	var wg sync.WaitGroup

	Fetch(url, &cache)
	wg.Add(1)
	go func() {
		go Fetch(cache.ArtistsURI, &cache.Artists)
		go Fetch(cache.LocationsURI, &cache.Locations)
		go Fetch(cache.DatesURI, &cache.Dates)
		go Fetch(cache.RelationURI, &cache.Relation)
		wg.Done()
	}()
	wg.Wait()
}
