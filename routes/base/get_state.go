package base

import (
	"github.com/Evolt0/def-braveTroops/consts/status"
	"github.com/Evolt0/def-braveTroops/proto"
	"github.com/Evolt0/srv-braveTroops/global"
	"github.com/Evolt0/srv-braveTroops/modules/base"
	"github.com/Evolt0/srv-braveTroops/utils"
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
