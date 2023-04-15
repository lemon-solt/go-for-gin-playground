package paraminpath

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	パスに含まれるパラメータ

https://gin-gonic.com/ja/docs/examples/param-in-path/
*/
func ParaminpathSample() *gin.Engine {
	router := gin.Default()

	// このハンドラは /user/john にはマッチするが、/user/ や /user にはマッチしない
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "そこに、 %s", name)
	})

	// しかし、下記は /user/john/ と /user/john/send にマッチする
	// もしほかのルーターが /user/john にマッチしなければ、/user/john/ にリダイレクトしてくれる
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	return router
}
