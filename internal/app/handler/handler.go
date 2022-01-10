package handler

import (
	"chansTask/internal/app/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) FromNChannelsToOneChannel(c echo.Context) error {
	x := c.QueryParam("x")
	y := c.QueryParam("y")
	xx, _ := strconv.Atoi(x)
	yy, _ := strconv.Atoi(y)
	res := h.service.FromNChannelsToOneChannel(xx, yy)
	return c.JSON(http.StatusOK, res)
}


