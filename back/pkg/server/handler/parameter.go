package handler

import (
	"encoding/json"
	"extractor/back/pkg/er"
	"extractor/back/pkg/logger"
	"extractor/back/pkg/parameter"
	"extractor/back/pkg/util"
	"net/http"
)

func GetParameters(w http.ResponseWriter, r *http.Request) {
	logger.Console("Get Parameters")
	util.SetWriter(w, 200)
	err := json.NewEncoder(w).Encode(parameter.Get())
	er.Check(err)
}

func AddParameter(w http.ResponseWriter, r *http.Request) {
	logger.Console("Add Parameter")
	var data parameter.Parameter
	err := json.NewDecoder(r.Body).Decode(&data)
	er.Check(err)

	status := parameter.AddParameter(data)
	util.SetWriter(w, status)
}

func SetParameter(w http.ResponseWriter, r *http.Request) {
	logger.Console("Set Parameter")
	var data parameter.Parameter
	err := json.NewDecoder(r.Body).Decode(&data)
	er.Check(err)

	status := parameter.SetParameter(data)
	util.SetWriter(w, status)
}

func DeleteParameter(w http.ResponseWriter, r *http.Request) {
	logger.Console("Delete Parameter")
	var data parameter.Parameter
	err := json.NewDecoder(r.Body).Decode(&data)
	er.Check(err)

	status := parameter.DeleteParameter(data)
	util.SetWriter(w, status)
}
