package main

import (
	"context"
	"slipneff/youtube-to-tg/internal/di"
	"slipneff/youtube-to-tg/pkg/utils/config"
	"slipneff/youtube-to-tg/pkg/utils/flags"
)

func main() {
	flags := flags.MustParseFlags()
	config := config.MustLoadConfig(flags.ConfigPath)
	container := di.New(context.Background(), config)
	container.Bot()
}
