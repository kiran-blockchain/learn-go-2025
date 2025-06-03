package repositories

import (
	"database/sql"
	"errors"
	"rest-api/interfaces"
	"rest-api/models"
	"rest-api/utils"
)

type userRepo struct {
	db *sql.DB
}

func (r *userRepo) GetUserById() (models.User, error) {
	// In a real project, this would interact with the DB.
	return models.User{ID: 1, Username: "Alice"}, nil

}

func (r *userRepo) GetAllUsers() ([]models.User, error) {
	rows, err := r.db.Query("select ID,Email,Username from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Email, &u.Username)

		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func NewUserRepository(database *sql.DB) interfaces.UserRepository {
	return &userRepo{db: database}
}

func (r *userRepo) CreateUser(u models.User) (int64, error) {
	// In a real project, this would interact with the DB.
    
	result, err := r.db.Exec("insert into users (username,email,password_hash) values (?,?,?)", u.Username, u.Email, u.PasswordHash)
	if err != nil {
		return 0, err
	} else {
		id, _ := result.LastInsertId()
		return id, nil
	}
}
//ideally it should return the JWT token as return parameter
func(r *userRepo) Login (u models.User)(string,error){
		rows, err := r.db.Query("select ID, Password_Hash from users where email = ?",u.Email)
		if err != nil {
		return "", err
	}
	defer rows.Close()
	var user  models.User
	for rows.Next() {
		//var user models.User
		err := rows.Scan(&user.ID,&user.PasswordHash)

		if err != nil {
			return "", err
		}
	}
	checkPasswordWithDB := utils.CheckPasswordHash(u.PasswordHash,user.PasswordHash)	
	if !checkPasswordWithDB {
		return "",errors.New("invalid password")
	}else {
		//create JWT and send the JWT 
		jwt,err:= utils.GenerateJWT(uint(user.ID))
		if err!=nil {
             return "",errors.New("unable to generate JWT")
		}else{
			return jwt,nil
		}
	}
}
