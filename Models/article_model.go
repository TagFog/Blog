package Models

import "web/Models/ctype"

type ArticleModel struct {
	MODEL
	Title      string      `gorm:"size:32" json:"title"`                    //标题
	Abstract   string      `json:"abstract"`                                //简介
	Content    string      `json:"content"`                                 //内容
	UserModel  UserModel   `gorm:"foreignKey:UserID" json:"-"`              //作者
	UserID     uint        `json:"user_id"`                                 //作者id
	TagModels  []TagModel  `gorm:"many2many:article_tag" json:"tag_models"` //文章标签
	Banner     BannerModel `json:"-"`                                       //封面
	BannerID   uint        `json:"banner_id"`
	BannerPath string      `json:"banner_path"`
	Tags       ctype.Array `gorm:"type:string;size:64" json:"tags"` //标签
}
