package handler

import (
	"encoding/json"
	"extractor/back/pkg/department"
	"extractor/back/pkg/er"
	"extractor/back/pkg/logger"
	"extractor/back/pkg/util"
	"net/http"
)

func GetDepartments(w http.ResponseWriter, r *http.Request) {
	logger.Console("Get Departments")
	util.SetWriter(w, 200)
	err := json.NewEncoder(w).Encode(department.Get())
	er.Check(err)
}
