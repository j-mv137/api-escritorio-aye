package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	router.HandleFunc("/api/revisions/{id}", makeHTTPHandlerFunc(s.GetRevisionByID))
	router.HandleFunc("/api/revisions/add", makeHTTPHandlerFunc(s.AddRevision))

	return nil
}

func (s *APIServer) GetRevisionByID(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) AddRevision(w http.ResponseWriter, r *http.Request) error {
	revision := new(Revision)
	err := json.NewDecoder(r.Body).Decode(revision)

	if err != nil {
		return err
	}

	err = s.Store.AddRevision(revision)

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
