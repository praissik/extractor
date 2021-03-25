package main

import (
	"extractor/back/pkg/report"
	"extractor/back/pkg/util"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	util.ConsoleLog("Server has been started")

	r := mux.NewRouter().StrictSlash(true)

	//HOME
	r.HandleFunc("/reports", report.GetReports).Methods("GET")
	r.HandleFunc("/report/generate", report.GenerateReport).Methods("GET")

	http.Handle("/", r)

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	log.Fatalln(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
	//log.Fatalln(http.ListenAndServe(":8080", r))
}
