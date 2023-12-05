package handlers

import (
	"encoding/json"
	"groupie-tracker/api/models"
	"io"
	"log"
	"net/http"
	"sync"
)

var url = "https://groupietrackers.herokuapp.com/api"
var cache models.Cache

func Fetch(url string, pointer interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(responseData, &pointer)
	if err != nil {
		return err
	}
	return nil
}

func Parse() {
	var wg sync.WaitGroup

	Fetch(url, &cache)
	wg.Add(1)
	go func() {
		go Fetch(cache.ArtistsURL, &cache.Artists)
		// go Fetch(cache.LocationsURL, &cache.Locations)
		// go Fetch(cache.DatesURL, &cache.Dates)
		// go Fetch(cache.RelationURL, &cache.Relation)
		wg.Done()
	}()
	wg.Wait()
}
