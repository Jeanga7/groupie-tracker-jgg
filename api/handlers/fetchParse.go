package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/api/models"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
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
		return err
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

//________________________________________________________________________________

var indexTpl *template.Template

//var tpl404 *template.Template

func init() {

	timeToWait := 30

	tStart := time.Now()
	var wg sync.WaitGroup
	waitCh := make(chan struct{})
	wg.Add(1)
	log.Printf("Parsing started, if something goes wrong, program will terminate in %v seconds.", timeToWait)
	go func() {
		go func() {
			Parse()
			wg.Done()
		}()
		wg.Wait()
		close(waitCh)
	}()

	select {
	case <-waitCh:
		elapsed := time.Since(tStart)
		log.Printf("Parsing took %.4fs\n", elapsed.Seconds())
		log.Println("Init complete.")
	case <-time.After(time.Duration(timeToWait) * time.Second):
		log.Printf("Parsing failed, terminating\n")
		os.Exit(1)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	indexTpl = template.Must(template.ParseGlob("templates/index/*.html"))
	//tpl404 = template.Must(template.ParseGlob("templates/404/*.html"))
	if r.URL.Path == "/favicon.ico" {
		http.ServeFile(w, r, "static/assets/favicon.ico")
	} else if r.URL.Path != "/" {
		fmt.Println("erreur")
		//callErrorPage(w, r, 404)
		return
	}

	switch r.Method {
	case "GET":
		indexTpl.ExecuteTemplate(w, "index.html", nil)
	default:
		fmt.Println("erreur")
		//callErrorPage(w, r, 405)
	}
}

// func callErrorPage(w http.ResponseWriter, r *http.Request, errorCode int) {
// 	var errorMsg string

// 	switch errorCode {
// 	case 404:
// 		w.WriteHeader(http.StatusNotFound)
// 		errorMsg = "404 Page not found"
// 	case 405:
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		errorMsg = "405 Wrong method"
// 	case 400:
// 		w.WriteHeader(http.StatusBadRequest)
// 		errorMsg = "400 Bad request"
// 	default:
// 		w.WriteHeader(http.StatusInternalServerError)
// 		errorMsg = "500 Internal error"
// 		errorCode = 500
// 	}

// 	// tpl404.ExecuteTemplate(w, "404.html", struct {
// 	// 	ErrorCode int
// 	// 	Error     string
// 	// }{
// 	// 	ErrorCode: errorCode,
// 	// 	Error:     errorMsg,
// 	// })
// }
