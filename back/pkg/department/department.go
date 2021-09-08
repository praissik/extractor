package department

import (
	"extractor/back/pkg/db"
)

type Department struct {
	ID           int32  `json:"ID"`
	Lp           int32  `json:"lp"`
	DepartmentID int32  `json:"departmentID"`
	Name         string `json:"name"`
}

var Departments map[int32]Department

func Init() {
	Departments = LoadAll()
}

func LoadAll() map[int32]Department {
	var p Department
	departments := make(map[int32]Department)

	rows := db.InvokeQuery(`SELECT id, id_department, name FROM departments ORDER BY name`)

	var i int32
	i = 1
	for rows.Next() {
		rows.Scan(&p.ID, &p.DepartmentID, &p.Name)
		p.Lp = i
		departments[p.DepartmentID] = p
		i = i + 1
	}

	return departments
}

func Get() map[int32]Department {
	return Departments
}
