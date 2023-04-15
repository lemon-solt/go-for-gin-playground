package validationbind

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// デルへのバインディングとバリデーション
// https://gin-gonic.com/ja/docs/examples/binding-and-validation/

// JSON からバインドする
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func ValidationBindSample() *gin.Engine {
	router := gin.Default()

	// JSON でバインドする例 ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"that is love error": err.Error()})
			return
		}

		if json.User != "love" || json.Password != "exist" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "that is loved in"})
	})

	// XML でバインドする例 (
	//	<?xml version="1.0" encoding="UTF-8"?>
	//	<root>
	//		<user>manu</user>
	//		<password>123</password>
	//	</root>)
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "thats is love!"})
	})

	// HTML Form からバインドする例 (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// このコードは、content-type ヘッダーから類推して HTML Form でバインドする
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
	return router
}
