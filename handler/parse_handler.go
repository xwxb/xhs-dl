package handler

import (
	"github.com/gin-gonic/gin"
	"xhs-dl/parse"
)

func Parse(c *gin.Context) {
	url := c.Query("orig")
	imageUrls := parse.Scrape(url)
	c.JSON(200, gin.H{
		"imageUrls": imageUrls,
	})
}
