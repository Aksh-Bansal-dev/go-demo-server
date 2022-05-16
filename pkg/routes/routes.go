package routes

import (
	"net/http"

	"example.com/go-demo-server/pkg/controller"
	"github.com/gorilla/mux"
)

type PingResponse struct {
	Message string `json:"message"`
}

var FileServer = http.FileServer(http.Dir("./static"))

func Routes(r *mux.Router) {

	r.HandleFunc("/movies", controller.GetMoviesHandler).Methods("GET")
	r.HandleFunc("/movies/{id}", controller.GetMovieHandler).Methods("GET")
	r.HandleFunc("/movies", controller.AddMovieHandler).Methods("POST")
	r.HandleFunc("/movies", controller.UpdateMovieHandler).Methods("PUT")
	r.HandleFunc("/movies/{id}", controller.DeleteMovieHandler).Methods("DELETE")

	http.Handle("/static/", http.StripPrefix("/static/", FileServer))
	http.HandleFunc("/ping", controller.PingHandler)
	http.Handle("/api/", r)
}
