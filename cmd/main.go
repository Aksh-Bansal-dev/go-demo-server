package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"example.com/go-demo-server/pkg/routes"
)

func main() {
    r := mux.NewRouter().PathPrefix("/api").Subrouter()

    r.HandleFunc("/movies",routes.GetMoviesHandler).Methods("GET")
    r.HandleFunc("/movies/{id}",routes.GetMovieHandler).Methods("GET")
    r.HandleFunc("/movies",routes.AddMovieHandler).Methods("POST")
    r.HandleFunc("/movies",routes.UpdateMovieHandler).Methods("PUT")
    r.HandleFunc("/movies/{id}",routes.DeleteMovieHandler).Methods("DELETE")

	http.Handle("/static/", http.StripPrefix("/static/",routes.FileServer))
    http.HandleFunc("/ping", routes.PingHandler)
    http.Handle("/api/", r)

    port := os.Getenv("PORT")
    if len(port)==0{
        port = "5000"
    }
    log.Println("Server started at port:", port)
    http.ListenAndServe(fmt.Sprintf(":%s",port), nil)
}
