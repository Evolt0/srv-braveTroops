package base

import (
	"github.com/Parker-Yang/def-braveTroops/consts/status"
	"github.com/Parker-Yang/def-braveTroops/proto"
	"github.com/Parker-Yang/srv-braveTroops/global"
	"github.com/Parker-Yang/srv-braveTroops/modules/base"
	"github.com/Parker-Yang/srv-braveTroops/utils"
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
