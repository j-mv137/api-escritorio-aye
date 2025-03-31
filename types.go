package main

import (
	"net/http"

	"cloud.google.com/go/civil"
)

type Revision struct {
	Id          int            `json:"id"`
	Tipo        string         `json:"tipo"`
	Fecha       civil.DateTime `json:"fecha"`
	Nombre      string         `json:"nombre"`
	Descripcion string         `json:"descripcion"`
	Domicilio   string         `json:"domicilio"`
	Telefono    string         `json:"telefono"`
}

type APIServer struct {
	ListenAdrr string
	Store      Storage
}

func NewAPIServer(listenAdrr string, store Storage) *APIServer {
	return &APIServer{
		ListenAdrr: listenAdrr,
		Store:      store,
	}
}

type APIHandlerFunc = func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string `json:"error"`
}
