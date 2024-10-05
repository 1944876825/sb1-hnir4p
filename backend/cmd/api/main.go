package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/blog-system/internal/auth"
	"github.com/yourusername/blog-system/internal/database"
	"github.com/yourusername/blog-system/internal/models"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	r := gin.Default()

	// Public routes
	r.POST("/api/register", models.RegisterUser(db))
	r.POST("/api/login", auth.LoginUser(db))

	// Protected routes
	authorized := r.Group("/api")
	authorized.Use(auth.JWTAuthMiddleware())
	{
		authorized.GET("/posts", models.GetPosts(db))
		authorized.POST("/posts", models.CreatePost(db))
		authorized.GET("/posts/:id", models.GetPost(db))
		authorized.POST("/posts/:id/comments", models.AddComment(db))
		authorized.POST("/posts/:id/like", models.LikePost(db))
	}

	r.Run(":8080")
}