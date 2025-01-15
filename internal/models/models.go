package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Email    string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password string `gorm:"type:varchar(100);not null"`
	Phrases  []Phrase
}

type Phrase struct {
	gorm.Model
	OriginalPhrase  string `gorm:"type:text;not null" json:"original_phrase"`
	EmojifiedPhrase string `gorm:"type:text;not null" json:"emojified_phrase"`
	UserID          uint   `gorm:"not null" json:"user_id"`
	User            User   `json:"-" gorm:"foreignKey:UserID"`
}
