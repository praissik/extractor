package util

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Data struct {
	Queries []Query `json:"queries"`
}

type Query struct {
	Name   string `json:"name"`
	Query  string `json:"query"`
	Needed []int  `json:"needed"`
}

func CheckErr(err error) {
	if err != nil {
		_, function, line, _ := runtime.Caller(1)
		fmt.Printf("[error] %s:%d %v \n", function, line, err)
	}
}

func SetWriter(writer http.ResponseWriter, statusCode int) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.WriteHeader(statusCode)
}

func ConsoleLog(message string) {
	time := time.Now().Format("2006-01-02 15:04:05")
	message = time + " " + message
	fmt.Println(message)
}

func GetReportsData() Data {
	var dataJson Data
	var data Data
	var query Query

	jsonQueriesPath := "jsons/queries.json"
	jsonQueries, _ := ioutil.ReadFile(jsonQueriesPath)

	err := json.Unmarshal([]byte(jsonQueries), &dataJson)
	if err != nil {
		panic(err)
	}

	for _, q := range dataJson.Queries {
		query.Name = q.Name
		query.Needed = q.Needed
		data.Queries = append(data.Queries, query)
	}

	return data
}

func SetCellValues(f *excelize.File, rows *sql.Rows) error {
	var x int
	var y int = 2

	columns, _ := rows.Columns()

	var values = make([]interface{}, len(columns))

	for i, _ := range values {
		var ii interface{}
		values[i] = &ii
	}

	var cellName string
	rows.Next()
	err := setColumnHeaders(f, rows, columns, values)

	if err != nil {
		return err
	}

	for {
		err := rows.Scan(values...)
		CheckErr(err)

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
	return nil
}

//func setColumnHeaders(f *excelize.File, rows *sql.Rows) {
func setColumnHeaders(f *excelize.File, rows *sql.Rows, columns []string, values []interface{}) error {
	var cellName string

	err := rows.Scan(values...)
	if err != nil {
		return err
	}

	for i, columnName := range columns {
		cellName, _ = excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue("Raport", cellName, columnName)
	}
	return nil
}

func FetchQueryData(id int) (string, string) {
	var name string
	var query string
	var data Data

	jsonQueriesPath := "jsons/queries.json"
	jsonQueries, _ := ioutil.ReadFile(jsonQueriesPath)

	err := json.Unmarshal([]byte(jsonQueries), &data)
	CheckErr(err)

	name = data.Queries[id].Name
	query = data.Queries[id].Query

	return name, query
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
