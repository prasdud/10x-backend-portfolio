package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/prasdud/10x-backend-portfolio/internal/store"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": store.Data.Ping.Message})
}

func GetResume(c *gin.Context) {
	c.File("../../static/resume.pdf")
}

func GetSkills(c *gin.Context) {
	c.JSON(200, store.Data.Skills)
}

func GetProjects(c *gin.Context) {
	c.JSON(200, store.Data.Projects)
}

func GetExperience(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(200, store.Data.Experiences)
		return
	}
	if exp, ok := store.Data.Experiences[id]; ok {
		c.JSON(200, gin.H{"experience": exp})
	} else {
		c.JSON(404, gin.H{"error": "Experience not found"})
	}
}

func GetContact(c *gin.Context) {
	c.JSON(200, store.Data.Contact)
}

func GetBlog(c *gin.Context) {
	c.Redirect(302, store.Data.Blog.URL)
}

func GetAbout(c *gin.Context) {
	c.JSON(200, store.Data.About)
}

func GetEasterEgg1(c *gin.Context) {
	c.JSON(200, store.Data.EasterEggs.Egg1)
}

func GetEasterEgg2(c *gin.Context) {
	c.JSON(200, store.Data.EasterEggs.Egg2)
}

func Help(c *gin.Context) {
	c.JSON(200, store.Data.Help)
}
