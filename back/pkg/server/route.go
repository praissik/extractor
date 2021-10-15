package server

import (
	"extractor/back/pkg/server/handler"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api/data", handler.GetData).Methods("GET")

	r.HandleFunc("/api/reports", handler.GetReports).Methods("GET")
	r.HandleFunc("/api/report", handler.SetReport).Methods("PUT")
	r.HandleFunc("/api/report", handler.AddReport).Methods("POST")
	r.HandleFunc("/api/report", handler.DeleteReport).Methods("DELETE")
	r.HandleFunc("/api/report/generate", handler.GenerateReport).Methods("POST")

	r.HandleFunc("/api/parameters", handler.GetParameters).Methods("GET")
	r.HandleFunc("/api/parameter", handler.SetParameter).Methods("PUT")
	r.HandleFunc("/api/parameter", handler.AddParameter).Methods("POST")
	r.HandleFunc("/api/parameter", handler.DeleteParameter).Methods("DELETE")

	r.HandleFunc("/api/departments", handler.GetDepartments).Methods("GET")
	r.HandleFunc("/api/department", handler.SetDepartment).Methods("PUT")
	r.HandleFunc("/api/department", handler.AddDepartment).Methods("POST")
	r.HandleFunc("/api/department", handler.DeleteDepartment).Methods("DELETE")

	return r
}
