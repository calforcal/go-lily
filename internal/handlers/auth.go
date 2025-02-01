package handlers

import (
	"net/http"
	"time"

	"github.com/calforcal/can-lily-eat-it/services/google"
	"github.com/calforcal/can-lily-eat-it/storage"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	e       *echo.Echo
	storage storage.DBer
	google  *google.GoogleService
}

func NewAuthHandler(e *echo.Echo, s storage.DBer) *AuthHandler {
	return &AuthHandler{
		e:       e,
		storage: s,
		google:  google.NewGoogleService(),
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	URL string `json:"url"`
}

type UserResponse struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (h *AuthHandler) Login(c echo.Context) error {
	authURL := h.google.GetAuthURL()
	return c.JSON(http.StatusOK, LoginResponse{URL: authURL})
}

func (h *AuthHandler) Callback(c echo.Context) error {
	code := c.QueryParam("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Code not found"})
	}

	userInfo, err := h.google.GetUserInfo(code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	user, err := h.storage.GetOrCreateUser(userInfo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, UserResponse{
		UUID:      user.UUID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}
