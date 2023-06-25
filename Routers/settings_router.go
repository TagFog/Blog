package Routers

import (
	"web/Api"
)

func (router RouterGroup) SettingRouter() {
	settingapi := Api.ApiGroupApp.SettingApi

	router.GET("", settingapi.SettingsInfoView)
}
