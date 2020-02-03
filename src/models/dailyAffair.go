package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type dailyAffair struct {
	gorm.Model
	Title string
	Extra string
	//Owner string
	UserId uint
}

type dailyOutput struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Extra     string `json:"extra"`
	CreatedAt string `json:"created_at"`
}

func GetAllDailyAffair(userId int) interface{} {
	data := make([]*dailyAffair, 0, 100)
	DB.Table("daily_affairs").Where("user_id = ?", userId).Find(&data)

	out := make([]*dailyOutput, 0, 100)
	for _, v := range data {
		out = append(out, &dailyOutput{
			ID:        v.ID,
			Title:     v.Title,
			Extra:     v.Extra,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return out
}

func AddDailyAffair(userId int, title string, extra string) error {
	tx := DB.Begin()
	if tx.Create(&dailyAffair{
		Title: title,
		Extra: extra,
		//Owner: owner,
		UserId: uint(userId),
	}).RowsAffected != 1 {
		tx.Rollback()
		return errors.New("Can'tInsertIntoDatabase")
	}
	tx.Commit()
	return nil
}

func DeleteDailyAffair(userId int, id int) error {
	temp := new(dailyAffair)

	DB.Where("id = ? AND user_id = ?", id, userId).Find(temp)
	if temp.ID <= 0 {
		return errors.New("NotFound")
	}

	tx := DB.Begin()
	if tx.Where("id = ? AND user_id = ?", id, userId).Delete(&dailyAffair{}).RowsAffected != 1 {
		tx.Rollback()
		return errors.New("Can'tInsertIntoDatabase")
	}
	tx.Commit()
	return nil
}

func ModifyDailyAffair(userId int,id int,title string,extra string) (ok bool) {
	tx := DB.Begin()
	if tx.Model(&dailyAffair{}).Where("id = ? AND user_id = ?", id, userId).Updates(&dailyAffair{
		Title: title,
		Extra: extra,
	}).RowsAffected != 1 {
		tx.Rollback()
		ok = false
		return
	}
	tx.Commit()
	ok = true
	return
}

//找到了返回true,没找到返回false
func FindDailyAffair(userId int, id int) (isFound bool) {
	temp := new(dailyAffair)
	DB.Where("id = ? AND user_id = ?", id, userId).Find(temp)

	if temp.ID <= 0 {
		isFound = false
		return
	}

	isFound = true
	return
}
