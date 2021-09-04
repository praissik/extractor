package util

import (
	"database/sql"
	"extractor/back/pkg/er"
	"net/http"
	"os"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func SetWriter(writer http.ResponseWriter, statusCode int) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.WriteHeader(statusCode)
}

func SetCellValues(f *excelize.File, rows *sql.Rows) {
	var x int
	var y int = 2

	columns, _ := rows.Columns()

	var values = make([]interface{}, len(columns))

	for i, _ := range values {
		var ii interface{}
		values[i] = &ii
	}

	var cellName string
	setColumnHeaders(f, rows, columns, values)

	for {
		err := rows.Scan(values...)
		er.Check(err)

		x = 1
		for i, _ := range columns {
			cellValue := *(values[i].(*interface{}))

			cellName, _ = excelize.CoordinatesToCellName(x, y)
			f.SetCellValue("Raport", cellName, cellValue)
			x++
		}
		y++

		if !rows.Next() {
			break
		}
	}
}

//func setColumnHeaders(f *excelize.File, rows *sql.Rows) {
func setColumnHeaders(f *excelize.File, rows *sql.Rows, columns []string, values []interface{}) {
	var cellName string

	rows.Scan(values...)

	for i, columnName := range columns {
		cellName, _ = excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue("Raport", cellName, columnName)
	}
}

func FileExistsNoLongerThan(filename string, minutes int) bool {
	info, err := os.Stat("./reports/" + filename + ".xlsx")
	if os.IsNotExist(err) {
		return false
	}

	modY, modM, modD := info.ModTime().Date()
	nowY, nowM, nowD := time.Now().Date()

	modH, modMin, _ := info.ModTime().Clock()
	nowH, nowMin, _ := time.Now().Clock()

	//fmt.Println(modY, modM, modD, modH, modMin)
	//fmt.Println(nowY, nowM, nowD, nowH, nowMin)

	if modY == nowY && modM == nowM && modD == nowD {
		if modH+1 == nowH {
			nowMin = nowMin + 60
		}
		if nowMin-minutes < modMin {
			return true
		}
	}
	return false
}
