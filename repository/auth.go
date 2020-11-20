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
	GetAuthByAPIKey(key string) (model.Auth, error)
}

// AuthRepository realization using mysql
type authRepo struct {
	db *sql.DB
}

// GetAuthByAPIKey return all auth entity by api-key
func (ar *authRepo) GetAuthByAPIKey(key string) (model.Auth, error) {
	row := ar.db.QueryRow("SELECT * FROM auth WHERE api-key=$1 LIMIT 1;", key)
	auth := model.Auth{}
	if err := row.Scan(&key, &auth.APIKey); err != nil {
		return auth, err
	}
	return auth, nil
}
