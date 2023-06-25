package settings_api

import (
	"github.com/gin-gonic/gin"
	"web/Models/res"
)

func (SettingApi) SettingsInfoView(c *gin.Context) {
	res.FailWithCode(1001, c)
}
