package models

import "github.com/jinzhu/gorm"

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
	DB.Table("daily_events").Where("user_id = ?", userId).Find(&data)

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
