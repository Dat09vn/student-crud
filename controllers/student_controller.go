package controllers

import (
	"net/http"
	"student-crud/database"
	"student-crud/models"

	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusCreated, student)
}

func GetStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func GetStudentByName(c *gin.Context) {
	name := c.Param("name")
	var students []models.Student
	if err := database.DB.Where("name = ?", name).Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(students) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No students found with that name"})
		return
	}
	c.JSON(http.StatusOK, students)
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	var input models.Student
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	student.Name = input.Name
	student.Email = input.Email
	student.Age = input.Age
	database.DB.Save(&student)
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Student{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
}
