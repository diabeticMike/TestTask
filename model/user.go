package model

// User entity
type User struct {
	ID   uint64 `db:"id"`
	Name string `db:"username"`
}

// UserProfile holds user entity additional info
type UserProfile struct {
	UserID    uint64 `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Phone     string `db:"phone"`
	Address   string `db:"address"`
	City      string `db:"city"`
}

// UserData holds user school
type UserData struct {
	UserID uint64 `db:"user_id"`
	School string `db:"school"`
}
