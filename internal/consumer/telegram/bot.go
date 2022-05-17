package telegram

import (
	"context"
	"fmt"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/korovkin/limiter"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"

	"github.com/sknv/passwordless-verifier/internal/usecase"
	"github.com/sknv/passwordless-verifier/pkg/closer"
	"github.com/sknv/passwordless-verifier/pkg/log"
)

type BotConfig struct {
	APIToken          string
	PollingTimeout    time.Duration
	MaxUpdatesAllowed int
	CallbackURL       string
	Debug             bool
}

type Usecase interface {
	SetVerificationChat(ctx context.Context, params *usecase.SetVerificationChatParams) error
}

type Bot struct {
	Config  BotConfig
	Usecase Usecase

	bot   *tgbotapi.BotAPI
	limit *limiter.ConcurrencyLimiter
}

func NewBot(config BotConfig, usecase Usecase) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(config.APIToken)
	if err != nil {
		return nil, err
	}

	bot.Buffer = 0 // no buffering allowed for updates channel, control concurrency manually with a limiter
	bot.Debug = config.Debug

	return &Bot{
		Config:  config,
		Usecase: usecase,

		bot:   bot,
		limit: limiter.NewConcurrencyLimiter(config.MaxUpdatesAllowed),
	}, nil
}

func (b *Bot) Run(ctx context.Context) {
	updateConfig := tgbotapi.UpdateConfig{
		Limit:   b.Config.MaxUpdatesAllowed,
		Timeout: int(b.Config.PollingTimeout.Seconds()),
	}

	updates := b.bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		if _, err := b.limit.Execute(func() { b.HandleUpdate(ctx, update) }); err != nil {
			log.Extract(ctx).WithError(err).Error("execute telegram update handler")
		}
	}
}

func (b *Bot) Close(ctx context.Context) error {
	b.bot.StopReceivingUpdates()

	return closer.CloseWithContext(ctx, func() error { return b.limit.WaitAndClose() })
}

func (b *Bot) HandleUpdate(ctx context.Context, update tgbotapi.Update) {
	if update.Message == nil { // ignore empty messages
		return
	}

	ctx, span := otel.Tracer("").Start(ctx, "telegram.HandleUpdate")
	defer span.End()

	logger := log.Extract(ctx).WithFields(logrus.Fields{
		"chat": update.Message.Chat,
		"text": update.Message.Text,
	})

	var err error
	switch { // route commands
	case strings.Contains(update.Message.Text, "/start"): // verification started
		err = b.startVerification(ctx, update.Message)
		if err != nil {
			err = fmt.Errorf("start verification: %w", err)
		}
	case update.Message.Contact != nil: // contact shared
		err = b.verifyContact(ctx, update.Message)
		if err != nil {
			err = fmt.Errorf("verify contact: %w", err)
		}
	default:
		err = b.unknownCommand(update.Message)
		if err != nil {
			err = fmt.Errorf("unknown command: %w", err)
		}
	}

	if err != nil {
		logger.WithError(err).Error("handle telegram update")
	} else {
		logger.Info("telegram update handled")
	}
}

func (b *Bot) unknownCommand(message *tgbotapi.Message) error {
	return b.reply(message, fmt.Sprintf(msgFormatUnknownCommand, message.Text))
}

func (b *Bot) reply(to *tgbotapi.Message, text string) error {
	msg := tgbotapi.NewMessage(to.Chat.ID, text)

	_, err := b.bot.Send(msg)
	return err
}
