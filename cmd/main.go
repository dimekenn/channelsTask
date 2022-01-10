package main

import (
	"chansTask/internal/app"
	"chansTask/internal/app/configs"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

//go:embed configs.json
var fs embed.FS

const configName = "configs.json"

func main()  {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//reading json file for configs
	data, readErr := fs.ReadFile(configName)
	if readErr != nil {
		log.Fatal(readErr)
	}

	//creating config entity to deserialize configs.json
	cfg := configs.NewConfig()
	if unmErr := json.Unmarshal(data, &cfg); unmErr != nil {
		log.Fatal(unmErr)
	}

	//channel for errors
	errCh := make(chan error, 1)

	go func(ctx context.Context, errCh chan error) {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		errCh <- fmt.Errorf("%v", <-sigCh)
	}(ctx, errCh)

	//new goroutine for REST api server
	go app.StartHTTPServer(ctx, errCh, cfg)

	log.Fatalf("%v", <-errCh)
}
