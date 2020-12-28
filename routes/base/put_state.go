package base

import (
	"net/http"
	"reflect"

	"github.com/Parker-Yang/srv-braveTroops/global"
	"github.com/Parker-Yang/srv-braveTroops/models"
	"github.com/Parker-Yang/srv-braveTroops/modules/base"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Put(ctx *gin.Context) {
	data := &models.PutState{}
	err := ctx.BindJSON(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	logrus.Debugln(reflect.TypeOf(data))
	controller := base.NewController(global.Conf.Client)
	state, err := controller.PutState(*data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, state)
}
