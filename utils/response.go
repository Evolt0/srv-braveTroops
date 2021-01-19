package utils

import (
	"github.com/Parker-Yang/def-braveTroops/consts/status"
	"github.com/Parker-Yang/def-braveTroops/proto/epkg"
	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, code status.Status, data interface{}) {
	ctx.JSON(int(code), gin.H{
		"code": code,
		"msg":  status.GetMsg(code),
		"data": data,
	})
	ctx.Abort()
}

func SuccessResp(ctx *gin.Context, code status.Status, data interface{}) {
	Response(ctx, code, data)
}

func FailedResp(ctx *gin.Context, code status.Status, data interface{}) {
	Response(ctx, code, data)
}

func UnWarpFailedResp(ctx *gin.Context, err error) {
	fail, message := epkg.UnwrapFail(err)
	FailedResp(ctx, fail, message)
}
