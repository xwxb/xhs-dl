package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xwxb/xhs-dl/pkg/parse"
)

func Parse(c *gin.Context) {
	url := c.Query("orig")
	imageUrls, err := parse.RoboParse(url)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"imageUrls": imageUrls,
	})
}
