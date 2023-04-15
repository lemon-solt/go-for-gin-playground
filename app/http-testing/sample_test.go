package httptesting

import (
	"net/http"
	tst "net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// テスト
// https://gin-gonic.com/ja/docs/testing/

func loveServer() *gin.Engine {
	r := gin.Default()
	r.GET("/loveExist", func(c *gin.Context) {
		text := c.Query("love")
		c.String(200, text)
	})
	return r
}

func TestPingRouteSample(t *testing.T) {
	router := loveServer()

	w := tst.NewRecorder()
	req, _ := http.NewRequest("GET", "/loveExist?love=あるんか", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "あるんか", w.Body.String())
}
