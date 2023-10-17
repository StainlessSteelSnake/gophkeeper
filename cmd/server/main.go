package main

import (
	"context"

	"github.com/StainlessSteelSnake/gophkeeper/internal/config"
	"github.com/StainlessSteelSnake/gophkeeper/internal/storage"
)

func main() {
	ctx := context.Background()
	cfg := config.NewConfiguration()
	_ = storage.NewStorage(ctx, cfg.DatabaseURI)
}
