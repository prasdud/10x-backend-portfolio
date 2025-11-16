package handlers

import (
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func GetResume(c *gin.Context) {
	c.File("./static/resume.pdf")
}

func GetSkills(c *gin.Context) {
	c.JSON(200, gin.H{
		"Python": "Master",
		"Go":     "Proficient",
	})
}

func GetProjects(c *gin.Context) {
	c.JSON(200, gin.H{
		"DISAC":  "Data integrity system for academic credentials",
		"Raptor": "Red team AI malware simulator",
	})
}

func GetExperience(c *gin.Context) {
	id := c.Param("id")
	experiences := map[string]string{
		"omni-tech": "Omni-Tech Solutions: SWE Intern",
		"dragbin":   "Dragbin: SWE Intern",
	}
	if exp, ok := experiences[id]; ok {
		c.JSON(200, gin.H{"experience": exp})
	} else {
		c.JSON(404, gin.H{"error": "Experience not found"})
	}
}

func GetContact(c *gin.Context) {
	c.JSON(200, gin.H{
		"Linkedin":     "LINKEDINURL",
		"X / Twitter":  "XURL",
		"Phone Number": "XXXXXXXX",
		"Email":        "X@Y.com",
	})
}

func GetBlog(c *gin.Context) {
	c.Redirect(302, "https://blog.prasdud.com")
}

func GetAbout(c *gin.Context) {
	c.JSON(200, gin.H{
		"Abdul": "who is he",
	})
}

func GetEasterEgg1(c *gin.Context) {
	c.JSON(200, gin.H{
		"Easter egg ": "#1",
	})
}

func GetEasterEgg2(c *gin.Context) {
	c.JSON(200, gin.H{
		"Easter egg": "#2",
	})
}
