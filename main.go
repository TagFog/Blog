package main

import (
	"web/Core"
	"web/Flag"
	"web/Global"
	"web/Routers"
)

func main() {
	Core.InitConf()
	Global.DB = Core.InitGorm()
	Global.Log = Core.InitLogger()
	r := Routers.InitRouter()

	op := Flag.Parse()
	if Flag.IsWebStop(op) {
		Flag.SwitchOption(op)
		return
	}

	add := Global.Config.System.Addr()
	Global.Log.Infof("web 运行在: %s", add)
	err := r.Run(add)
	if err != nil {
		Global.Log.Fatalf(err.Error())
	}
}
