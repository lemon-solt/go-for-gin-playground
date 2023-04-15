package requestbind

import "github.com/gin-gonic/gin"

type ExistStruct struct {
	ExistParams string `form:"pointerParams"`
}

type LoveStruct struct {
	NestParams string `form:"nestParams"`
}

type ThatStruct struct {
	Params      string `form:"params"`
	Params2     string `form:"params2"`
	LoveParams  LoveStruct
	ExistStruct *ExistStruct
}

func That(c *gin.Context) {
	var b ThatStruct
	c.Bind(&b)
	c.JSON(200, gin.H{
		"thatResponse":    b.Params,
		"thatResponse2":   b.Params2,
		"nestResponse":    b.LoveParams,
		"pointerResponse": b.ExistStruct,
	})
}

func RequestBindSample() *gin.Engine {
	router := gin.Default()
	router.GET("/that", That)
	// router.GET("/love", GetDataC)
	// router.GET("/exists", GetDataD)
	return router
}
