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

// Register handles user registration
// @Summary Register a new user
// @Tags Auth
// @Description Register a new user with username, email and password
// @Accept json
// @Produce json
// @Param input body model.RegisterRequest true "Register Input"
// @Success 201 {object} model.ServerResponse
// @Failure 400 {object} model.ServerResponse
// @Router /auth/register [post]
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

// Login handles user authentication
// @Summary Login user
// @Tags Auth
// @Description Login with email and password to get JWT token
// @Accept json
// @Produce json
// @Param input body model.LoginRequest true "Login Input"
// @Success 200 {object} model.ServerResponse
// @Failure 400 {object} model.ServerResponse
// @Router /auth/login [post]
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

// GetMe returns the current user's info based on the JWT token
// @Summary Get current user info
// @Tags Users
// @Description Get details of the currently logged-in user
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} model.ServerResponse
// @Router /users/me [get]
func (h *AuthHandler) GetMe(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return c.JSON(http.StatusOK, model.ServerResponse{
		Status: "success",
		Data:   claims,
	})
}
