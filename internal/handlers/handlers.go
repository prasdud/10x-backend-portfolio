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
	// Get query parameters
	level := c.Query("level")
	category := c.Query("category")

	// If no filters, return all skills
	if level == "" && category == "" {
		c.JSON(200, store.Data.Skills)
		return
	}

	// Filter skills based on query params
	filtered := []store.Skill{}
	for _, skill := range store.Data.Skills {
		match := true
		if level != "" && skill.Level != level {
			match = false
		}
		if category != "" && skill.Category != category {
			match = false
		}
		if match {
			filtered = append(filtered, skill)
		}
	}

	c.JSON(200, filtered)
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
	// Find experience by id
	for _, exp := range store.Data.Experiences {
		if exp.ID == id {
			c.JSON(200, exp)
			return
		}
	}
	c.JSON(404, gin.H{"error": "Experience not found"})
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
