package models

type Document struct {
	ID             string                `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Path           string                `gorm:"type:varchar(200);unique"json:"path"`
	Title          string                `gorm:"type:varchar(50);" json:"title"`
	Cover          string                `gorm:"type:varchar(200)"json:"cover"`    // 文档阅读时顶部图，在外部作为文档卡片的封面，在这里为封面图片的地址
	Up             int                   `gorm:"type:int"json:"up"`                // 👍
	Down           int                   `gorm:"type:int"json:"down"`              // 👎
	ReadCount      int                   `gorm:"type:int" json:"readCount"`        // 阅读数量
	CommentCount   int                   `gorm:"type:int" json:"commentCount"`     // 评论数量
	WordsCount     int                   `gorm:"type:int" json:"words_count"`      // 文章字数
	CreateTime     int64                 `gorm:"type:bigint"json:"createTime"`     // 创建时间
	LastModifyTime int64                 `gorm:"type:bigint"json:"lastModifyTime"` // 最后修改时间
	CategoryID     string                `gorm:"type:uuid;"`                       // 隶属的分类
	UserID         string                `gorm:"type:uuid;"`                       // 文档创建者
	Contributors   []DocumentContributor `json:"contributors"`
}
