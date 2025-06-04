package interfaces

import "github.com/kiran-blockchain/rest-api-gorm/models"

type UserRepository interface {
	GetAllUsers() ([]models.User,error)
}