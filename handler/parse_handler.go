package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xwxb/xhs-dl/pkg/parse"
)

func Parse(c *gin.Context) {
	url := c.Query("orig")
	imageUrls := parse.Scrape(url)
	c.JSON(200, gin.H{
		"imageUrls": imageUrls,
	})
}
