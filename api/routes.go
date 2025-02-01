package api

import (
	"github.com/calforcal/can-lily-eat-it/internal/handlers"
	"github.com/calforcal/can-lily-eat-it/storage"
	"github.com/labstack/echo/v4"
)

type ApiRouter struct {
	e       *echo.Echo
	storage storage.DBer
}

func NewApiRouter(e *echo.Echo, s storage.DBer) *ApiRouter {
	return &ApiRouter{e: e, storage: s}
}

func (r *ApiRouter) RegisterRoutes() {
	ah := handlers.NewApiHandler(r.e, r.storage)
	auth := handlers.NewAuthHandler(r.e, r.storage)

	r.e.GET("/", ah.GetRoot)

	api := r.e.Group("/api")
	api.GET("/login", auth.Login)
	api.GET("/callback", auth.Callback)
}
