package Models

import (
	"web/Models/ctype"
)

type Role int

type UserModel struct {
	MODEL
	NickName   string           `gorm:"size:36" json:"nick_name"`
	UserName   string           `gorm:"size:36" json:"user_name"`
	Password   string           `gorm:"size:128" json:"password"`
	Avatar     string           `gorm:"size:256" json:"avatar"`
	Email      string           `gorm:"size:128" json:"email"`
	Tel        string           `gorm:"size:18" json:"tel"`
	Addr       string           `gorm:"size:64" json:"addr"`
	Token      string           `gorm:"size:64" json:"token"`
	Ip         string           `gorm:"size:20" json:"ip"`
	Role       ctype.Role       `gorm:"size4;default:1" json:"role"`
	SignStatus ctype.SignStatus `gorm:"type=smallint(6)" json:"signStatus"`
	//ArticleModels  []ArticleModel   `gorm:"foreignKey:AuthID" json:"-"`
	//CollectsModels []ArticleModel   `gorm:"many2many:auth2_collect;joinForeignKey:AuthID;JoinReference:ArticleID" json:"-"`
}
