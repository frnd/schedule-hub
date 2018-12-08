package controllers

import (
	"strconv"
	"time"

	"github.com/frnd/schedule-hub/forms"
	"github.com/frnd/schedule-hub/models"

	"github.com/gin-gonic/gin"
)

//AssignmentController ...
type AssignmentController struct{}

var assignmentModel = new(models.AssignmentModel)

//Create ...
func (ctrl AssignmentController) Create(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	var assignmentForm forms.AssignmentForm

	if c.BindJSON(&assignmentForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": assignmentForm})
		c.Abort()
		return
	}

	employeeID, err := assignmentModel.Create(assignmentForm)

	if err != nil {
		c.JSON(406, gin.H{"message": "Assignment could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Assignment created", "id": employeeID})
}

//All ...
func (ctrl AssignmentController) All(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	employeeId, err1 := strconv.ParseInt(c.Param("employeeId"), 10, 64)
	projectId, err2 := strconv.ParseInt(c.Param("projectId"), 10, 64)
	fromDate, err3 := strconv.ParseInt(c.Query("from"), 10, 64)
	toDate, err4 := strconv.ParseInt(c.Query("to"), 10, 64)

	if err3 != nil || err4 != nil {
		t := time.Now()
		firstDay := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
		lastDay := firstDay.AddDate(0, 1, 0).Add(time.Nanosecond * -1)
		if err3 != nil {
			fromDate = firstDay.Unix()
		}
		if err4 != nil {
			toDate = lastDay.Unix()
		}
	}

	if err1 != nil {
		employeeId = -1
	}

	if err2 != nil {
		projectId = -1
	}

	if err1 != nil && err2 != nil {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	} else {
		data, err := assignmentModel.All(projectId, employeeId, fromDate, toDate)

		if err != nil {
			c.JSON(406, gin.H{"Message": "Could not get the employees", "error": err.Error()})
			c.Abort()
			return
		}

		c.JSON(200, gin.H{"data": data})
	}
}

//One ...
func (ctrl AssignmentController) One(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	employeeId, err1 := strconv.ParseInt(c.Param("employeeId"), 10, 64)
	projectId, err2 := strconv.ParseInt(c.Param("projectId"), 10, 64)
	date, err3 := strconv.ParseInt(c.Param("date"), 10, 64)

	if err1 == nil && err2 == nil && err3 == nil {

		data, err := assignmentModel.One(employeeId, projectId, date)
		if err != nil {
			c.JSON(404, gin.H{"Message": "Assignment not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"data": data})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}

//Update ...
func (ctrl AssignmentController) Update(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	employeeId, err1 := strconv.ParseInt(c.Param("employeeId"), 10, 64)
	projectId, err2 := strconv.ParseInt(c.Param("projectId"), 10, 64)
	date, err3 := strconv.ParseInt(c.Param("date"), 10, 64)

	if err1 == nil && err2 == nil && err3 == nil {

		var assignmentForm forms.AssignmentUpdateForm

		if c.BindJSON(&assignmentForm) != nil {
			c.JSON(406, gin.H{"message": "Invalid parameters", "form": assignmentForm})
			c.Abort()
			return
		}

		err := assignmentModel.Update(employeeId, projectId, date, assignmentForm)
		if err != nil {
			c.JSON(406, gin.H{"Message": "Assignment could not be updated", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Assignment updated"})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}

//Delete ...
func (ctrl AssignmentController) Delete(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	employeeId, err1 := strconv.ParseInt(c.Param("employeeId"), 10, 64)
	projectId, err2 := strconv.ParseInt(c.Param("projectId"), 10, 64)
	date, err3 := strconv.ParseInt(c.Param("date"), 10, 64)

	if err1 == nil && err2 == nil && err3 == nil {

		err := assignmentModel.Delete(employeeId, projectId, date)
		if err != nil {
			c.JSON(406, gin.H{"Message": "Assignment could not be deleted", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Assignment deleted"})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}

func (ctrl AssignmentController) Find(c *gin.Context) {
	c.JSON(200, gin.H{"Message": "WIP"})
}
