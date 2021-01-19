package base

import (
	"github.com/gin-gonic/gin"
)

func Routes(root *gin.Engine) {
	base := root.Group("/base")
	{
		base.PUT("/", Put)
		base.GET("/:key", Get)
	}
	bases := root.Group("/bases")
	{
		bases.GET("/history", getHistory)
	}
}
