package models

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model

	TagId       uint
	UserId      uint
	ArticleBody string `gorm:"size:1520"`
}

type outputArticle struct {
	TagId       uint
	ArticleBody string
}

func AddArticle(articleBody string, tagId uint, userId uint) (ok bool) {
	tempArticle := new(Article)
	tempArticle = &Article{
		//Model:       gorm.Model{},
		TagId:       tagId,
		UserId:      userId,
		ArticleBody: articleBody,
	}

	tx := DB.Begin()
	if tx.Create(tempArticle).RowsAffected != 1 {
		tx.Rollback()
		ok = false
		return
	}
	tx.Commit()
	ok = true
	return
}

func GetArticle(tagId int, userId int) interface{} {
	data := new(Article)
	DB.Where("user_id = ? AND tag_id = ?", userId, tagId).Find(data)
	out := outputArticle{
		TagId:       data.TagId,
		ArticleBody: data.ArticleBody,
	}
	return out
}
