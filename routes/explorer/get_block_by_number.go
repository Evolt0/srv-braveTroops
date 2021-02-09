package explorer

import (
	"github.com/Evolt0/def-braveTroops/consts/status"
	"github.com/Evolt0/srv-braveTroops/global"
	"github.com/Evolt0/srv-braveTroops/utils"
	"github.com/gin-gonic/gin"
)

func GetBlockByNumber(ctx *gin.Context) {
	number, err := utils.QueryUInt64(ctx, "number")
	if err != nil {
		utils.FailedResp(ctx, status.BadRequest, nil)
		return
	}
	block, err := global.Conf.Client.LedgerQueryDecodedBlock(number)
	if err != nil {
		utils.FailedResp(ctx, status.InternalServerError, nil)
		return
	}
	utils.SuccessResp(ctx, status.Ok, block)
}
