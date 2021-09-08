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
