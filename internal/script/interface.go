package script

import (
	"slipneff/youtube-to-tg/pkg/utils/config"
)

type Script struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Script {
	return &Script{
		cfg: cfg,
	}
}
