package svc

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"go.uber.org/zap"
)

// global b function interface
var b *gotgbot.Bot

type tgbotService struct {
	bot *gotgbot.Bot
}

var TgbotService = new(tgbotService)

// This bot demonstrates some example interactions with telegram callback queries.
// It has a basic start command which contains a button. When pressed, this button is edited via a callback query.
func (t *tgbotService) TgCallbackInit() *tgbotService {
	// Get token from the environment variable
	// token := os.Getenv("token")
	// token := viper.GetString("tgbot.token")
	token := "737275:AAEcYBmJ6W0M"
	if token == "" {
		panic("TOKEN is empty")
	}
	var err error
	// Create bot from environment value.
	b, err = gotgbot.NewBot(token, nil)
	if err != nil {
		zap.L().Error("failed to create new bot:", zap.Error(err))
		// panic("failed to create new bot: " + err.Error())
	}
	t.bot = b
	return t
}

// 发送消息到tg
// chatid是群或者聊天对话的id，通过以下方式获取
// get user id @userinfobot /start
// get group id  https://api.telegram.org/bot[token,不包括括号]/getUpdates
func (t *tgbotService) SendMessageAlone(chatId int64, text string) error {
	_, err := b.SendMessage(chatId, text, &gotgbot.SendMessageOpts{})
	if err != nil {
		zap.L().Fatal("failed to send start message: %s\n", zap.Error(err))
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}

func (t *tgbotService) SendMessage(ctx *ext.Context, chatId int64, text string) error {
	_, err := ctx.EffectiveMessage.Reply(b, text, &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}
