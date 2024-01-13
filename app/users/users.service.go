package users

import (
	"blog/core"
	"blog/database/models"
	"blog/database/repositories"
)

type UsersServiceInterface interface {
	FindAll() []models.User
	FindOne(id string) (models.User, core.Error)
}

type UsersService struct {
	repository repositories.UserRepositoryInterface
}

func NewUsersService() *UsersService {
	return &UsersService{
		repository: repositories.NewUserRepository(),
	}
}

func (self *UsersService) FindAll() []models.User {
	return self.repository.FindAll()
}

func (self *UsersService) FindOne(id string) (models.User, core.Error) {
	user := self.repository.FindByID(id)

	if user.ID == 0 {
		return user, core.Error{"reason": "User not found"}
	}

	return user, nil
}
