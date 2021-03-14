package user

import (
	"github.com/Evolt0/def-braveTroops/consts/status"
	"github.com/Evolt0/srv-braveTroops/global"
	"github.com/Evolt0/srv-braveTroops/modules/user"
	"github.com/Evolt0/srv-braveTroops/utils"
	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	controller := user.NewController(global.Conf.Client)
	state, err := controller.List()
	if err != nil {
		utils.UnWarpFailedResp(ctx, err)
		return
	}
	utils.SuccessResp(ctx, status.Ok, state)
}
