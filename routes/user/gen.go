package user

import (
	"github.com/Evolt0/def-braveTroops/consts/status"
	"github.com/Evolt0/srv-braveTroops/global"
	"github.com/Evolt0/srv-braveTroops/modules/user"
	"github.com/Evolt0/srv-braveTroops/utils"
	"github.com/gin-gonic/gin"
)

// 创建用户
func Gen(ctx *gin.Context) {
	controller := user.NewController(global.Conf.Client)
	state, err := controller.Gen()
	if err != nil {
		utils.FailedResp(ctx, status.InternalServerError, nil)
		return
	}
	utils.SuccessResp(ctx, status.Ok, state)
}
