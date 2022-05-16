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
	routes.Routes(r)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
	}
	log.Println("Server started at port:", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
