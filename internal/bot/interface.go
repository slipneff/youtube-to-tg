package bot

import (
	"slipneff/youtube-to-tg/internal/script"

	"github.com/mymmrac/telego"
)

type Bot struct {
	bot    *telego.Bot
	script *script.Script
}

func New(bot *telego.Bot, script *script.Script) *Bot {
	return &Bot{bot: bot, script: script}
}
