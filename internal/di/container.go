package di

import (
	"context"
	"slipneff/youtube-to-tg/internal/bot"
	"slipneff/youtube-to-tg/internal/script"
	"slipneff/youtube-to-tg/pkg/utils/config"

	"github.com/mymmrac/telego"
)

type Container struct {
	cfg *config.Config
	ctx context.Context

	telebot *telego.Bot
	script  *script.Script
	bot     *bot.Bot
}

func New(ctx context.Context, cfg *config.Config) *Container {
	return &Container{cfg: cfg, ctx: ctx}
}

func (c *Container) Bot() *telego.Bot {
	if c.telebot == nil {
		bot, err := telego.NewBot(c.cfg.BotToken, telego.WithDefaultDebugLogger())
		if err != nil {
			panic(err)
		}
		c.telebot = bot
	}

	return c.telebot
}

func (c *Container) NewBot() *bot.Bot {
	if c.bot == nil {
		c.bot = bot.New(c.Bot(), c.NewScript())
	}

	return c.bot
}

func (c *Container) NewScript() *script.Script {
	if c.script == nil {
		c.script = script.New(c.cfg)
	}

	return c.script
}
