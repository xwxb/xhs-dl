package main

import (
	"flag"
	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.String("port", "8080", "Port for the server to run on")
	flag.Parse()

	r := gin.Default() // Creates a gin router with default middleware: logger and recovery
	//gin.SetMode(gin.ReleaseMode)
	//gin.TrustedProxies = []string{
	r.LoadHTMLGlob("templates/*")

	initRouters(r)
	r.Run(":" + *port)
}
