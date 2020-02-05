package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type affair struct {
	gorm.Model
	Title    string
	Deadline time.Time
	Extra    string
	//Owner string
	UserId uint
}

type outputAffair struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Deadline  string `json:"deadline"`
	Extra     string `json:"extra"`
	CreatedAt string `json:"created_at"`
}

func GetAllAffairs(userId int) interface{} {
	data := make([]*affair, 0, 100)
	DB.Table("affairs").Where("user_id = ?", userId).Find(&data)

	out := make([]*outputAffair, 0, 100)
	for _, v := range data {
		out = append(out, &outputAffair{
			ID:        v.ID,
			Title:     v.Title,
			Deadline:  v.Deadline.Format("2006-01-02 15:04:05"),
			Extra:     v.Extra,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return out
}

func AddAffair(userId int, title string, extra string, deadline time.Time) (ok bool) {
	tx := DB.Begin()
	if tx.Create(&affair{
		Title:    title,
		Deadline: deadline,
		Extra:    extra,
		//Owner:    owner,
		UserId: uint(userId),
	}).RowsAffected != 1 {
		tx.Rollback()
		ok = false
		return
	}
	tx.Commit()
	ok = true
	return
}

func DeleteAffair(userId int, id int) error {
	temp := new(affair)
	DB.Where("id = ? AND user_id = ?", id, userId).Find(temp)
	if temp.ID <= 0 {
		return errors.New("NotFound")
	}

	tx := DB.Begin()
	if tx.Where("id = ? AND user_id = ?", id, userId).Delete(&affair{}).RowsAffected != 1 {
		tx.Rollback()
		return errors.New("Can'tInsertIntoDatabase")
	}
	tx.Commit()
	return nil
}

func ModifyAffair(userId int, id int, title string, extra string, deadline time.Time) (ok bool) {
	tx := DB.Begin()
	if tx.Model(&affair{}).Where("id = ? AND user_id = ?", id, userId).Updates(&affair{
		Title:    title,
		Deadline: deadline,
		Extra:    extra,
	}).RowsAffected != 1 {
		tx.Rollback()
		ok = false
		return
	}
	tx.Commit()
	ok = true
	return
}

func FindAffair(userId int, id int) (isFound bool) {
	temp := new(affair)
	DB.Where("id = ? AND user_id = ?", id, userId).Find(temp)
	if temp.ID <= 0 {
		isFound = false
		return
	}
	isFound = true
	return
}
