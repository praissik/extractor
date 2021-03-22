package handler

import (
	"encoding/json"
	"net/http"
	"pcList/pkg/auth"
	"pcList/pkg/pcInfo"
	"pcList/pkg/util"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	var data auth.Data

	err := json.NewDecoder(r.Body).Decode(&data)
	util.CheckErr(err)

	d, status := auth.CheckAuth(data)

	util.SetWriter(w, status)

	err = json.NewEncoder(w).Encode(d)
	util.CheckErr(err)

	return
}

func GetPcList(w http.ResponseWriter, r *http.Request) {
	var filters pcInfo.Filters
	var err error

	rString, _ := r.URL.Query()["string"]
	if len(rString) > 0 {
		filters.String = rString[0]
	}

	rStatus, _ := r.URL.Query()["status"]
	filters.Status, _ = strconv.Atoi(rStatus[0])

	pcList := pcInfo.GetPcList()

	if filters.String == "" && filters.Status == 0 {

		util.SetWriter(w, 200)

		err = json.NewEncoder(w).Encode(pcList)
		util.CheckErr(err)

		return
	}

	var filteredPcList map[int]pcInfo.PcInfo

	if filters.String != "" && filters.Status != 0 {
		filteredPcList = pcInfo.FilterPcListByStatus(pcList, filters.Status)
		filteredPcList = pcInfo.FilterPcListByString(filteredPcList, filters.String)
	} else if filters.String != "" {
		filteredPcList = pcInfo.FilterPcListByString(pcList, filters.String)
	} else if filters.Status != 0 {
		filteredPcList = pcInfo.FilterPcListByStatus(pcList, filters.Status)
	}

	util.SetWriter(w, 200)

	err = json.NewEncoder(w).Encode(filteredPcList)
	util.CheckErr(err)

	return
}

func AddPcInfo(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	util.CheckErr(err)

	var pI pcInfo.PcInfo
	mapstructure.Decode(data, &pI)

	pcInfo.AddPcInfo(pI)

	return
}

func SetPcInfo(w http.ResponseWriter, r *http.Request) {

	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	util.CheckErr(err)

	var pcInfoToSet pcInfo.PcInfo
	mapstructure.Decode(data, &pcInfoToSet)

	pcList := pcInfo.GetPcList()

	pcInfo.Set(pcList[pcInfoToSet.Lp], pcInfoToSet)

	util.SetWriter(w, 200)
}

func UpdatePcInfo(w http.ResponseWriter, r *http.Request) {

	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	util.CheckErr(err)

	var pcInfoToUpdate pcInfo.PcInfo
	mapstructure.Decode(data, &pcInfoToUpdate)

	pcInfo.Update(pcInfoToUpdate)

}

func DeletePcInfo(w http.ResponseWriter, r *http.Request) {

}
