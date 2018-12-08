package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/frnd/schedule-hub/controllers"
	"github.com/frnd/schedule-hub/db"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()

	// TODO move redis host and port to an env variable.
	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("schedule-hub-session", store))

	r.Use(CORSMiddleware())

	db.Init()

	v1 := r.Group("/v1")
	{
		/*** START USER ***/
		user := new(controllers.UserController)

		v1.POST("/user/signin", user.Signin)
		v1.POST("/user/signup", user.Signup)
		v1.GET("/user/signout", user.Signout)

		/*** START Article ***/
		article := new(controllers.ArticleController)

		v1.POST("/article", article.Create)
		v1.GET("/articles", article.All)
		v1.GET("/article/:id", article.One)
		v1.PUT("/article/:id", article.Update)
		v1.DELETE("/article/:id", article.Delete)

		/*** START Project ***/
		project := new(controllers.ProjectController)

		v1.POST("/project", project.Create)
		v1.GET("/projects", project.All)
		v1.GET("/project/:id", project.One)
		v1.PUT("/project/:id", project.Update)
		v1.DELETE("/project/:id", project.Delete)

		/*** START Employee ***/
		employee := new(controllers.EmployeeController)

		v1.POST("/employee", employee.Create)
		v1.GET("/employees", employee.All)
		v1.GET("/employee/:id", employee.One)
		v1.PUT("/employee/:id", employee.Update)
		v1.DELETE("/employee/:id", employee.Delete)

		/*** START Assignment ***/
		assignment := new(controllers.AssignmentController)

		v1.POST("/assignment", assignment.Create)
		v1.GET("/assignment/employee/:employeeId", assignment.All)
		v1.GET("/assignment/project/:projectId", assignment.All)
		v1.GET("/assignment/employee/:employeeId/project/:projectId", assignment.All)
		v1.GET("/assignment/employee/:employeeId/project/:projectId/date/:date", assignment.One)
		v1.PUT("/assignment/employee/:employeeId/project/:projectId/date/:date", assignment.Update)
		v1.DELETE("/assignment/employee/:employeeId/project/:projectId/date/:date", assignment.Delete)
	}

	r.LoadHTMLGlob("./public/html/*")

	r.Static("/public", "./public")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"goVersion":             runtime.Version(),
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	r.Run()
}
