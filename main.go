package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xwxb/xhs-dl/consts"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	initRouters(r)
	r.Run(consts.DefaultPort)
}
