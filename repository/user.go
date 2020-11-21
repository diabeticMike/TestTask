package repository

import (
	"database/sql"

	"github.com/TestTask/model"
)

// UserRepository for working with user, user_data, user_profile tables
type UserRepository interface {
	GetUserByName(username string) (*model.User, error)
	GetUserProfileByID(userID uint) (*model.UserProfile, error)
	GetUserDataByID(userID uint) (*model.UserData, error)
}

type userRepo struct {
	db *sql.DB
}

func (ur *userRepo) GetUserByName(username string) (*model.User, error) {
	row := ur.db.QueryRow("SELECT * FROM user WHERE username=$1 LIMIT 1;", username)
	user := model.User{}
	if err := row.Scan(&user.ID, &user.Name); err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepo) GetUserProfileByID(userID uint) (*model.UserProfile, error) {
	row := ur.db.QueryRow("SELECT * FROM user_profile WHERE id=$1 LIMIT 1;", userID)
	userProfile := model.UserProfile{}
	if err := row.Scan(&userProfile.UserID, &userProfile.FirstName, &userProfile.LastName,
		&userProfile.Phone, &userProfile.Address, &userProfile.City); err != nil {
		return nil, err
	}
	return &userProfile, nil
}

func (ur *userRepo) GetUserDataByID(userID uint) (*model.UserData, error) {
	row := ur.db.QueryRow("SELECT * FROM user_data WHERE id=$1 LIMIT 1;", userID)
	userData := model.UserData{}
	if err := row.Scan(&userData.UserID, &userData.School); err != nil {
		return nil, err
	}
	return &userData, nil
}
