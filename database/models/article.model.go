package models

import "time"

type Article struct {
	ID      uint   `gorm:"primarykey"`
	UserID  uint   `gorm:"not null"`
	Slug    string `gorm:"type:varchar(191);unique;not null"`
	Title   string `gorm:"type:varchar(191);not null"`
	Content string `gorm:"type:text;not null"`

	User User `gorm:"foreignkey:UserID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
