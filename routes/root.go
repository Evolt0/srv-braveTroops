package routes

import (
	"github.com/Evolt0/srv-braveTroops/routes/base"
	"github.com/Evolt0/srv-braveTroops/routes/explorer"
	"github.com/Evolt0/srv-braveTroops/routes/ledger"
	"github.com/Evolt0/srv-braveTroops/routes/mining"
	"github.com/Evolt0/srv-braveTroops/routes/user"
	"github.com/gin-gonic/gin"
)

func SetRoutes(engine *gin.Engine) {
	base.Routes(engine)
	mining.Routes(engine)
	ledger.Routes(engine)
	user.Routes(engine)
	explorer.Routes(engine)
}
