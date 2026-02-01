package usecase

import (
	"HW_5/internal/model"
	"HW_5/internal/storage"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	repo *storage.Storage
}

func NewAuthUsecase(repo *storage.Storage) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}


func (u *AuthUsecase) Register(req model.RegisterRequest) (int, error) {
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	// create user model
	user := model.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	// save 
	return u.repo.CreateUser(context.Background(), user) 
}

// verify + generate jwt
func (u *AuthUsecase) Login(req model.LoginRequest) (string, error) {
	
	user, err := u.repo.GetUserByEmail(context.Background(), req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), //72 hour
	})

	
	tokenString, err := token.SignedString([]byte("my_secret_key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
