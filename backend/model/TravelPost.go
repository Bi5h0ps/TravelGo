package model

import "time"

type TravelPost struct {
	ID          int       `json:"post_id" gorm:"primaryKey;autoIncrement;column:ID"`
	PostTitle   string    `json:"title" gorm:"title"`
	Destination string    `json:"destination" gorm:"destination"`
	StartDate   time.Time `json:"start_date" gorm:"start_date;type:date"`
	EndDate     time.Time `json:"end_date" gorm:"end_date;type:date"`
	Tags        string    `json:"tags" gorm:"tags"`
	IsDeleted   bool      `json:"-" gorm:"is_deleted"`
}
