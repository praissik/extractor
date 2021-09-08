package handler

import (
	"encoding/json"
	"extractor/back/pkg/data"
	"extractor/back/pkg/er"
	"extractor/back/pkg/logger"
	"extractor/back/pkg/util"
	"net/http"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	logger.Console("Get Data")
	util.SetWriter(w, 200)
	err := json.NewEncoder(w).Encode(data.Get())
	er.Check(err)
}
