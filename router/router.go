package router

import (
	"github.com/gorilla/mux"

	"assignment3/cmd"
	"assignment3/mapstore"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	c := cmd.NewCustomerController(mapstore.NewMapStore())
	r.HandleFunc("/customers", c.GetAll).Methods("GET")
	r.HandleFunc("/customers/{id}", c.GetByID).Methods("GET")
	r.HandleFunc("/customers", c.Post).Methods("POST")
	r.HandleFunc("/customers/{id}", c.Put).Methods("PUT")
	r.HandleFunc("/customers/{id}", c.Delete).Methods("DELETE")
	return r
}
