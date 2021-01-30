package base

import (
	"github.com/Evolt0/def-braveTroops/consts/status"
	"github.com/Evolt0/def-braveTroops/proto"
	"github.com/Evolt0/srv-braveTroops/global"
	"github.com/Evolt0/srv-braveTroops/modules/base"
	"github.com/Evolt0/srv-braveTroops/utils"
	"github.com/gin-gonic/gin"
)

func getHistory(ctx *gin.Context) {
	if key, ok := ctx.GetQuery("key"); ok {
		controller := base.NewController(global.Conf.Client)
		State := &proto.GetHistoryRequest{
			Key: key,
		}
		state, err := controller.GetHistory(State)
		if err != nil {
			utils.UnWarpFailedResp(ctx, err)
			return
		}
		utils.SuccessResp(ctx, status.Ok, state)
		return
	}
	utils.FailedResp(ctx, status.BadRequest, nil)
}
