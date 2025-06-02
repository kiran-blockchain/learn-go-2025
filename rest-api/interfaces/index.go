package interfaces

import(
	"rest-api/models"
)
type UserRepository interface {
	GetAllUsers() ([]models.User,error)
	GetUserById() (models.User,error)
	CreateUser(models.User)(int64,error)
}