package report

import (
	"extractor/back/pkg/db"
	"extractor/back/pkg/er"
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
	ID           int32   `json:"id"`
	DepartmentID int32   `json:"departmentID"`
	ReportName   string  `json:"reportName"`
	SourceName   string  `json:"sourceName"`
	ParametersID []int32 `json:"parametersID"`
	Active       bool    `json:"active"`
}

type ExtractReportData struct {
	Status     bool   `json:"status"`
	ReportName string `json:"reportName"`
	FileName   string `json:"fileName"`
}

type RequestReportData struct {
	ID      int32  `json:"id"`
	Akronim string `json:"akronim"`
	Kod     string `json:"kod"`
	Symbol  string `json:"symbol"`
	RokOS   string `json:"rokOS"`
}

var Reports map[int32]Report

func Init() {
	Reports = LoadAll()
	SetReports()
}

func LoadAll() map[int32]Report {

	var r Report
	reports := make(map[int32]Report)

	rows := db.InvokeQuery(`SELECT id, department_id, active, report_name, source_name FROM reports`)

	for rows.Next() {
		rows.Scan(&r.ID, &r.DepartmentID, &r.Active, &r.ReportName, &r.SourceName)
		reports[r.ID] = r
	}
	return reports
}

func Get() map[int32]Report {
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

func AddReport(data Report) int {

	usedReportName := false

	for _, r := range Get() {
		if strings.ToLower(r.ReportName) == strings.ToLower(data.ReportName) {
			usedReportName = true
			break
		}
	}

	if !usedReportName {

		if data.DepartmentID == 0 {
			logger.Console("Department not selected")
			return 402
		}

		result := db.InvokeExec("INSERT INTO reports (department_id, active, report_name, source_name) VALUES (@p1, 1, @p2, @p3)", data.DepartmentID, data.ReportName, data.SourceName)
		rAR, _ := result.RowsAffected()

		newReportID := db.InvokeQueryRow("SELECT MAX(id) FROM reports")
		newReportID.Scan(&data.ID)

		if rAR > 0 {
			rARp := true
			for _, parameterID := range data.ParametersID {
				result = db.InvokeExec("INSERT INTO reportsParameters (report_id, parameter_id) VALUES (@p1, @p2)", data.ID, parameterID)
				temp, _ := result.RowsAffected()
				if temp == 0 {
					rARp = false
					break
				}
			}
			if rARp {
				Init()
				return 200
			} else {
				for _, parameterID := range data.ParametersID {
					_ = db.InvokeExec("DELETE FROM reportsParameters WHERE report_id = @p1 AND parameter_id = @p2)", data.ID, parameterID)
				}
			}
		}
		return 400
	} else {
		return 401
	}
}

func SetReport(data Report) int {

	usedReportName := false

	for _, p := range Get() {
		if p.ID != data.ID && strings.ToLower(p.ReportName) == strings.ToLower(data.ReportName) {
			usedReportName = true
			break
		}
	}

	if !usedReportName {
		result := db.InvokeExec("UPDATE reports SET report_name = @p1, source_name = @p2, department_id = @p3, active = @p4 WHERE id = @p5", data.ReportName, data.SourceName, data.DepartmentID, data.Active, data.ID)
		rA, _ := result.RowsAffected()

		if rA > 0 {
			r := Reports[data.ID]

			for _, currentParameter := range r.ParametersID {
				if !util.Contains(data.ParametersID, currentParameter) {
					db.InvokeExec("DELETE FROM reportsParameters WHERE report_id = @p1 AND parameter_id = @p2", r.ID, currentParameter)
				}
			}

			for _, p := range data.ParametersID {
				if !util.Contains(r.ParametersID, p) {
					db.InvokeExec("INSERT INTO reportsParameters (report_id, parameter_id) VALUES (@p1, @p2)", r.ID, p)
				}
			}
			Init()
			return 200
		}
		return 400
	} else {
		return 401
	}
}

func DeleteReport(data Report) int {

	result := db.InvokeExec("DELETE FROM reportsParameters WHERE report_id = @p1", data.ID)
	rA, err := result.RowsAffected()
	er.Check(err)

	result = db.InvokeExec("DELETE FROM reports WHERE id = @p1", data.ID)
	rA, err = result.RowsAffected()
	er.Check(err)

	if rA > 0 {
		delete(Reports, data.ID)
		return 200
	}
	return 400
}

func GenerateReport(requestReportData RequestReportData) ExtractReportData {
	var extractReportData ExtractReportData

	numberOfParameters, reportName, sourceName := FetchReportData(requestReportData.ID)

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
		extractReportData.FileName = "report_temp/" + uuid + ".xlsx"

		url := "../front/static/" + extractReportData.FileName

		util.SetCellValues(f, rows)

		f.SaveAs(url)

		go rm(url)
	}

	go logger.DB(logger.GENERATE, logger.REPORT, requestReportData.ID, extractReportData.Status)

	return extractReportData
}

func FetchReportData(ID int32) (int32, string, string) {
	var numberOfParameters int32
	var reportName string
	var sourceName string

	row := db.InvokeQueryRow(`SELECT (SELECT COUNT(id) FROM reportsParameters WHERE report_id = @p1), report_name, source_name FROM reports WHERE id = @p1`, ID)

	row.Scan(&numberOfParameters, &reportName, &sourceName)

	return numberOfParameters, reportName, sourceName
}

// If numberOfParameters == 0, the PrepareQuery function return the view,
// otherwise the PrepareQuery function will return the procedure
func PrepareQuery(numberOfParameters int32, sourceName string, requestReportData RequestReportData) string {
	if numberOfParameters == 0 {
		return `SELECT * FROM ` + sourceName
	} else {
		var (
			parameterID              int32
			procedureQueryParameters []string
		)

		rows := db.InvokeQuery(`SELECT parameter_id FROM reportsParameters WHERE report_id = @p1`, requestReportData.ID)

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
