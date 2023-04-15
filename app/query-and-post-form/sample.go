package queryandpostform

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// フォーム投稿によるクエリ文字列
// https://gin-gonic.com/ja/docs/examples/query-and-post-form/
func QueryAndPostSample() *gin.Engine {

	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		s := fmt.Sprintf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		fmt.Println("愛: ", s)
		// 愛:  id: 1; page: 1; name: menu; message: is that exist love?
	})
	return router
}
