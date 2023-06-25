package Models

type TagModel struct {
	MODEL
	Title    string         `gorm:"size:16" json:"title"`           //标签名称
	Articles []ArticleModel `gorm:"many2many:article_tag" json:"-"` //关联的标签
}
