package explorer

import (
	"github.com/Evolt0/def-braveTroops/consts/status"
	"github.com/Evolt0/srv-braveTroops/global"
	"github.com/Evolt0/srv-braveTroops/utils"
	"github.com/gin-gonic/gin"
)

func GetTx(ctx *gin.Context) {
	txID := ctx.Query("tx")
	block, err := global.Conf.Client.LedgerQueryDecodedBlockByTxID(txID)
	if err != nil {
		utils.FailedResp(ctx, status.InternalServerError, nil)
		return
	}
	utils.SuccessResp(ctx, status.Ok, block)
}
