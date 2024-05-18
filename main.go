package main

import (
	"flag"
	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.String("port", "8080", "Port for the server to run on")
	flag.Parse()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	initRouters(r)
	r.Run(":" + *port)
}
