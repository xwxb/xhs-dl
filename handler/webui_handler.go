package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xwxb/xhs-dl/pkg/parse"
)

func WebUIIndex(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func ParseResultHtml(c *gin.Context) {
	url := c.Query("orig")
	imageUrls, err := parse.RoboParse(url)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.HTML(200, "result.html", imageUrls)
}
