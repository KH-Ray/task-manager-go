package controllers

import (
	"net/http"
	"strings"
	"task-manager-be/db"
	"task-manager-be/models"

	"github.com/gin-gonic/gin"
)

type CreateTaskInput struct {
    Name string `json:"name" binding:"required"`
}

type UpdateTaskInput struct {
    Name      string `json:"name"`
    Completed *bool  `json:"completed"`
}

func FindTasks(c *gin.Context) {
    var tasks []models.Task
    db.DB.Find(&tasks)

    c.JSON(http.StatusOK, tasks)
}

func FindTask(c *gin.Context) {
    var task models.Task

    if err := db.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Task not found!"})
        return
    }

    c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
    var input CreateTaskInput

    if err := c.ShouldBindJSON(&input); err !=  nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if len(input.Name) > 20 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Input exceeds maximum length of 20 characters."})
        return
    }

    task := models.Task{
        Name: strings.Trim(input.Name, " "),
        Completed: false,
    }
    db.DB.Create(&task)

    c.JSON(http.StatusOK, task)
}

func DeleteBook(c *gin.Context) {
    var task models.Task

    if err := db.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Task not found!"})
        return
    }

    db.DB.Delete(&task)

    c.JSON(http.StatusOK, gin.H{"data": true})
}

func UpdateTask(c *gin.Context) {
    var task models.Task
    if err := db.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Task not found!"})
        return
    }

    var input UpdateTaskInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db.DB.Model(&task).Updates(input)

    c.JSON(http.StatusOK, task)
}
