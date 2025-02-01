package handlers

import (
	"net/http"

	"github.com/calforcal/can-lily-eat-it/storage"
	"github.com/labstack/echo/v4"
)

type ApiHandler struct {
	e       *echo.Echo
	storage storage.DBer
}

func NewApiHandler(e *echo.Echo, s storage.DBer) *ApiHandler {
	return &ApiHandler{e: e, storage: s}
}

func (h *ApiHandler) GetRoot(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
}

// func (h *ApiHandler) GetUser(c echo.Context) error {
// 	userID := c.Param("id")
// 	userIDInt, err := strconv.Atoi(userID)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
// 	}
// 	user, err := h.storage.GetUser(userIDInt)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 	}
// 	return c.JSON(http.StatusOK, user)
// }
