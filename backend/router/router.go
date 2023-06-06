package router

import (
	"todo/controller"
	"todo/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(g *gin.Engine, c controller.Controller) *gin.Engine {
	g.POST("/login", func(ctx *gin.Context) {
		c.Login.Login(ctx)
	})

	r := g.Group("/api")
	r.Use(middleware.ValidateToken())
	r.GET("/tasks", func(ctx *gin.Context) {
		c.Task.GetTaskList(ctx)
	})
	r.POST("/tasks", func(ctx *gin.Context) {
		c.Task.AddTask(ctx)
	})
	r.PUT("/tasks/:id", func(ctx *gin.Context) {
		c.Task.UpdateTask(ctx)
	})
	r.DELETE("/tasks/:id", func(ctx *gin.Context) {
		c.Task.DeleteTask(ctx)
	})

	r.GET("/:id", func(ctx *gin.Context) {
		c.Task.GetTaskList(ctx)
	})

	return g
}
