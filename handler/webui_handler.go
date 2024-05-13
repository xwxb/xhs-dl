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
	imageUrls := parse.Scrape(url)
	c.HTML(200, "result.html", imageUrls)
}
