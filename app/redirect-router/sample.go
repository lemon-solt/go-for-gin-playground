package redirectrouter

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// リダイレクト
// https://gin-gonic.com/ja/docs/examples/redirects/

func RedicretSample() *gin.Engine {

	r := gin.Default()
	r.GET("/googleRedirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})

	r.GET("/postRedirectGet", func(c *gin.Context) {
		fmt.Println("redicret: post")
	})

	r.POST("/loveRedirectExist", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, "<h1>愛はあるよ</h1>")
	})

	r.POST("/throwRedirectUrl", func(c *gin.Context) {
		fmt.Println("POSTされた情報に愛は")
		c.Redirect(http.StatusTemporaryRedirect, "/loveRedirectExist")
	})

	r.GET("/routerHandle", func(c *gin.Context) {
		c.Request.URL.Path = "/routerLoveExist"
		r.HandleContext(c)
	})
	r.GET("/routerLoveExist", func(c *gin.Context) {
		fmt.Println("handler url: ", c.Request.URL)
		c.JSON(200, gin.H{"愛は": "あるんか"})
	})

	return r
}
