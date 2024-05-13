package main

import (
	"github.com/gin-gonic/gin"
	"xhs-dl/handler"
)

func initRouters(r *gin.Engine) {
	r.GET("/webui", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//url/api/v1/parse
	api := r.Group("/api")
	{
		apiv1 := api.Group("/v1")
		apiv1.GET("/parse", handler.Parse)
	}
}
