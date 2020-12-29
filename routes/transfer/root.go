package transfer

import (
	"github.com/gin-gonic/gin"
)

// 货币交易
func Routes(root *gin.Engine) {
	transfer := root.Group("/transfer")
	{
		transfer.POST("/", transferOut)
	}
}

// 转账
func transferOut(context *gin.Context) {
	
}
