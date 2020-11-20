package model

// Auth struct for authorization
type Auth struct {
	ID     uint64 `db:"id"`
	APIKey string `db:"auth-key"`
}
