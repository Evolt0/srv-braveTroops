package routes

import (
	"github.com/Parker-Yang/srv-braveTroops/routes/base"
	"github.com/gin-gonic/gin"
)

func SetRoutes(engine *gin.Engine) {
	base.Routes(engine)
}
