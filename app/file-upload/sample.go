package fileupload

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ファイルアップロード
// https://gin-gonic.com/ja/docs/examples/upload-file/multiple-file/

func FileUploadSingleSample() *gin.Engine {
	router := gin.Default()
	// マルチパートフォームが利用できるメモリの制限を設定する(デフォルトは 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// 単一のファイル
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// 特定のディレクトリにファイルをアップロードする
		c.SaveUploadedFile(file, fmt.Sprintf("./file-upload/files/%v", file.Filename))

		c.String(http.StatusOK, fmt.Sprintf("'%s' 愛のあるファイル: ", file.Filename))
	})
	return router
}

func FileUploadMultiSample() *gin.Engine {
	router := gin.Default()
	// マルチパートフォームが利用できるメモリの制限を設定する(デフォルトは 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// マルチパートフォーム
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// 特定のディレクトリにファイルをアップロードする
			c.SaveUploadedFile(file, fmt.Sprintf("./file-upload/files/%v", file.Filename))
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	return router
}
