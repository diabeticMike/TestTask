package model

// User entity
type User struct {
	ID   uint64 `db:"id"`
	Username string `db:"username"`
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

// WholeUser for holding all user fields together 
type WholeUser struct {
	User
	UserProfile
	UserData
}
