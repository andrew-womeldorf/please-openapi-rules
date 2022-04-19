package main

import (
	"examples/validate"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	SpecController := validate.NewSpecApiController()

	OverwrittenApiService := validate.NewOverwrittenApiService()
	DefaultApiController := validate.NewDefaultApiController(OverwrittenApiService)

	router := validate.NewRouter(SpecController, DefaultApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
