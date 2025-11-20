package model

import "time"

type Question struct {
	ID         int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Text       string `gorm:"not null" json:"text"`
	Answers []Answer `gorm:"foreignKey:QuestionID"`
	CreatedAt time.Time `json:"created_at"`
}