package main

import (
	"os"
	"student-crud/controllers"
	"student-crud/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.GET("/students", controllers.GetStudents)
	r.POST("/students", controllers.CreateStudent)
	r.GET("/students/:id", controllers.GetStudent)
	r.PUT("/students/:id", controllers.UpdateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
