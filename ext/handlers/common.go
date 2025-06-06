package handlers

import (
	"github.com/EvgTG/gotgbot/v2"
	"github.com/EvgTG/gotgbot/v2/ext"
)

type Response func(b *gotgbot.Bot, ctx *ext.Context) error
