package forms

//EmployeeForm ...
type EmployeeForm struct {
	Name    string `form:"name" json:"name" binding:"required,max=200"`
}
