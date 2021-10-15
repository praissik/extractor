package handler

import (
	"encoding/json"
	"extractor/back/pkg/er"
	"extractor/back/pkg/logger"
	"extractor/back/pkg/report"
	"extractor/back/pkg/util"
	"net/http"
)

func GetReports(w http.ResponseWriter, r *http.Request) {
	logger.Console("Get Reports")
	util.SetWriter(w, 200)
	err := json.NewEncoder(w).Encode(report.Get())
	er.Check(err)
}

func AddReport(w http.ResponseWriter, r *http.Request) {
	logger.Console("Add Report")
	var data report.Report
	err := json.NewDecoder(r.Body).Decode(&data)
	er.Check(err)

	status := report.AddReport(data)
	util.SetWriter(w, status)
}

func SetReport(w http.ResponseWriter, r *http.Request) {
	logger.Console("Set Report")
	var data report.Report
	err := json.NewDecoder(r.Body).Decode(&data)
	er.Check(err)

	status := report.SetReport(data)
	util.SetWriter(w, status)
}

func DeleteReport(w http.ResponseWriter, r *http.Request) {
	logger.Console("Delete Report")
	var data report.Report
	err := json.NewDecoder(r.Body).Decode(&data)
	er.Check(err)

	status := report.DeleteReport(data)
	util.SetWriter(w, status)
}

func GenerateReport(w http.ResponseWriter, r *http.Request) {
	logger.Console("Generate Report")
	var data report.RequestReportData
	err := json.NewDecoder(r.Body).Decode(&data)
	er.Check(err)

	report := report.GenerateReport(data)

	util.SetWriter(w, 200)

	err = json.NewEncoder(w).Encode(report)
	er.Check(err)
}
