package handler

import (
	"HW_5/internal/model"
	"HW_5/internal/usecase"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	usecase *usecase.AuthUsecase
}

func NewAuthHandler(usecase *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: usecase}
}

// registration
func (h *AuthHandler) Register(c echo.Context) error {
	var req model.RegisterRequest
	// Parse 
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ServerResponse{Status: "error", Message: "Invalid request"})
	}

	// Usecase 
	id, err := h.usecase.Register(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ServerResponse{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, model.ServerResponse{
		Status: "success",
		Data:   map[string]int{"id": id},
	})
}

// authentication
func (h *AuthHandler) Login(c echo.Context) error {
	var req model.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ServerResponse{Status: "error", Message: "Invalid request"})
	}

	token, err := h.usecase.Login(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ServerResponse{Status: "error", Message: "Invalid email or password"})
	}

	return c.JSON(http.StatusOK, model.ServerResponse{
		Status: "success",
		Data:   model.AuthResponse{Token: token},
	})
}

// returns the current user
func (h *AuthHandler) GetMe(c echo.Context) error {
	
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	
	return c.JSON(http.StatusOK, model.ServerResponse{
		Status: "success",
		Data:   claims,
	})
}
