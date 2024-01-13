package repositories

import (
	"blog/database"
	"blog/database/models"
	"blog/pkg/helpers"
	"sync"
	"time"

	"gorm.io/gorm"
)

type TokenRepositoryInterface interface {
	Connection() *gorm.DB
	Clearing(userId uint, wg *sync.WaitGroup)
	Create(userId uint) (string, string, string)
	FindByAccess(access string) models.Token
	DeleteByAccess(access string)
	FindByRefresh(refresh string) models.Token
	FindByRefreshAndAccess(access, refresh string) models.Token
}

type TokenRepository struct {
	DB *gorm.DB
}

func NewTokenRepository() *TokenRepository {
	return &TokenRepository{
		DB: database.Connection(),
	}
}

func (self *TokenRepository) Connection() *gorm.DB {
	return self.DB
}

func (self *TokenRepository) Clearing(userId uint, wg *sync.WaitGroup) {
	defer wg.Done()
	self.DB.
		Where("refresh_expires_at < ?", time.Now()).
		Where("user_id = ?", userId).
		Delete(&models.Token{})
}

func (self *TokenRepository) Create(userId uint) (string, string, string) {
	var token models.Token

	accessTokenUUID, _ := helpers.UUID()
	refreshTokenUUID, _ := helpers.UUID()

	token.UserID = userId
	token.AccessToken = accessTokenUUID
	token.RefreshToken = refreshTokenUUID
	token.Invoked = false

	self.DB.Create(&token).Scan(&token)

	return accessTokenUUID, refreshTokenUUID, token.AccessExpiresAt.Format(time.RFC3339)
}

func (self *TokenRepository) FindByAccess(access string) models.Token {
	var token models.Token

	self.DB.
		Joins("User").
		Where(&models.Token{AccessToken: helpers.SHA1(access)}, "access_token", "invoked").
		Where("access_expires_at > ?", time.Now()).
		First(&token)

	return token
}

func (self *TokenRepository) DeleteByAccess(access string) {
	self.DB.
		Where(&models.Token{AccessToken: access}, "access_token").
		Delete(&models.Token{})
}

func (self *TokenRepository) FindByRefresh(refresh string) models.Token {
	var token models.Token
	self.DB.
		Joins("User").
		Where(&models.Token{RefreshToken: helpers.SHA1(refresh)}, "refresh_token", "invoked").
		Where("refresh_expires_at > ?", time.Now()).
		First(&token)
	return token
}

func (self *TokenRepository) FindByRefreshAndAccess(access, refresh string) models.Token {
	var token models.Token
	self.DB.
		Joins("User").
		Where(&models.Token{AccessToken: helpers.SHA1(access), RefreshToken: helpers.SHA1(refresh)}, "access_token", "refresh_token", "invoked").
		Where("refresh_expires_at > ?", time.Now()).
		First(&token)
	return token
}
