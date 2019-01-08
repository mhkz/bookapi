package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 文章模型
type Article struct {
	Model

	Title         string `json:"title"`               // 文章标题
	Keywords      string `json:"keywords"`            // 文章关键字 seo
	Description   string `json:"title"`               // 文章描述
	Content       string `json:"content"`             // 文章内容
	Thumb         string `json:"thumb"`               // 文章缩略图
	State         int    `json:"state"`               // 文章发布状态 -1 回收站 0 草稿 1 已发布
	Public        int    `json:"public"`              // 文章公开状态
	Origin        int    `json:"origin"`              // 文章转载状态 0 原创 1 转载 2 混合
	Password      string `json:"password"`            // 文章密码
	CreatedAt     string `json:"created_at"`          // 创建时间
	ModifiedAt    string `json:"modified_at"`         // 修改时间
	TagID         int    `json:"tag_id" gorm:"index"` // 标签 id
	Category      int    `json:"category"`            // 分类信息
	coverImageUrl string `json:"cover_image_url"`     // 封面url
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}

	return false
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.TagID)

	return
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:       data["tag_id"].(int),
		Title:       data["title"].(string),
		Description: data["desc"].(string),
		Content:     data["content"].(string),
		CreatedAt:   data["created_by"].(string),
		State:       data["state"].(int),
	})

	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

// 硬删除文章
func CleanAllArticle() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{})

	return true
}
