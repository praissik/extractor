package parameter

import (
	"extractor/back/pkg/db"
)

type Parameter struct {
	ID          int32  `json:"ID"`
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
