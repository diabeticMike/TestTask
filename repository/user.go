package repository

import (
	"database/sql"

	"github.com/TestTask/model"
)

// NewUserRepo return UserRepository instance
func NewUserRepo(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

// UserRepository for working with user, user_data, user_profile tables
type UserRepository interface {
	GetUserByUsername(username string) (*model.User, error)
	GetUserProfileByID(userID uint64) (*model.UserProfile, error)
	GetUserDataByID(userID uint64) (*model.UserData, error)
}

type userRepo struct {
	db *sql.DB
}

func (ur *userRepo) GetUserByUsername(username string) (*model.User, error) {
	row := ur.db.QueryRow("SELECT * FROM user WHERE username=? LIMIT 1;", username)
	user := model.User{}
	if err := row.Scan(&user.ID, &user.Username); err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepo) GetUserProfileByID(userID uint64) (*model.UserProfile, error) {
	row := ur.db.QueryRow("SELECT * FROM user_profile WHERE user_id=? LIMIT 1;", userID)
	userProfile := model.UserProfile{}
	if err := row.Scan(&userProfile.UserID, &userProfile.FirstName, &userProfile.LastName,
		&userProfile.Phone, &userProfile.Address, &userProfile.City); err != nil {
		return nil, err
	}
	return &userProfile, nil
}

func (ur *userRepo) GetUserDataByID(userID uint64) (*model.UserData, error) {
	row := ur.db.QueryRow("SELECT * FROM user_data WHERE user_id=? LIMIT 1;", userID)
	userData := model.UserData{}
	if err := row.Scan(&userData.UserID, &userData.School); err != nil {
		return nil, err
	}
	return &userData, nil
}
