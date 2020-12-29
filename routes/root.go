package routes

import (
	"github.com/Parker-Yang/srv-braveTroops/routes/base"
	"github.com/Parker-Yang/srv-braveTroops/routes/mining"
	"github.com/Parker-Yang/srv-braveTroops/routes/transfer"
	"github.com/gin-gonic/gin"
)

func SetRoutes(engine *gin.Engine) {
	base.Routes(engine)
	mining.Routes(engine)
	transfer.Routes(engine)
}
