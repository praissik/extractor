package parameter

import (
	"extractor/back/pkg/db"
	"extractor/back/pkg/report"
	"extractor/back/pkg/util"
	"strings"
)

type Parameter struct {
	ID          int32  `json:"id"`
	ParameterID int32  `json:"parameterID"`
	Name        string `json:"name"`
}

var Parameters map[int32]Parameter

func Init() {
	Parameters = LoadAll()
}

func LoadAll() map[int32]Parameter {
	var p Parameter
	parameters := make(map[int32]Parameter)

	rows := db.InvokeQuery(`SELECT id, id_parameter, name FROM parameters`)

	for rows.Next() {
		rows.Scan(&p.ID, &p.ParameterID, &p.Name)

		parameters[p.ParameterID] = p
	}

	return parameters
}

func Get() map[int32]Parameter {
	return Parameters
}

func AddParameter(data Parameter) int {

	usedParameterName := false

	for _, p := range Get() {
		if strings.ToLower(p.Name) == strings.ToLower(data.Name) {
			usedParameterName = true
			break
		}
	}

	if !usedParameterName {
		newParamID := db.InvokeQueryRow("SELECT MAX(id_parameter) FROM parameters")
		newParamID.Scan(&data.ParameterID)

		data.ParameterID += 1

		result := db.InvokeExec("INSERT INTO parameters (id_parameter, name) VALUES (@p1, @p2)", data.ParameterID, data.Name)
		rA, _ := result.RowsAffected()
		if rA > 0 {
			Init()
			return 200
		}
		return 400
	} else {
		return 401
	}
}

func SetParameter(data Parameter) int {

	usedParameterName := false

	for _, p := range Get() {
		if p.ID != data.ID && strings.ToLower(p.Name) == strings.ToLower(data.Name) {
			usedParameterName = true
			break
		}
	}

	if !usedParameterName {
		result := db.InvokeExec("UPDATE parameters SET name = @p1 WHERE id_parameter = @p2", data.Name, data.ParameterID)
		rA, _ := result.RowsAffected()

		if rA > 0 {
			p := Parameters[data.ParameterID]
			p.Name = data.Name

			Parameters[data.ParameterID] = p
			return 200
		}
		return 400
	} else {
		return 401
	}
}

func DeleteParameter(data Parameter) int {

	usedParameter := false
	for _, r := range report.Get() {
		if util.Contains(r.ParametersID, data.ParameterID) {
			usedParameter = true
			break
		}
	}

	if !usedParameter {
		result := db.InvokeExec("DELETE FROM parameters WHERE id_parameter = @p1", data.ParameterID)
		rA, _ := result.RowsAffected()

		if rA > 0 {
			delete(Parameters, data.ParameterID)
			return 200
		}
		return 400
	} else {
		return 401
	}
}
