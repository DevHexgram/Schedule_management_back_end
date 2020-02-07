package models

import (
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

func AddTagAndArticle(tagName string, username string, articleBody *string, userId uint) (ok bool) {
	tempArticleTag := new(ArticleTag)

	tempArticleTag = &ArticleTag{
		//Model:      gorm.Model{},
		Name:       tagName,
		CreatedBy:  username,
		ModifiedBy: username,
		UserId:     userId,
	}

	tx := DB.Begin()
	if tx.Create(tempArticleTag).RowsAffected != 1 {
		tx.Rollback()
		ok = false
		return
	}

	tempArticle := new(Article)
	tempArticle = &Article{
		//Model:       gorm.Model{},
		TagId:       tempArticleTag.ID,
		UserId:      userId,
		ArticleBody: *articleBody,
	}

	if tx.Create(tempArticle).RowsAffected != 1 {
		tx.Rollback()
		ok = false
		return
	}
	tx.Commit()
	ok = true
	return ok
}

func ModifyTag(tagName string, tagId int, userId int) (ok bool) {
	tx := DB.Begin()
	if tx.Model(&ArticleTag{}).Where("id = ? AND user_id = ?", tagId, userId).Update(&ArticleTag{
		//Model:      gorm.Model{},
		Name: tagName,
		//CreatedBy:  "",
		//ModifiedBy: "",
		//UserId:     0,
	}).RowsAffected != 1 {
		tx.Rollback()
		ok = false
		return
	}
	tx.Commit()
	ok = true
	return
}

func ModifyArticleWithTag(articleBody *string, tagId int, userId int) (ok bool) {
	tx := DB.Begin()
	if tx.Model(&Article{}).Where("tag_id = ? AND user_id = ?", tagId, userId).Update(&Article{
		//Model:       gorm.Model{},
		//TagId:       0,
		//UserId:      0,
		ArticleBody: *articleBody,
	}).RowsAffected != 1 {
		tx.Rollback()
		ok = false
		return
	}
	tx.Commit()
	ok = true
	return
}

//func AddTag(TagName string, username string, userId uint) (uint, error) {
//	tempArticleTag := new(ArticleTag)
//
//	tempArticleTag = &ArticleTag{
//		Model:      gorm.Model{},
//Name:       TagName,
//CreatedBy:  username,
//ModifiedBy: username,
//UserId:     userId,
//}
//
//tx := DB.Begin()
//if tx.Create(tempArticleTag).RowsAffected != 1 {
//	tx.Rollback()
//	return 0, errors.New("Can'tInsertIntoDatabase")
//}
//tx.Commit()
//return tempArticleTag.ID, nil
//}

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

func FindTag(tagId int, userId int) (isFound bool) {
	tempTag := new(ArticleTag)
	DB.Where("id = ? AND user_id = ?", tagId, userId).Find(tempTag)
	if tempTag.ID <= 0 {
		isFound = false
		return
	}
	isFound = true
	return
}

func DeleteTagAndArticle(tagId int, userId int) (ok bool) {
	tx := DB.Begin()
	if tx.Where("id = ? AND user_id = ?", tagId, userId).Delete(&ArticleTag{}).RowsAffected != 1 {
		tx.Rollback()
		ok = false
		return
	}
	if tx.Where("tag_id = ?", tagId).Delete(&Article{}).RowsAffected != 1 {
		tx.Rollback()
		ok = false
		return
	}

	tx.Commit()
	ok = true
	return
}
