package user

import (
	"github.com/gin-gonic/gin"
)

func Routes(ctx *gin.Engine) {
	user := ctx.Group("/user")
	user.POST("/",Create)
}
