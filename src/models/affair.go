package models

import (
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

type output struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Deadline  string `json:"deadline"`
	Extra     string `json:"extra"`
	CreatedAt string `json:"created_at"`
}

func GetAllAffairs(userId int) interface{} {
	data := make([]*affair, 0, 100)
	DB.Table("affairs").Where("user_id = ?", userId).Find(&data)

	out := make([]*output, 0, 100)
	for _, v := range data {
		out = append(out, &output{
			ID:        v.ID,
			Title:     v.Title,
			Deadline:  v.Deadline.Format("2006-01-02 15:04:05"),
			Extra:     v.Extra,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return out
}


