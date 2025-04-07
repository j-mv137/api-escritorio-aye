package main

import (
	"net/http"

	"cloud.google.com/go/civil"
)

type Order struct {
	Id          int            `json:"id"`
	Tipo        string         `json:"tipo"`
	Fecha       civil.DateTime `json:"fecha"`
	Nombre      string         `json:"nombre"`
	Domicilio   string         `json:"domicilio"`
	Telefono    string         `json:"telefono"`
	Estado      string         `json:"estado"`
	Descripcion string         `json:"descripcion"`
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
