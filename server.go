package main

import (
	"atp-go-rest/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.NewRouter()
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}
