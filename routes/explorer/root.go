package explorer

import (
	"github.com/gin-gonic/gin"
)

func Routes(root *gin.Engine) {
	explorer := root.Group("/explorer")
	explorer.GET("/blockNumber", GetBlockByNumber)
	explorer.GET("/blockHash", GetBlockByHash)
	explorer.GET("/blockTx", GetBlockByTx)
	explorer.GET("/blockChainInfo", GetBlockChainInfo)
	explorer.GET("/getTx", GetTx)
}
