package repository

import (
	"database/sql"

	"github.com/TestTask/model"
)

// NewAuthRepo return AuthRepository instance
func NewAuthRepo(db *sql.DB) AuthRepository {
	return &authRepo{db: db}
}

// AuthRepository interface for getting auth info from db
type AuthRepository interface {
	GetAuthByAPIKey(key string) (*model.Auth, error)
}

// AuthRepository realization using mysql
type authRepo struct {
	db *sql.DB
}

// GetAuthByAPIKey return all auth entity by api-key
func (ar *authRepo) GetAuthByAPIKey(key string) (*model.Auth, error) {
	row := ar.db.QueryRow("SELECT * FROM auth WHERE api_key=?;", key)
	auth := model.Auth{}
	if err := row.Scan(&auth.ID, &auth.APIKey); err != nil {
		return nil, err
	}
	return &auth, nil
}
