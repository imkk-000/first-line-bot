package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/callback", func(responseWriter http.ResponseWriter, request *http.Request) {
		events, err := bot.ParseRequest(request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				responseWriter.WriteHeader(400)
			} else {
				responseWriter.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			log.Printf("event: %v\n", event)
			if event.Type == linebot.EventTypeMessage {
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(time.Now().String())).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	})
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
