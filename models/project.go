package models

import (
	"errors"
	"time"

	"github.com/frnd/schedule-hub/db"
	"github.com/frnd/schedule-hub/forms"
)

//Project ...
type Project struct {
	ID          int64  `db:"id, primarykey, autoincrement" json:"id"`
	Key         string `db:"key" json:"key"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	UpdatedAt   int64  `db:"updated_at" json:"updated_at"`
	CreatedAt   int64  `db:"created_at" json:"created_at"`
}

//ProjectModel ...
type ProjectModel struct{}

//Create ...
func (m ProjectModel) Create(form forms.ProjectForm) (projectID int64, err error) {
	getDb := db.GetDB()

	i := "INSERT INTO public.project(key, name, description, updated_at, created_at) " +
		"VALUES($1, $2, $3, $4, $5) RETURNING id"
	// r, err := getDb.Exec(i, form.Key, form.Name, form.Description, time.Now().Unix(), time.Now().Unix())
	getDb.QueryRow(i, form.Key, form.Name, form.Description, time.Now().Unix(), time.Now().Unix()).Scan(&projectID)

	if err != nil {
		return 0, err
	}

	return projectID, err
}

//One ...
func (m ProjectModel) One(id int64) (project Project, err error) {
	s := "SELECT p.id, p.key, p.name, p.description, p.updated_at, p.created_at " +
		"FROM public.project p " +
		"WHERE p.id = $1 " +
		"LIMIT 1"
	err = db.GetDB().SelectOne(&project, s, id)
	return project, err
}

//All ...
func (m ProjectModel) All() (projects []Project, err error) {
	s := "SELECT p.id, p.key, p.name, p.description, p.updated_at, p.created_at " +
		"FROM public.project p " +
		"ORDER BY p.id DESC"
	_, err = db.GetDB().Select(&projects, s)
	return projects, err
}

//Update ...
func (m ProjectModel) Update(userID int64, id int64, form forms.ProjectForm) (err error) {
	_, err = m.One(id)

	if err != nil {
		return errors.New("Project not found")
	}

	u := "UPDATE public.project SET key=$1, name= $2, description=$3, updated_at=$4 WHERE id=$5"
	_, err = db.GetDB().Exec(u, form.Key, form.Name, form.Description, time.Now().Unix(), id)

	return err
}

//Delete ...
func (m ProjectModel) Delete(userID, id int64) (err error) {
	_, err = m.One(id)

	if err != nil {
		return errors.New("Project not found")
	}

	d := "DELETE FROM public.project WHERE id=$1"
	_, err = db.GetDB().Exec(d, id)

	return err
}
