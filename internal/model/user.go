package model


type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}


type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}


type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}


type AuthResponse struct {
	Token string `json:"token"`
}
