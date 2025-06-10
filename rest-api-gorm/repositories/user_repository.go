package repositories

import (
	"errors"

	"github.com/kiran-blockchain/rest-api-gorm/interfaces"
	"github.com/kiran-blockchain/rest-api-gorm/models"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) interfaces.UserRepository {
	return &userRepo{db: database}
}

func (u *userRepo) GetAllUsers()([]models.User,error){
	var users []models.User
	u.db.Table("users").Find(&users)
	//fmt.Println(users)
	return users,errors.New("I got an error")
}

