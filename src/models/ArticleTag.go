package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type ArticleTag struct {
	gorm.Model

	Name       string
	CreatedBy  string
	ModifiedBy string
	//Secret bool `gorm:"DEFAULT:true"`//true:仅自己可见,默认true
	UserId uint
	//State int
}

type outputArticleTag struct {
	TagId      uint
	CreatedAt  time.Time
	UpdateAt   time.Time
	Name       string
	CreatedBy  string
	ModifiedBy string
}

func AddTag(TagName string, username string, userId uint) (uint, error) {
	tempArticleTag := new(ArticleTag)

	tempArticleTag = &ArticleTag{
		//Model:      gorm.Model{},
		Name:       TagName,
		CreatedBy:  username,
		ModifiedBy: username,
		UserId:     userId,
	}

	tx := DB.Begin()
	if tx.Create(tempArticleTag).RowsAffected != 1 {
		tx.Rollback()
		return 0, errors.New("Can'tInsertIntoDatabase")
	}
	tx.Commit()
	return tempArticleTag.ID, nil
}

func GetAllTags(userId int) interface{} {
	out := make([]*outputArticleTag, 0, 100)
	data := make([]*ArticleTag, 0, 100)

	DB.Where("user_id = ?", userId).Find(&data)
	for _, v := range data {
		out = append(out, &outputArticleTag{
			TagId:      v.ID,
			CreatedAt:  v.CreatedAt,
			UpdateAt:   v.UpdatedAt,
			Name:       v.Name,
			CreatedBy:  v.CreatedBy,
			ModifiedBy: v.ModifiedBy,
		})
	}
	return out
}
