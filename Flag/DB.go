package Flag

import (
	"web/Global"
	"web/Models"
)

func Makemigration() {
	//fmt.Println("迁移表结构")
	var err error
	Global.DB.SetupJoinTable(&Models.UserModel{}, "ConllectsModels", &Models.User2Collects{})
	Global.DB.SetupJoinTable(&Models.MenuModel{}, "Banners", &Models.MenuImageModel{})
	err = Global.DB.Set("gorm:tables_options", "ENGINE=InnoDB").
		AutoMigrate(
			&Models.BannerModel{},
			&Models.TagModel{},
			&Models.MenuModel{},
			&Models.ArticleModel{},
			&Models.MenuImageModel{},
			&Models.UserModel{},
			&Models.User2Collects{},
		)
	if err != nil {
		Global.Log.Error("{ error } 生成数据库表结构失败")
		return
	}
	Global.Log.Infof("{ success } 生成数据库表结构成功")
}
