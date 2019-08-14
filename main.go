package main

import (
	"log"
	"net/http"

	"github.com/Khamliuk/testsCI/controller"
	"github.com/Khamliuk/testsCI/handler"
	"github.com/Khamliuk/testsCI/mongo"
)

func main() {
	db, err := mongo.New()
	if err != nil {
		log.Fatalf("could not create new db connection: %v", err)
	}
	service := controller.New(db)
	api := handler.New(service)
	log.Fatal(http.ListenAndServe(":8080", api))
}
