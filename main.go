package main

import (
	routes "github.com/bagustri_3/SemestaArusTeknologi/routes"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "success"})
	})
}
