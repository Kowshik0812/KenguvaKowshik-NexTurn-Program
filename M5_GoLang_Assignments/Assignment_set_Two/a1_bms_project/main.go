package main

import (
	db "blogmanager/config"
	"blogmanager/controller"
	"blogmanager/middleware"
	"blogmanager/repository"
	"blogmanager/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitializeDatabase()

	blogRepo := repository.NewBlogRepository(db.GetDB())
	blogService := service.NewBlogService(blogRepo)
	blogController := controller.NewBlogController(blogService)

	r := gin.Default()

	r.Use(middleware.LoggingMiddleware())

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware(db.GetDB()))

	api.POST("/blog", blogController.CreateBlog)
	api.GET("/blog/:id", blogController.GetBlog)
	api.GET("/blog", blogController.GetAllBlogs)
	api.PUT("/blog/:id", blogController.UpdateBlog)
	api.DELETE("/blog/:id", blogController.DeleteBlog)

	r.Run(":8080")
}
