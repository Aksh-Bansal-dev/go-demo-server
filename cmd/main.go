package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "example.com/go-server/pkg/routes"
)

func main() {
	http.Handle("/", routes.FileServer)
    http.HandleFunc("/ping", routes.PingHandler)

    port := os.Getenv("PORT")
    if len(port)==0{
        port = "5000"
    }
    log.Println("Server started at port:", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s",port), nil))
}
