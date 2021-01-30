package mining

import (
	"github.com/Evolt0/def-braveTroops/consts/status"
	"github.com/Evolt0/def-braveTroops/proto"
	"github.com/Evolt0/srv-braveTroops/global"
	"github.com/Evolt0/srv-braveTroops/modules/mining"
	"github.com/Evolt0/srv-braveTroops/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"reflect"
)

func GetTarget(ctx *gin.Context) {
	data := &proto.BodyData{}
	err := ctx.ShouldBindJSON(data)
	if err != nil {
		utils.FailedResp(ctx, status.BadRequest, nil)
		return
	}
	logrus.Debugln(reflect.TypeOf(data))
	controller := mining.NewController(global.Conf.Client)
	state, err := controller.GetTarget(data)
	if err != nil {
		utils.UnWarpFailedResp(ctx, err)
		return
	}
	utils.SuccessResp(ctx, status.Ok, state)
}
