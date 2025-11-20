package model

import (
	"time"

	"github.com/google/uuid"
)

type Answer struct {
	ID         int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	QuestionID int64     `gorm:"not null" json:"question_id"`
	UserID     uuid.UUID `gorm:"not null;default:gen_random_uuid()" json:"user_id"`
	Text       string    `gorm:"not null" json:"text"`
	CreatedAt  time.Time `json:"created_at"`

	Question Question `gorm:"foreignKey:QuestionID;constraint:OnDelete:CASCADE" json:"-"`
}
