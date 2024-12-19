package main

import (
	"context"
	"log"

	"github.com/aziret/s3-mini-storage/internal/app"
)

func main() {
	ctx := context.Background()
	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	_ = a
}
