package data

import (
	"extractor/back/pkg/department"
	"extractor/back/pkg/logger"
	"extractor/back/pkg/parameter"
	"extractor/back/pkg/report"
)

type Data struct {
	Reports     map[int32]report.Report         `json:"reports"`
	Parameters  map[int32]parameter.Parameter   `json:"parameters"`
	Departments map[int32]department.Department `json:"departments"`
}

func Get() Data {
	var data Data

	data.Reports = report.Get()
	data.Parameters = parameter.Get()
	data.Departments = department.Get()

	go logger.DB(logger.GET, logger.DATA, 0, true)

	return data
}
