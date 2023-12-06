package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"

	"assignment3/domain"
)

type CustomerController struct {
	store domain.CustomerStore
}

var (
	errMissingFields = errors.New("Missing required fields")
)

func NewCustomerController(store domain.CustomerStore) *CustomerController {
	return &CustomerController{
		store: store,
	}
}

func (cc *CustomerController) Post(w http.ResponseWriter, r *http.Request) {
	var c domain.Customer
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := cc.store.Create(c); err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	slog.Info("New customer added")
	w.WriteHeader(http.StatusCreated)
}

func (cc *CustomerController) Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var c domain.Customer
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := cc.store.Update(id, c); err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	m := fmt.Sprintf("Customer with id %s updated\n", vars["id"])
	slog.Info(m)
	w.WriteHeader(http.StatusNoContent)
}

func (cc *CustomerController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := cc.store.Delete(vars["id"]); err != nil {
		slog.Error(err.Error())
		return
	}
	m := fmt.Sprintf("Customer with id %s deleted\n", vars["id"])
	slog.Info(m)
}

func (cc *CustomerController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if c, err := cc.store.GetByID(vars["id"]); err != nil {
		slog.Error(err.Error())
		return
	} else {
		c, err := json.Marshal(c)
		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(c)
	}
}

func (cc *CustomerController) GetAll(w http.ResponseWriter, r *http.Request) {
	if customers, err := cc.store.GetAll(); err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		customers, err := json.Marshal(customers)
		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(customers)
	}
}
