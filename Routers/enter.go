package Routers

import (
	"github.com/gin-gonic/gin"
	"web/Global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(Global.Config.System.Env)
	r := gin.Default()
	apiRouterGroup := r.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	routerGroupApp.SettingRouter()
	return r
}
