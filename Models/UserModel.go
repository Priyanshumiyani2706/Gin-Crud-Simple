package Model

//User model
type User struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// Declares table name to connect and performe databse query
func (b *User) TableName() string {
	return "users"
}
