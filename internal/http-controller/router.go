package http_controller

import (
	"fmt"
	"github.com/labstack/echo/v4"

	"net/http"
	"url-shortener/internal/config"
)

func (hc *HttpController) saveUrl(ctx echo.Context) error {
	const op = "http-controller.router.saveUrl"

	urlToSave := ctx.QueryParam("urlToSave")
	alias := ctx.QueryParam("alias")
	if urlToSave == "" || alias == "" {
		hc.l.Info("Alias or urlToSave is empty")
		return ctx.String(http.StatusInternalServerError, "alias or url to save could not be empty")
	}

	err := hc.uc.SaveURL(ctx.Request().Context(), urlToSave, alias)
	if err != nil {
		hc.l.Error(op, "error save URL by alias: ", err)
		return ctx.String(http.StatusInternalServerError, "could not to save url")
	}

	return ctx.String(http.StatusOK, "URL saved successfully")
}

func (hc *HttpController) deleteUrl(ctx echo.Context) error {
	const op = "http-controller.router.deleteUrl"

	alias := ctx.QueryParam("alias")
	if alias == "" {
		hc.l.Info("Alias is empty")
		return ctx.String(http.StatusInternalServerError, "alias could not be empty")
	}

	err := hc.uc.DeleteURL(ctx.Request().Context(), alias)
	if err != nil {
		hc.l.Error(op, "error deleting URL by alias: ", err)
		return ctx.String(http.StatusNotFound, "could not find url")
	}

	return ctx.String(http.StatusOK, "URL deleted successfully")
}

func (hc *HttpController) redirectToUrl(ctx echo.Context) error {
	const op = "http-controller.router.getUrl"

	alias := ctx.QueryParam("alias")
	if alias == "" {
		hc.l.Info("Alias is empty")
		return ctx.String(http.StatusInternalServerError, "alias could not be empty")
	}

	url, err := hc.uc.GetURL(ctx.Request().Context(), alias)
	if err != nil {
		hc.l.Error(op, "error getting URL by alias: ", err)
		return ctx.String(http.StatusNotFound, "could not find url")
	}

	return ctx.Redirect(http.StatusFound, url)
}

func (hc *HttpController) Start(c config.HTTPConfig) {
	hc.e.POST("/save", hc.saveUrl)
	hc.e.DELETE("/delete", hc.deleteUrl)
	hc.e.GET("/alias", hc.redirectToUrl)
	hc.e.Start(fmt.Sprintf("localhost:%s", c.Port))
}
