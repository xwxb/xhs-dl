package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xwxb/xhs-dl/handler"
)

func initRouters(r *gin.Engine) {
	r.Use(gin.Recovery())

	r.GET("/webui", handler.WebUIIndex)
	webui := r.Group("/webui")
	{
		webui.GET("/", func(c *gin.Context) {
			c.Redirect(301, "/webui")
		})
		webui.GET("/index.html", handler.WebUIIndex)
		webui.GET("/result", handler.ParseResultHtml)
	}

	//url/api/v1/parse
	api := r.Group("/api")
	{
		apiv1 := api.Group("/v1")
		apiv1.GET("/parse", handler.Parse)
	}
}
