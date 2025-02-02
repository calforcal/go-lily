package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/calforcal/can-lily-eat-it/services/google"
	"github.com/calforcal/can-lily-eat-it/services/google/auth"
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

type TokenResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	Expiry       time.Time `json:"expiry"`
	ExpiresIn    int       `json:"expires_in,omitempty"`
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

	// Get state parameter and validate it (TODO: implement state validation)
	state := c.QueryParam("state")
	if state == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "State parameter missing"})
	}

	token, err := h.google.GetToken(code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: fmt.Sprintf("Failed to exchange code: %v", err)})
	}

	userInfo, err := h.google.GetUserInfo(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: fmt.Sprintf("Failed to get user info: %v", err)})
	}

	user, err := h.storage.GetOrCreateUser(userInfo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: fmt.Sprintf("Failed to process user: %v", err)})
	}

	tokenResponse, err := auth.IssueJwt(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: fmt.Sprintf("Failed to issue JWT: %v", err)})
	}

	return c.JSON(http.StatusOK, tokenResponse)
}
