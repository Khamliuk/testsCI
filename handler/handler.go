package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Khamliuk/testsCI/model"
	"github.com/gorilla/mux"
)

type Service interface {
	Create(ctx context.Context, req model.Person) (*model.Person, error)
	List(ctx context.Context) ([]model.Person, error)
	Update(ctx context.Context, req model.Person) error
	Delete(ctx context.Context, id string) error
}
type API struct {
	*mux.Router
	service Service
}

func New(service Service) API {
	router := mux.NewRouter().PathPrefix("/person").Subrouter()
	api := API{
		router,
		service,
	}
	router.Path("").
		Methods(http.MethodGet).
		HandlerFunc(api.persons)
	router.Path("").
		Methods(http.MethodPost).
		HandlerFunc(api.createPerson)
	router.Path("").
		Methods(http.MethodPut).
		HandlerFunc(api.updatePerson)
	router.Path("").
		Methods(http.MethodDelete).
		HandlerFunc(api.deletePerson)
	return api
}

func (a API) persons(w http.ResponseWriter, r *http.Request) {
	resp, err := a.service.List(r.Context())
	if err != nil {
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			log.Println(err)
		}
		return
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(err)
	}
}

func (a API) createPerson(w http.ResponseWriter, r *http.Request) {
	var req model.Person
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			log.Println(err)
		}
		return
	}
	resp, err := a.service.Create(r.Context(), req)
	if err != nil {
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			log.Println(err)
		}
		return
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(err)
	}
}

func (a API) updatePerson(w http.ResponseWriter, r *http.Request) {
	var req model.Person
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			log.Println(err)
		}
		return
	}
	err = a.service.Update(r.Context(), req)
	if err != nil {
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			log.Println(err)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (a API) deletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := a.service.Delete(r.Context(), vars["id"])
	if err != nil {
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			log.Println(err)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
