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

func AddDepartment(w http.ResponseWriter, r *http.Request) {
	logger.Console("Add Department")
	var data department.Department
	err := json.NewDecoder(r.Body).Decode(&data)
	er.Check(err)

	status := department.AddDepartment(data)
	util.SetWriter(w, status)
}

func SetDepartment(w http.ResponseWriter, r *http.Request) {
	logger.Console("Set Department")
	var data department.Department
	err := json.NewDecoder(r.Body).Decode(&data)
	er.Check(err)

	status := department.SetDepartment(data)
	util.SetWriter(w, status)
}

func DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	logger.Console("Delete Department")
	var data department.Department
	err := json.NewDecoder(r.Body).Decode(&data)
	er.Check(err)

	status := department.DeleteDepartment(data)
	util.SetWriter(w, status)
}
