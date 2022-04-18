package main

import (
	"examples/validate"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	OverwrittenApiService := validate.NewOverwrittenApiService()
	DefaultApiController := validate.NewDefaultApiController(OverwrittenApiService)

	router := validate.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
