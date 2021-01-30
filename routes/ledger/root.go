package ledger

import (
	"github.com/gin-gonic/gin"
)

// 货币交易
func Routes(root *gin.Engine) {
	ledger := root.Group("/ledger")
	{
		ledger.POST("/transfer", Transfer)
		ledger.GET("/list", List)
	}
}
