package models

import (
	"archie/connection/postgres_conn"
	"archie/utils"
)

type Document struct {
	ID             string                `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Content        string                `gorm:"type:text" json:"content"`
	Title          string                `gorm:"type:varchar(50);" json:"title"`
	Headings       string                `gorm:"type:varchar(500)" json:"headings"`
	Excerpt        string                `gorm:"type:varchar(120);" json:"excerpt"`
	Cover          string                `gorm:"type:varchar(200)"json:"cover"`       // 文档阅读时顶部图，在外部作为文档卡片的封面，在这里为封面图片的地址
	Up             int                   `gorm:"type:int"json:"up"`                   // 👍
	Down           int                   `gorm:"type:int"json:"down"`                 // 👎
	ReadCount      int                   `gorm:"type:int" json:"readCount"`           // 阅读数量
	WordsCount     int                   `gorm:"type:int" json:"words_count"`         // 文章字数
	CreateTime     int32                 `gorm:"type:bigint"json:"createTime"`        // 创建时间
	LastModifyTime int32                 `gorm:"type:bigint"json:"lastModifyTime"`    // 最后修改时间
	CategoryID     string                `gorm:"type:varchar(36);" json:"categoryID"` // 隶属的分类
	UserID         string                `gorm:"type:uuid;" json:"userID"`            // 文档创建者
	OrganizationID string                `gorm:"type:uuid" json:"organizationID"`
	IsPublic       bool                  `gorm:"type:bool" json:"isPublic"` // 是否公开
	Contributors   []DocumentContributor `json:"contributors"`
}

type Heading struct {
	Level   int    `json:"level"`
	Content string `json:"content"`
}

type ParsedDocument struct {
	Document
	Headings []Heading `json:"headings"`
}

func (doc *Document) New() error {
	now := utils.Now()
	doc.CreateTime = now
	doc.LastModifyTime = now

	return postgres_conn.DB.Instance().Create(doc).Find(doc).Error
}

func (doc *Document) FindAll(docs *[]Document) error {
	return postgres_conn.DB.Instance().Find(&docs).Error
}

func (doc *Document) Detail() error {
	return postgres_conn.DB.Instance().Where("id = ?", doc.ID).Find(doc).Error
}

//func parseDocument(rawDocs []Document, target *[]ParsedDocument) error {
//	for _, rawDoc := range rawDocs {
//		var headings []Heading
//
//		if err := json.Unmarshal([]byte(rawDoc.Headings), &headings); err != nil {
//			return err
//		}
//
//		*target = append(*target, ParsedDocument{
//			Document: rawDoc,
//			Headings: headings,
//		})
//	}
//}
