package base

import (
	"net/http"

	"github.com/Parker-Yang/def-braveTroops/proto"
	"github.com/Parker-Yang/srv-braveTroops/global"
	"github.com/Parker-Yang/srv-braveTroops/modules/base"
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
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, state)
}
