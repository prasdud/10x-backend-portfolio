package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prasdud/10x-backend-portfolio/internal/handlers"

)


func pingHandler(c *gin.Context){
	c.JSON(200, gin.H{"message": "pong"})
}


func main(){
	
	r := gin.Default()
	
	v1 := r.Group("/api/v1")
	v1.GET("/ping", handlers.ping)
	v1.GET("/resume", handlers.getResume)
	v1.GET("/skills", handlers.getSkills)
	v1.GET("/projects", handlers.getProjects)
	v1.GET("/experience", handlers.getExperience)
	v1.GET("/contact", handlers.getContact)
	v1.GET("/blog", handlers.getBlog)
	v1.GET("/about", handlers.getAbout)
	v1.GET("/easteregg1", handlers.getEasterEgg1)
	v1.GET("/easteregg2", handlers.getEasterEgg2)



	r.GET("/ping", pingHandler)

	r.Run(":7741")
}
