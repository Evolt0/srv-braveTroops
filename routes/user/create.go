package user

import (
	"github.com/Parker-Yang/def-braveTroops/consts/status"
	"github.com/Parker-Yang/def-braveTroops/proto"
	"github.com/Parker-Yang/srv-braveTroops/global"
	"github.com/Parker-Yang/srv-braveTroops/modules/user"
	"github.com/Parker-Yang/srv-braveTroops/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"reflect"
)

//
func Create(ctx *gin.Context) {
	data := &proto.UserReq{}
	err := ctx.ShouldBindJSON(data)
	if err != nil {
		utils.FailedResp(ctx, status.BadRequest, nil)
		return
	}
	logrus.Debugln(reflect.TypeOf(data))
	controller := user.NewController(global.Conf.Client)
	state, err := controller.Create(data)
	if err != nil {
		utils.UnWarpFailedResp(ctx, err)
		return
	}
	utils.SuccessResp(ctx, status.Ok, state)
}
