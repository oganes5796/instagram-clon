package main

import (
	"context"

	"github.com/oganes5796/instagram-clon/internal/app"
	"github.com/oganes5796/instagram-clon/internal/logger"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		logger.Fatal("failed to init app:", zap.Error(err))
	}

	err = a.Run()
	if err != nil {
		logger.Fatal("failed to run app:", zap.Error(err))
	}
}
