package main

import (
	"api-services/routers"
	"log"
	"net/http"
	"os"
)

func main() {
	port := ""
	p := os.Getenv("PORT")
	if p == "" {
		port = "3021"
	}
	r := routers.SetupRouter()
	log.Printf("Running Service on %s port", port)
	log.Println(http.ListenAndServe(":"+port, r))
}
