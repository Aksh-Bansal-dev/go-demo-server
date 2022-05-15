package routes

import (
	"encoding/json"
    "net/http"
)

type PingResponse struct{
	Message string `json:"message"`
}

var FileServer = http.FileServer(http.Dir("./static"))

func PingHandler(w http.ResponseWriter, r *http.Request){
	if r.Method!="GET"{
		http.Error(w,"Method not supported", 405)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PingResponse{Message: "pong"})
}