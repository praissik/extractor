package department

import (
	"extractor/back/pkg/db"
	"extractor/back/pkg/er"
	"extractor/back/pkg/report"
	"strings"
)

type Department struct {
	ID           int32  `json:"id"`
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

	rows := db.InvokeQuery(`SELECT id, id_department, name FROM departments`)

	for rows.Next() {
		rows.Scan(&p.ID, &p.DepartmentID, &p.Name)

		departments[p.DepartmentID] = p
	}

	return departments
}

func Get() map[int32]Department {
	return Departments
}

func AddDepartment(data Department) int {

	usedDepartmentName := false

	for _, p := range Get() {
		if strings.ToLower(p.Name) == strings.ToLower(data.Name) {
			usedDepartmentName = true
			break
		}
	}

	if !usedDepartmentName {
		newDepID := db.InvokeQueryRow("SELECT MAX(id_department) FROM departments")
		newDepID.Scan(&data.DepartmentID)

		data.DepartmentID += 1

		result := db.InvokeExec("INSERT INTO departments (id_department, name) VALUES (@p1, @p2)", data.DepartmentID, data.Name)
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

func SetDepartment(data Department) int {

	usedDepartmentName := false

	for _, p := range Get() {
		if p.ID != data.ID && strings.ToLower(p.Name) == strings.ToLower(data.Name) {
			usedDepartmentName = true
			break
		}
	}

	if !usedDepartmentName {
		result := db.InvokeExec("UPDATE departments SET name = @p1 WHERE id_department = @p2", data.Name, data.DepartmentID)
		rA, _ := result.RowsAffected()

		if rA > 0 {
			p := Departments[data.DepartmentID]
			p.Name = data.Name

			Departments[data.DepartmentID] = p
			return 200
		}
		return 400
	} else {
		return 401
	}
}

func DeleteDepartment(data Department) int {

	usedDepartment := false
	for _, r := range report.Get() {
		if r.DepartmentID == data.DepartmentID {
			usedDepartment = true
			break
		}
	}

	if !usedDepartment {
		result := db.InvokeExec("DELETE FROM departments WHERE id_department = @p1", data.DepartmentID)
		rA, err := result.RowsAffected()
		er.Check(err)

		if rA > 0 {
			delete(Departments, data.DepartmentID)
			return 200
		}
		return 400
	} else {
		return 401
	}
}
