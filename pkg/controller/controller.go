package controller

import (
	"encoding/json"
	"net/http"

	"example.com/go-demo-server/pkg/db"
	"github.com/gorilla/mux"
)

type PingResponse struct {
	Message string `json:"message"`
}

var FileServer = http.FileServer(http.Dir("./static"))

func PingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not supported", 405)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PingResponse{Message: "pong"})
}

func GetMoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.GetAll())
}

func GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	res, err := db.Get(vars["id"])
	if err != nil {
		http.Error(w, "Movie not found", 404)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func AddMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	var body db.Album
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Could not add movie, try again later", 400)
		return
	}
	db.Add(body)
	json.NewEncoder(w).Encode(struct {
		Ok bool `json:"ok"`
	}{Ok: true})
}

func UpdateMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body db.Album
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Could not add movie, try again later", 400)
		return
	}
	db.Update(body.ID, body)
	json.NewEncoder(w).Encode(struct {
		Ok bool `json:"ok"`
	}{Ok: true})
}

func DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	if err := db.Delete(vars["id"]); err != nil {
		http.Error(w, "Movie not found", 404)
		return
	}
	json.NewEncoder(w).Encode(struct {
		Ok bool `json:"ok"`
	}{Ok: true})
}
