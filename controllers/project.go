package controllers

import (
	"strconv"

	"github.com/frnd/schedule-hub/forms"
	"github.com/frnd/schedule-hub/models"

	"github.com/gin-gonic/gin"
)

//ProjectController ...
type ProjectController struct{}

var projectModel = new(models.ProjectModel)

//Create ...
func (ctrl ProjectController) Create(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	var projectForm forms.ProjectForm

	if c.BindJSON(&projectForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": projectForm})
		c.Abort()
		return
	}

	projectID, err := projectModel.Create(projectForm)

	if projectID > 0 && err != nil {
		c.JSON(406, gin.H{"message": "Project could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Project created", "id": projectID})
}

//All ...
func (ctrl ProjectController) All(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	data, err := projectModel.All()

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the projects", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": data})
}

//One ...
func (ctrl ProjectController) One(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	id := c.Param("id")

	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		data, err := projectModel.One(id)
		if err != nil {
			c.JSON(404, gin.H{"Message": "Project not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"data": data})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}

//Update ...
func (ctrl ProjectController) Update(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		var projectForm forms.ProjectForm

		if c.BindJSON(&projectForm) != nil {
			c.JSON(406, gin.H{"message": "Invalid parameters", "form": projectForm})
			c.Abort()
			return
		}

		err := projectModel.Update(userID, id, projectForm)
		if err != nil {
			c.JSON(406, gin.H{"Message": "Project could not be updated", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Project updated"})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter", "error": err.Error()})
	}
}

//Delete ...
func (ctrl ProjectController) Delete(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		err := projectModel.Delete(userID, id)
		if err != nil {
			c.JSON(406, gin.H{"Message": "Project could not be deleted", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Project deleted"})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}
