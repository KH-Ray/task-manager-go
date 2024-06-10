package main

import (
	"task-manager-be/controllers"
	"task-manager-be/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
    r.Use(cors.Default())

    db.ConnectDB()

    v1 := r.Group("/api/v1")
    v1.GET("/tasks", controllers.FindTasks)
    v1.GET("/tasks/:id", controllers.FindTask)
    v1.POST("/tasks", controllers.CreateTask)
    v1.DELETE("/tasks/:id", controllers.DeleteBook)
    v1.PATCH("/tasks/:id", controllers.UpdateTask)

    r.Run(":3001")
}
