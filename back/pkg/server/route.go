package server

import (
	"extractor/back/pkg/server/handler"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api/reports", handler.GetReports).Methods("GET")
	r.HandleFunc("/api/parameters", handler.GetParameters).Methods("GET")
	r.HandleFunc("/api/departments", handler.GetDepartments).Methods("GET")
	r.HandleFunc("/api/report/generate", handler.GenerateReport).Methods("POST")

	return r
}
