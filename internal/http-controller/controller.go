package http_controller

import (
	"context"
	"github.com/labstack/echo/v4"
	"log/slog"
)

type httpUsecase interface {
	SaveURL(context.Context, string, string) error
	GetURL(context.Context, string) (string, error)
	DeleteURL(context.Context, string) error
}

type HttpController struct {
	e  *echo.Echo
	uc httpUsecase
	l  *slog.Logger
}

func New(ec *echo.Echo, uc httpUsecase, log *slog.Logger) *HttpController {
	return &HttpController{
		e:  ec,
		uc: uc,
		l:  log,
	}
}
