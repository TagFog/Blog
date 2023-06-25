package Api

import "web/Api/settings_api"

type ApiGroup struct {
	SettingApi settings_api.SettingApi
}

//实例化对象
var ApiGroupApp = new(ApiGroup)
