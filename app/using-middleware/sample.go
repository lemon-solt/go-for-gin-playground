package usingmiddleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// goroutineエンドポイント
// https://gin-gonic.com/ja/docs/examples/goroutines-inside-a-middleware/
func GoroutineApiSample() *gin.Engine {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		// goroutine 内で使用するコピーを生成します
		cCp := c.Copy()
		go func() {
			// time.Sleep() を使って、長時間かかる処理をシミュレートします。5秒です。
			time.Sleep(5 * time.Second)

			// コピーされた context である "cCp" を使ってください。重要！
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		// time.Sleep() を使って、長時間かかる処理をシミュレートします。5秒です。
		time.Sleep(5 * time.Second)

		// goroutine を使ってなければ、context をコピーする必要はありません。
		log.Println("Done! in path " + c.Request.URL.Path)
	})
	return r
}
