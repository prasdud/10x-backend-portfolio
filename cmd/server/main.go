package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prasdud/10x-backend-portfolio/internal/handlers"
	"github.com/prasdud/10x-backend-portfolio/internal/store"
)

func main() {
	// Load portfolio data from JSON
	if err := store.Init(); err != nil {
		log.Fatalf("Failed to load portfolio data: %v", err)
	}

	r := gin.Default()

	version_enabled := false
	var version *gin.RouterGroup

	if version_enabled {
		version = r.Group("api/v1")
	} else {
		version = r.Group("/")
	}

	version.GET("/ping", handlers.Ping)
	version.GET("/resume", handlers.GetResume)
	version.GET("/skills", handlers.GetSkills)
	version.GET("/projects", handlers.GetProjects)
	version.GET("/experience", handlers.GetExperience)
	version.GET("/experience/:id", handlers.GetExperience)
	version.GET("/contact", handlers.GetContact)
	version.GET("/blog", handlers.GetBlog)
	version.GET("/about", handlers.GetAbout)
	version.GET("/easteregg1", handlers.GetEasterEgg1)
	version.GET("/easteregg2", handlers.GetEasterEgg2)
	version.GET("/help", handlers.Help)

	r.Run(":7741")
}
