package handler

import (
	"chansTask/internal/app/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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
	ch := c.QueryParam("channels")
	ran := c.QueryParam("randomin")
	channelsCount, chErr := strconv.Atoi(ch)
	if chErr!=nil{
		log.Infof("bad request")
		return echo.NewHTTPError(http.StatusBadRequest, "cannot parse param")
	}
	randomIn, ranErr := strconv.Atoi(ran)
	if ranErr!=nil{
		log.Infof("bad request")
		return echo.NewHTTPError(http.StatusBadRequest, "cannot parse param")
	}
	res := h.service.FromNChannelsToOneChannel(channelsCount, randomIn)
	return c.JSON(http.StatusOK, res)
}


