package main

import (
	"groupie-tracker/pkg/api/handlers"
)

func main() {
	handlers.HandleRequests()
	/* fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", http.FileServer(http.Dir("./static")))
	//http.HandleFunc("/artists", )
	//http.HandleFunc("/locations", handlers.HandleLocations())
	http.ListenAndServe(":8080", nil) */
}
