package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/domain/model"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/domain/repository"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/infrastructure/database"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/rootcommand/environment"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	result, err := handler(w, r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	fmt.Println(result)
	w.WriteHeader(200)
}

func handler(w http.ResponseWriter, r *http.Request) (string, error) {
	environment := environment.GetEnvironment()
	if err := environment.Validate(); err != nil {
		return "", err
	}

	bot, err := messaging_api.NewMessagingApiAPI(
		environment.Line.LINE_CHANNEL_TOKEN,
	)
	if err != nil {
		return "", err
	}

	db, err := NewDatabase(environment)
	if err != nil {
		return "", err
	}
	RootHandler(environment, bot, db).ServeHTTP(w, r)

	return "finish", nil
}

func RootHandler(environment environment.Environment, bot *messaging_api.MessagingApiAPI, karaoke repository.KaraokeSongRepository) http.Handler {
	RandomPickKaraokeSong := func(text string) ([]model.KaraokeSong, error) {
		switch text {
		case "DAM":
			return karaoke.Dam()
		case "JOYSOUND":
			return karaoke.Joysound()
		default:
			return karaoke.Ramdom()
		}
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cb, err := webhook.ParseRequest(
			environment.Line.LINE_CHANNEL_SECRET,
			r,
		)
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(cb.Events) == 0 {
			return
		}

		for _, event := range cb.Events {
			switch e := event.(type) {
			case webhook.MessageEvent:
				switch message := e.Message.(type) {
				case webhook.TextMessageContent:
					karaokeSongList, err := RandomPickKaraokeSong(message.Text)
					if err != nil {
						fmt.Println(err)
					}
					lineMessages := make([]messaging_api.MessageInterface, 0, len(karaokeSongList))
					for _, karaokeSong := range karaokeSongList {
						lineMessages = append(lineMessages, karaokeSong.GenerateLineTextMessage())
					}
					bot.ReplyMessage(
						&messaging_api.ReplyMessageRequest{
							ReplyToken: e.ReplyToken,
							Messages:   lineMessages,
						},
					)
				}
			}
		}
	})
}

func NewDatabase(environment environment.Environment) (database.Database, error) {
	port, err := strconv.Atoi(environment.Database.DATABASE_POSTGRESQL_PORT)
	if err != nil {
		return database.Database{}, err
	}
	return database.NewDatabase(
		environment.Database.DATABASE_POSTGRESQL_HOST,
		uint16(port),
		environment.Database.DATABASE_POSTGRESQL_USER,
		environment.Database.DATABASE_POSTGRESQL_PASSWORD,
		environment.Database.DATABASE_POSTGRESQL_DATABASE,
		environment.Database.DATABASE_POSTGRESQL_SSLMODE,
		environment.Database.DATABASE_POSTGRESQL_SCHEMA,
	)
}
