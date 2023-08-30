package main

import (
	"github.com/gin-gonic/gin"
	"go_test/route"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	route.InitRouter(engine)

	engine.Run() // listen and serve on 0.0.0.0:8080
}
