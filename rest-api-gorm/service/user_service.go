package services

import (
	"github.com/kiran-blockchain/rest-api-gorm/interfaces"
	"github.com/kiran-blockchain/rest-api-gorm/models"
)


 

type UserService interface {
    GetAllUsers() ([]models.User,error)
}

type userService struct {
    repo interfaces.UserRepository
}

//pass the repository dependency 
func NewUserService(r interfaces.UserRepository) UserService {
    return &userService{repo: r}
}

func (s *userService) GetAllUsers() ([]models.User,error) {
    return s.repo.GetAllUsers()
}