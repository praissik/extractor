package report

import (
	"extractor/back/pkg/db"
	"extractor/back/pkg/logger"
	"extractor/back/pkg/util"
	"os"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/lithammer/shortuuid"
)

const (
	PARAM_ID_AKRONIM = 1
	PARAM_ID_KOD     = 2
	PARAM_ID_SYMBOL  = 3
	PARAM_ID_ROK_OS  = 4
)

type Report struct {
	ID           int32   `json:"ID"`
	ReportID     int32   `json:"reportID"`
	DepartmentID int32   `json:"departmentID"`
	Name         string  `json:"name"`
	ParametersID []int32 `json:"parametersID"`
}

type ExtractReportData struct {
	Status     bool   `json:"status"`
	ReportName string `json:"reportName"`
	File       string `json:"fileName"`
}

type RequestReportData struct {
	ReportID int32  `json:"reportID"`
	Akronim  string `json:"akronim"`
	Kod      string `json:"kod"`
	Symbol   string `json:"symbol"`
	RokOS    string `json:"rokOS"`
}

var Reports map[int32]Report

func Init() {
	Reports = LoadAll()
	SetReports()
}

func LoadAll() map[int32]Report {

	var r Report
	reports := make(map[int32]Report)

	rows := db.InvokeQuery(`SELECT id, id_report, department_id, report_name FROM reports`)

	for rows.Next() {
		rows.Scan(&r.ID, &r.ReportID, &r.DepartmentID, &r.Name)
		reports[r.ReportID] = r
	}

	return reports
}

func Get() map[int32]Report {
	go logger.DB(logger.GET, logger.REPORT, 0, true)
	return Reports
}

func SetReports() {
	var reportID int32
	var parameterID int32

	rows := db.InvokeQuery(`SELECT report_id, parameter_id FROM reportsParameters`)

	for rows.Next() {
		rows.Scan(&reportID, &parameterID)

		r := Reports[reportID]
		r.ParametersID = append(r.ParametersID, parameterID)
		Reports[reportID] = r
	}
}

func GenerateReport(requestReportData RequestReportData) ExtractReportData {
	var extractReportData ExtractReportData

	numberOfParameters, reportName, sourceName := FetchReportData(requestReportData.ReportID)

	query := PrepareQuery(numberOfParameters, sourceName, requestReportData)

	rows := db.InvokeQuery(query)

	if !rows.Next() {
		extractReportData.Status = false
	} else {
		f := excelize.NewFile()
		f.SetSheetName("Sheet1", "Raport")

		uuid := shortuuid.New()

		extractReportData.Status = true
		extractReportData.ReportName = reportName + ".xlsx"
		extractReportData.File = "report_temp/" + uuid + ".xlsx"

		url := "../front/static/" + extractReportData.File

		util.SetCellValues(f, rows)

		f.SaveAs(url)

		go rm(url)
	}

	go logger.DB(logger.GENERATE, logger.REPORT, requestReportData.ReportID, extractReportData.Status)

	return extractReportData
}

func FetchReportData(ReportID int32) (int32, string, string) {
	var numberOfParameters int32
	var reportName string
	var sourceName string

	row := db.InvokeQueryRow(`SELECT (SELECT COUNT(id) FROM reportsParameters WHERE report_id = @p1), report_name, source_name FROM reports WHERE id_report = @p1`, ReportID)

	row.Scan(&numberOfParameters, &reportName, &sourceName)

	return numberOfParameters, reportName, sourceName
}

// If numberOfParameters == 0, the tarReports query will invoke the view,
// otherwise the tarReports query will invoke the procedure
func PrepareQuery(numberOfParameters int32, sourceName string, requestReportData RequestReportData) string {
	if numberOfParameters == 0 {
		return `SELECT * FROM ` + sourceName
	} else {
		var (
			parameterID              int32
			procedureQueryParameters []string
		)

		rows := db.InvokeQuery(`SELECT parameter_id FROM reportsParameters WHERE report_id = @p1`, requestReportData.ReportID)

		i := 1
		for rows.Next() {
			rows.Scan(&parameterID)
			switch parameterID {
			case PARAM_ID_AKRONIM:
				procedureQueryParameters = append(procedureQueryParameters, ("'" + requestReportData.Akronim + "'"))

			case PARAM_ID_KOD:
				procedureQueryParameters = append(procedureQueryParameters, ("'" + requestReportData.Kod + "'"))

			case PARAM_ID_SYMBOL:
				procedureQueryParameters = append(procedureQueryParameters, ("'" + requestReportData.Symbol + "'"))

			case PARAM_ID_ROK_OS:
				procedureQueryParameters = append(procedureQueryParameters, ("'" + requestReportData.RokOS + "'"))
			}
			i = i + 1
		}

		procedureQuery := `EXEC ` + sourceName + ` ` + strings.Join(procedureQueryParameters, ", ")

		return procedureQuery
	}
}

func rm(path string) {
	time.Sleep(10000 * time.Millisecond)
	os.Remove(path)
}
