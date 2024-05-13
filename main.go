package main

import (
	"github.com/gin-gonic/gin"
	"xhs-dl/consts"
)

func main() {
	r := gin.Default()
	initRouters(r)
	r.Run(consts.DefaultPort)
}
