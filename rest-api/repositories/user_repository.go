package repositories

import (
	"database/sql"
	"rest-api/interfaces"
	"rest-api/models"
)

type userRepo struct{
	db *sql.DB
}

func (r *userRepo) GetUserById() (models.User,error) {
	// In a real project, this would interact with the DB.
	return models.User{ID: 1, Username: "Alice"},nil
	
}

func (r *userRepo) GetAllUsers()([]models.User,error){
	rows,err:= r.db.Query("select ID,Email,Username from users")
	if(err!=nil){
		return nil,err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next(){
		var u models.User
		  err:= rows.Scan(&u.ID,&u.Email,&u.Username)
		  
		 if err!=nil {
			return nil,err
		 }
		 users = append(users, u)
	}
	return users,nil
}

func NewUserRepository(database *sql.DB) interfaces.UserRepository {
	return &userRepo{db: database}
}


