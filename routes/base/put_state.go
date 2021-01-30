package base

import (
	"github.com/Evolt0/def-braveTroops/consts/status"
	"github.com/Evolt0/def-braveTroops/proto"
	"github.com/Evolt0/srv-braveTroops/global"
	"github.com/Evolt0/srv-braveTroops/modules/base"
	"github.com/Evolt0/srv-braveTroops/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"reflect"
)

func Put(ctx *gin.Context) {
	data := &proto.PutState{}
	err := ctx.BindJSON(data)
	if err != nil {
		utils.FailedResp(ctx, status.BadRequest, nil)
		return
	}
	logrus.Debugln(reflect.TypeOf(data))
	controller := base.NewController(global.Conf.Client)
	state, err := controller.PutState(data)
	if err != nil {
		utils.UnWarpFailedResp(ctx, err)
		return
	}
	utils.SuccessResp(ctx, status.Ok, state)
}
