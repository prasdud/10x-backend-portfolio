package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prasdud/10x-backend-portfolio/internal/handlers"
)

func main() {

	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1.GET("/ping", handlers.Ping)
	v1.GET("/resume", handlers.GetResume)
	v1.GET("/skills", handlers.GetSkills)
	v1.GET("/projects", handlers.GetProjects)
	v1.GET("/experience/:id", handlers.GetExperience)
	v1.GET("/contact", handlers.GetContact)
	v1.GET("/blog", handlers.GetBlog)
	v1.GET("/about", handlers.GetAbout)
	v1.GET("/easteregg1", handlers.GetEasterEgg1)
	v1.GET("/easteregg2", handlers.GetEasterEgg2)
	v1.GET("/help", handlers.Help)

	r.Run(":7741")
}
