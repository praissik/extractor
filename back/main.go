package main

import (
	"extractor/back/pkg/department"
	"extractor/back/pkg/logger"
	"extractor/back/pkg/parameter"
	"extractor/back/pkg/report"
	"extractor/back/pkg/server"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	logger.Console("Initializing data")
	report.Init()
	parameter.Init()
	department.Init()
	logger.Console("Server has been started")

	r := server.Router()

	http.Handle("/", r)

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	log.Fatalln(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
