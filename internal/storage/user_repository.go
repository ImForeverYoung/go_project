package storage

import (
	"HW_5/internal/model"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

// new user
func (s *Storage) CreateUser(ctx context.Context, user model.User) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `INSERT INTO users (username, email, password_hash)
			  VALUES ($1, $2, $3) RETURNING id`

	var id int
	err := s.conn.QueryRow(ctx, query, user.Username, user.Email, user.PasswordHash).Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "CreateUser failed: %v\n", err)
		return 0, err
	}

	return id, nil
}


func (s *Storage) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `SELECT id, username, email, password_hash FROM users WHERE email=$1`

	var user model.User
	err := s.conn.QueryRow(ctx, query, email).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.User{}, fmt.Errorf("user not found")
		}
		fmt.Fprintf(os.Stderr, "GetUserByEmail failed: %v\n", err)
		return model.User{}, err
	}

	return user, nil
}
