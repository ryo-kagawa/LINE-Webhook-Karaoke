package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/domain/model"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/domain/repository"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/infrastructure/database"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/infrastructure/database/initialize"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}
}

func run() error {
	environment := GetEnvironment()
	err := environment.Validate()
	if err != nil {
		return err
	}

	args := GetArgs()

	dsn := database.GenerateDsn(
		environment.DATABASE_USER,
		environment.DATABASE_PASSWORD,
		environment.DATABASE_URL,
	)

	if args.initialize {
		err = initialize.Initialize(dsn)
		if err != nil {
			return err
		}
	}

	bot, err := messaging_api.NewMessagingApiAPI(
		environment.LINE_CHANNEL_TOKEN,
	)
	if err != nil {
		return err
	}

	db, err := database.GetDatabase(dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	karaokeSongDatabase := database.NewKaraokeSongDatabase(db)
	http.HandleFunc("/", TimerHandler(RootHandler(environment, bot, karaokeSongDatabase)))

	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		return err
	}

	return nil
}

func TimerHandler(handler http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("requestTime: %vns\n", time.Since(startTime).Nanoseconds())
	})
}

func RootHandler(environment Environment, bot *messaging_api.MessagingApiAPI, karaoke repository.KaraokeSongRepository) http.Handler {
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
			environment.LINE_CHANNEL_SECRET,
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
