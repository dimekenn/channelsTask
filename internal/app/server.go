package app

import (
	"chansTask/internal/app/configs"
	"chansTask/internal/app/handler"
	"chansTask/internal/app/service"
	"context"
	"github.com/labstack/echo/v4"
)

//server for REST apis
func StartHTTPServer(ctx context.Context, errCh chan<- error, cfg *configs.Configs) {
	app := echo.New()

	chService := service.NewService()
	chHandler := handler.NewHandler(chService)

	app.GET("/api/v1/animal", chHandler.FromNChannelsToOneChannel)

	errCh <- app.Start(cfg.Port)
}