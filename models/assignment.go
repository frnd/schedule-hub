package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/frnd/schedule-hub/db"
	"github.com/frnd/schedule-hub/forms"
)

//Assignment ...
type AssignmentID struct {
	ProjectID  int64 `db:"project_id, primarykey, autoincrement" json:"projectId"`
	EmployeeID int64 `db:"employee_id, primarykey, autoincrement" json:"employeeId"`
	Date       int64 `db:"date, primarykey" json:"date"`
}

type Assignment struct {
	AssignmentID
	Pct       int      `db:"pct" json:"pct"`
	UpdatedAt int64    `db:"updated_at" json:"updated_at"`
	CreatedAt int64    `db:"created_at" json:"created_at"`
	Project   *JSONRaw `db:"project" json:"project"`
	Employee  *JSONRaw `db:"employee" json:"employee"`
}

//AssignmentModel ...
type AssignmentModel struct{}

//Create ...
func (m AssignmentModel) Create(form forms.AssignmentForm) (assignments []AssignmentID, err error) {
	getDb := db.GetDB()

	i := "INSERT INTO public.assignment(project_id, employee_id, date, pct, updated_at, created_at) " +
		"VALUES($1, $2, $3, $4, $5, $6) RETURNING project_id, employee_id, date"

	if err != nil {
		return make([]AssignmentID, 0), err
	}

	var p, e, d int64
	if form.Date != 0 {
		day := time.Unix(form.Date, 0).Truncate(24 * time.Hour).Unix()
		err = getDb.QueryRow(i, form.ProjectID, form.EmployeeID, day, form.Pct, time.Now().Unix(), time.Now().Unix()).Scan(&p, &e, &d)
		if err != nil {
			return make([]AssignmentID, 0), err
		}
		assignments = append(assignments, AssignmentID{ProjectID: p, EmployeeID: e, Date: d})
	} else {
		transaction, err := getDb.Begin()
		date := time.Unix(form.StartDate, 0).Truncate(24 * time.Hour)
		end := time.Unix(form.EndDate, 0).Truncate(24 * time.Hour)
		for date.Before(end) || date.Equal(end) {
			day := date.Unix()
			err = transaction.QueryRow(i, form.ProjectID, form.EmployeeID, day, form.Pct, time.Now().Unix(), time.Now().Unix()).Scan(&p, &e, &d)
			if err != nil {
				transaction.Rollback()
				return make([]AssignmentID, 0), err
			}
			assignments = append(assignments, AssignmentID{ProjectID: p, EmployeeID: e, Date: d})
			date = date.AddDate(0, 0, 1)
		}
		transaction.Commit();
	}

	return assignments, err
}

//One ...
func (m AssignmentModel) One(projectId int64, employeeId int64, date int64) (assignment Assignment, err error) {
	s := "SELECT a.project_id, a.employee_id, a.date, a.pct, a.updated_at, a.created_at, " +
		"json_build_object('id', p.id, 'key', p.key, 'name', p.name, 'description', p.description, 'updated_at', p.updated_at, 'created_at', p.created_at) AS project, " +
		"json_build_object('id', e.id, 'name', e.name, 'updated_at', e.updated_at, 'created_at', e.created_at) AS employee " +
		"FROM public.assignment a " +
		"INNER JOIN public.project p ON a.project_id = p.id " +
		"INNER JOIN public.employee e ON a.employee_id = e.id " +
		"WHERE a.employee_id = $1 " +
		"AND a.project_id = $2 " +
		"AND a.date = $3 " +
		"LIMIT 1"
	println(s)
	err = db.GetDB().SelectOne(&assignment, s, projectId, employeeId, date)
	return assignment, err
}

//All ...
func (m AssignmentModel) All(projectId int64, employeeId int64, start int64, end int64) (assignments []Assignment, err error) {
	s := "SELECT a.project_id, a.employee_id, a.date, a.pct, a.updated_at, a.created_at, " +
		"json_build_object('id', p.id, 'key', p.key, 'name', p.name, 'description', p.description, 'updated_at', p.updated_at, 'created_at', p.created_at) AS project, " +
		"json_build_object('id', e.id, 'name', e.name, 'updated_at', e.updated_at, 'created_at', e.created_at) AS employee " +
		"FROM public.assignment a " +
		"INNER JOIN public.project p ON a.project_id = p.id " +
		"INNER JOIN public.employee e ON a.employee_id = e.id " +
		"WHERE a.date >= $1 " +
		"AND a.date <= $2 "
	var c = 3
	if employeeId != -1 {
		s = s + "AND a.employee_id = $" + strconv.Itoa(c)
		c++
	}
	if projectId != -1 {
		s = s + "AND a.project_id = $" + strconv.Itoa(c)
	}

	if projectId != -1 && employeeId != -1 {
		_, err = db.GetDB().Select(&assignments, s, start, end, employeeId, projectId)
	} else if employeeId != -1 {
		_, err = db.GetDB().Select(&assignments, s, start, end, employeeId)
	} else if projectId != -1 {
		_, err = db.GetDB().Select(&assignments, s, start, end, projectId)
	} else {
		err = errors.New("filter at least by project or employee")
	}
	return assignments, err
}

//Update ...
func (m AssignmentModel) Update(projectId int64, employeeId int64, date int64, form forms.AssignmentUpdateForm) (err error) {
	_, err = m.One(projectId, employeeId, date)

	if err != nil {
		return errors.New("assignment not found")
	}

	u := "UPDATE public.assignment " +
		"SET pct= $1, updated_at=$2 " +
		"WHERE project_id=$3 " +
		"AND employee_id=$4 " +
		"AND date = $5"
	_, err = db.GetDB().Exec(u, form.Pct, time.Now().Unix(), projectId, employeeId, date)

	return err
}

//Delete ...
func (m AssignmentModel) Delete(projectId int64, employeeId int64, date int64) (err error) {
	_, err = m.One(projectId, employeeId, date)

	if err != nil {
		return errors.New("assignment not found")
	}

	d := "DELETE FROM public.assignment a " +
		"WHERE a.employee_id = $1 " +
		"AND a.project_id = $2 " +
		"AND a.date = $3 "
	_, err = db.GetDB().Exec(d, projectId, employeeId, date)

	return err
}
