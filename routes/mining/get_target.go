package mining

import (
	"github.com/Evolt0/def-braveTroops/consts/status"
	"github.com/Evolt0/srv-braveTroops/global"
	"github.com/Evolt0/srv-braveTroops/modules/mining"
	"github.com/Evolt0/srv-braveTroops/utils"
	"github.com/gin-gonic/gin"
)

func GetTarget(ctx *gin.Context) {
	controller := mining.NewController(global.Conf.Client)
	state, err := controller.GetTarget()
	if err != nil {
		utils.UnWarpFailedResp(ctx, err)
		return
	}
	utils.SuccessResp(ctx, status.Ok, state)
}
