package mining

import (
	"github.com/gin-gonic/gin"
)

// 产出数字货币
func Routes(root *gin.Engine) {
	mining := root.Group("/mining")
	{
		mining.POST("/pow", PoW)
		mining.GET("/list", List)
	}
}
