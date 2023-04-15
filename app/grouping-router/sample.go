package groupingrouter

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// ルーティングをグループ化する
// https://gin-gonic.com/ja/docs/examples/grouping-routes/

func loveEndpoint(c *gin.Context) {
	fmt.Println("愛はあるんかログ: ", c.PostForm("love"))

	c.Header("Content-Type", "text/html")
	file, _ := os.ReadFile("./grouping-router/index.html")
	s := string(file)
	c.String(http.StatusOK, s)
}

func GroupingRouterSample() *gin.Engine {
	router := gin.Default()

	// v1 のグループ
	v1 := router.Group("/v1")
	{
		v1.POST("/isLoveExist", loveEndpoint)
	}

	// v2 のグループ
	v2 := router.Group("/v2")
	{
		v2.POST("/isLoveExist", loveEndpoint)
	}
	return router
}
