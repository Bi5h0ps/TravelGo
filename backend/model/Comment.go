package model

import "time"

type Comment struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement;column:ID"`
	PostId    int       `json:"post_id" grom:"post_id"`
	Username  string    `json:"username" gorm:"username"`
	Content   string    `json:"content" gorm:"content"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	IsDeleted bool      `json:"is_deleted" gorm:"is_deleted"`
}
