package telegram

import (
	"context"
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/korovkin/limiter"

	"github.com/sknv/passwordless-verifier/pkg/closer"
	"github.com/sknv/passwordless-verifier/pkg/log"
)

type BotConfig struct {
	APIToken          string
	PollingTimeout    time.Duration
	MaxUpdatesAllowed int
	Debug             bool
}

type Bot struct {
	Config BotConfig

	bot   *tgbotapi.BotAPI
	limit *limiter.ConcurrencyLimiter
}

func NewBot(config BotConfig) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(config.APIToken)
	if err != nil {
		return nil, err
	}

	bot.Debug = config.Debug
	fmt.Printf("Authorized on account %s", bot.Self.UserName)

	return &Bot{
		Config: config,

		bot:   bot,
		limit: limiter.NewConcurrencyLimiter(config.MaxUpdatesAllowed),
	}, nil
}

func (b *Bot) Run(ctx context.Context) {
	updateConfig := tgbotapi.UpdateConfig{
		Offset:  0, // TODO: fetch offset from the db
		Limit:   b.Config.MaxUpdatesAllowed,
		Timeout: int(b.Config.PollingTimeout.Seconds()),
	}

	updates := b.bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		if _, err := b.limit.Execute(func() { b.handleUpdate(ctx, update) }); err != nil {
			log.Extract(ctx).WithError(err).Error("handle telegram update")
		}
	}
}

func (b *Bot) Close(ctx context.Context) error {
	b.bot.StopReceivingUpdates()

	return closer.CloseWithContext(ctx, func() error { return b.limit.WaitAndClose() })
}

func (b *Bot) handleUpdate(ctx context.Context, update tgbotapi.Update) {
	log.Extract(ctx).WithField("update", update).Info("handle update")
}
