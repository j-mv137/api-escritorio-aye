package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	router.HandleFunc("/api/orders/{state}", makeHTTPHandlerFunc(s.GetOrderByState))
	router.HandleFunc("/api/orders/add", makeHTTPHandlerFunc(s.PostOrder))

	return nil
}

func (s *APIServer) GetOrderByState(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (s *APIServer) PostOrder(w http.ResponseWriter, r *http.Request) error {
	order := new(Order)
	err := json.NewDecoder(r.Body).Decode(order)

	if err != nil {
		return err
	}

	err = s.Store.AddOrder(order)

	if err != nil {
		return err
	}

	w.WriteHeader(204)
	return nil
}

func makeHTTPHandlerFunc(f APIHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
