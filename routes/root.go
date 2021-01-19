package routes

import (
	"github.com/Parker-Yang/srv-braveTroops/routes/base"
	"github.com/Parker-Yang/srv-braveTroops/routes/mining"
	"github.com/Parker-Yang/srv-braveTroops/routes/ledger"
	"github.com/Parker-Yang/srv-braveTroops/routes/user"
	"github.com/gin-gonic/gin"
)

func SetRoutes(engine *gin.Engine) {
	base.Routes(engine)
	mining.Routes(engine)
	ledger.Routes(engine)
	user.Routes(engine)
}
