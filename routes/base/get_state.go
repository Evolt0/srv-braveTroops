package base

import (
	"github.com/Parker-Yang/def-braveTroops/consts/status"
	"github.com/Parker-Yang/def-braveTroops/proto"
	"github.com/Parker-Yang/srv-braveTroops/global"
	"github.com/Parker-Yang/srv-braveTroops/modules/base"
	"github.com/Parker-Yang/srv-braveTroops/utils"
	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	key := ctx.Param("key")
	controller := base.NewController(global.Conf.Client)
	State := &proto.GetState{
		Key: key,
	}
	state, err := controller.GetState(State)
	if err != nil {
		utils.UnWarpFailedResp(ctx, err)
		return
	}
	utils.SuccessResp(ctx, status.Ok, state)
}
