package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xwxb/xhs-dl/handler"
)

func initRouters(r *gin.Engine) {
	webui := r.Group("/webui")
	{
		webui.GET("/", handler.WebUIIndex)
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
