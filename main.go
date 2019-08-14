package main

import (
	"log"
	"net/http"

	"github.com/rockspoon/testsCI/controller"
	"github.com/rockspoon/testsCI/handler"
	"github.com/rockspoon/testsCI/mongo"
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
