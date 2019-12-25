package libebot

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

func Omubot(w http.ResponseWriter, req *http.Request) {
	bot, err := linebot.New(
		"CHANNEL_SECRET",
		"CHANNEL_TOKEN",
	)
	if err != nil {
		log.Fatal(err)
	}

	events, err := bot.ParseRequest(req)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
	w.WriteHeader(200)
}
