package routerlog

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RouterLogSample() *gin.Engine {
	gin.DisableConsoleColor()
	f, _ := os.Create("love.log")

	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("愛はあるんかログ %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r.POST("/loveLogRouter1", func(c *gin.Context) {
		c.JSON(http.StatusOK, "そこに")
	})

	r.GET("/loveLogRouter2", func(c *gin.Context) {
		c.JSON(http.StatusOK, "愛は")
	})

	r.GET("/loveLogRouter3", func(c *gin.Context) {
		c.JSON(http.StatusOK, "あるんか")
	})
	return r
}
