package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"warungpintar/handler"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", handler.Home)
	r.HandleFunc("/welcome", handler.Welcome).Queries("name", "{name}").Methods("GET")
	r.HandleFunc("/history", handler.History).Methods("GET")
	staticDir := "/score/"
	r.PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir,
			http.FileServer(http.Dir("."+staticDir))))
	r.HandleFunc("/update-score", handler.Score)
	// swag := http.FileServer(http.Dir("./swaggerui/"))
	// r.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", swag))
	swag := "/swaggerui/"
	r.PathPrefix(swag).Handler(http.StripPrefix(swag,
		http.FileServer(http.Dir("."+swag))))

	go handler.HandleMessages()

	log.Fatal(http.ListenAndServe(":8080", r))
}
