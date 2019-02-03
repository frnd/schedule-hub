package forms

//AssignmentForm ...
type AssignmentForm struct {
	EmployeeID int64 `form:"employeeId" json:"employeeId" binding:"required,max=200"`
	ProjectID  int64 `form:"projectId" json:"projectId" binding:"required,max=200"`
	Pct        int   `form:"pct" json:"pct" binding:"required,max=200"`
	StartDate  int64 `form:"startDate" json:"startDate" binding:""`
	EndDate    int64 `form:"endDate" json:"endDate" binding:""`
	Date       int64 `form:"date" json:"date" binding:""`
}

type AssignmentUpdateForm struct {
	Pct        int   `form:"pct" json:"pct" binding:"required,max=200"`
}
