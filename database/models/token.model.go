package models

import (
	"blog/core/config"
	"blog/pkg/helpers"
	"time"

	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type Token struct {
	ID           uint   `gorm:"primarykey"`
	UserID       uint   `gorm:"not null"`
	AccessToken  string `gorm:"type:varchar(64);unique;not null"`
	RefreshToken string `gorm:"type:varchar(64);unique;not null"`
	Invoked      bool   `gorm:"default:false"`

	User User `gorm:"foreignkey:UserID"`

	AccessExpiresAt  time.Time `gorm:"not null"`
	RefreshExpiresAt time.Time `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Token) BeforeCreate(tx *gorm.DB) (err error) {
	t.AccessToken = helpers.SHA1(t.AccessToken)
	t.RefreshToken = helpers.SHA1(t.RefreshToken)

	accessExpiresAt, refreshExpiresAt := GetTokenAccessExpires()

	t.AccessExpiresAt = accessExpiresAt
	t.RefreshExpiresAt = refreshExpiresAt

	return
}

func GetTokenAccessExpires() (accessExpiresAt, refreshExpiresAt time.Time) {
	access, refresh := config.GetTokensExpires()

	accessExpiresAt = carbon.Now().AddHours(access).ToStdTime()
	refreshExpiresAt = carbon.Now().AddDays(refresh).ToStdTime()

	return
}
