package Flag

import sys_flag "flag"

//表结构迁移
type Option struct {
	Version bool
	DB      bool
}

func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据")
	sys_flag.Parse() //顺序结构
	return Option{
		DB: *db,
	}
}

//是否停止web项目
func IsWebStop(option Option) bool {
	if option.DB {
		return true
	}
	return false
}

func SwitchOption(option Option) {
	if option.DB {
		Makemigration()
	}
}
