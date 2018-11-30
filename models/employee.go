package models

import (
	"errors"
	"time"

	"github.com/frnd/schedule-hub/db"
	"github.com/frnd/schedule-hub/forms"
)

//Employee ...
type Employee struct {
	ID          int64  `db:"id, primarykey, autoincrement" json:"id"`
	Name        string `db:"name" json:"name"`
	UpdatedAt   int64  `db:"updated_at" json:"updated_at"`
	CreatedAt   int64  `db:"created_at" json:"created_at"`
}

//EmployeeModel ...
type EmployeeModel struct{}

//Create ...
func (m EmployeeModel) Create(form forms.EmployeeForm) (projectID int64, err error) {
	getDb := db.GetDB()

	i := "INSERT INTO public.employee(name, updated_at, created_at) " +
		"VALUES($1, $2, $3) RETURNING id"
	r, err := getDb.Exec(i, form.Name, time.Now().Unix(), time.Now().Unix())

	if err != nil {
		return 0, err
	}

	projectID, err = r.LastInsertId()

	return projectID, err
}

//One ...
func (m EmployeeModel) One(id int64) (project Employee, err error) {
	s := "SELECT e.id, e.name, e.updated_at, e.created_at " +
		"FROM public.employee e " +
		"WHERE e.id = $1 " +
		"LIMIT 1"
	err = db.GetDB().SelectOne(&project, s, id)
	return project, err
}

//All ...
func (m EmployeeModel) All() (projects []Employee, err error) {
	s := "SELECT e.id, e.name, e.updated_at, e.created_at " +
		"FROM public.employee e " +
		"ORDER BY e.id DESC"
	_, err = db.GetDB().Select(&projects, s)
	return projects, err
}

//Update ...
func (m EmployeeModel) Update(userID int64, id int64, form forms.EmployeeForm) (err error) {
	_, err = m.One(id)

	if err != nil {
		return errors.New("Employee not found")
	}

	u := "UPDATE public.employee SET name=$1, updated_at=$2 WHERE id=$3"
	_, err = db.GetDB().Exec(u, form.Name, time.Now().Unix(), id)

	return err
}

//Delete ...
func (m EmployeeModel) Delete(userID, id int64) (err error) {
	_, err = m.One(id)

	if err != nil {
		return errors.New("Employee not found")
	}

	d := "DELETE FROM public.employee WHERE id=$1"
	_, err = db.GetDB().Exec(d, id)

	return err
}
