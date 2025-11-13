package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/prasdud/10x-backend-portfolio/internal/handlers"
)

func ping(c *gin.Context){
	c.JSON(200, gin.H{"message": "pong"})
}


func getResume(c *gin.Context){
	c.File("https://prasdud.com/314159265358.pdf")
}

func getSkills(c *gin.Context){
	c.JSON(200, gin.H{
			"Python": "Master",
			"Go":"Proficient",
	})
}


func getProjects(c *gin.Context){
	c.JSON(200, gin.H{
			"DISAC": "Data integrity system for academic credentials",
			"Raptor":"Red team AI malware simulator",
	})
}


func getExperience(c *gin.Context){
	c.JSON(200, gin.H{
			"Omni-Tech Solutions": "SWE Intern",
			"Dragbin":"SWE Intern",
	})
}


func getContact(c *gin.Context){
	c.JSON(200, gin.H{
			"Linkedin": "LINKEDINURL",
			"X / Twitter": "XURL",
			"Phone Number": "XXXXXXXX",
			"Email": "X@Y.com",
	})
}


func getBlog(c *gin.Context){
	c.Redirect(302, "https://blog.prasdud.com"
}


func getAbout(c *gin.Context){
	c.JSON(200, gin.H{
			"Abdul": "who is he"
	})
}


func getEasterEgg1(c *gin.Context){
	c.JSON(200, gin.H{
			"Easter egg ": "#1"
	})
}

func getEasterEgg2(c *gin.Context){
	c.JSON(200, gin.H{
			"Easter egg": "#2"
	})
}


