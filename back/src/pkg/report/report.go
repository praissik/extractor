package report

import (
	"database/sql"
	"encoding/json"
	"extractor/pkg/db"
	"extractor/pkg/util"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/lithammer/shortuuid"
)

type Data struct {
	IsOk       int    `json:"isOk"`
	ReportName string `json:"reportName"`
	FileName   string `json:"fileName"`
}

func GetReports(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Get names of reports")

	data := util.GetReportsData()

	util.SetWriter(w, 200)
	//fmt.Println("data:", data)
	err := json.NewEncoder(w).Encode(data)
	util.CheckErr(err)
}

func GenerateReport(w http.ResponseWriter, r *http.Request) {

	var data Data

	idFV := r.FormValue("id")
	id, _ := strconv.Atoi(idFV)

	name, query := util.FetchQueryData(id)

	//fmt.Println("Generate report")
	akronim := r.FormValue("akronim")
	kodTowaru := r.FormValue("kodTowaru")
	symbolTowaru := r.FormValue("symbolTowaru")
	rokOS := r.FormValue("rokOS")

	db := db.Connect()
	rows, err := db.Query(query, sql.Named("Akronim", akronim),
		sql.Named("KodTowaru", kodTowaru),
		sql.Named("SymbolTowaru", symbolTowaru),
		sql.Named("RokOS", rokOS))
	defer db.Close()

	if err != nil {
		data.IsOk = -1
		fmt.Println(err)
		util.CheckErr(err)
	} else {

		f := excelize.NewFile()
		f.SetSheetName("Sheet1", "Raport")

		err := util.SetCellValues(f, rows)

		if err != nil {
			data.IsOk = 0
			fmt.Println(err)
		} else {

			uuid := shortuuid.New()

			data.IsOk = 1
			data.ReportName = name + ".xlsx"
			data.FileName = uuid + ".xlsx"
			url := "./front/static/" + data.FileName
			/*
				fmt.Println(uuid)
				fmt.Println(data.ReportName)
				fmt.Println(url)
			*/
			err := f.SaveAs(url)
			util.CheckErr(err)

			go rm(url)
		}
	}

	util.SetWriter(w, 200)

	err = json.NewEncoder(w).Encode(data)
	util.CheckErr(err)
}

func rm(path string) {
	time.Sleep(10000 * time.Millisecond)
	os.Remove(path)
}
