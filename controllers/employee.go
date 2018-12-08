package controllers

import (
	"strconv"

	"github.com/frnd/schedule-hub/forms"
	"github.com/frnd/schedule-hub/models"

	"github.com/gin-gonic/gin"
)

//EmployeeController ...
type EmployeeController struct{}

var employeeModel = new(models.EmployeeModel)

//Create ...
func (ctrl EmployeeController) Create(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	var employeeForm forms.EmployeeForm

	if c.BindJSON(&employeeForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": employeeForm})
		c.Abort()
		return
	}

	employeeID, err := employeeModel.Create(employeeForm)

	if employeeID > 0 && err != nil {
		c.JSON(406, gin.H{"message": "Employee could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Employee created", "id": employeeID})
}

//All ...
func (ctrl EmployeeController) All(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	data, err := employeeModel.All()

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the employees", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": data})
}

//One ...
func (ctrl EmployeeController) One(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	id := c.Param("id")

	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		data, err := employeeModel.One(id)
		if err != nil {
			c.JSON(404, gin.H{"Message": "Employee not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"data": data})
	}
}

//Update ...
func (ctrl EmployeeController) Update(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		var employeeForm forms.EmployeeForm

		if c.BindJSON(&employeeForm) != nil {
			c.JSON(406, gin.H{"message": "Invalid parameters", "form": employeeForm})
			c.Abort()
			return
		}

		err := employeeModel.Update(userID, id, employeeForm)
		if err != nil {
			c.JSON(406, gin.H{"Message": "Employee could not be updated", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Employee updated"})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter", "error": err.Error()})
	}
}

//Delete ...
func (ctrl EmployeeController) Delete(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		err := employeeModel.Delete(userID, id)
		if err != nil {
			c.JSON(406, gin.H{"Message": "Employee could not be deleted", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Employee deleted"})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}
