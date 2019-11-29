package main

import (
	"fmt"
	"log"
	"main/oxfordApi"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func indexPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func sendToUser(replyToken string, definition string) {
	message := linebot.NewTextMessage(definition)
	_, err := bot.ReplyMessage(replyToken, message).Do()
	if err != nil {
		panic(err)
	}

}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	log.Printf("events: %#v\nerr: %v", events, err)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		log.Printf("%+v\n", event)
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				// if it is an English word, search the dictionary
				definition := oxfordApi.GetDefinition(message.Text)
				log.Printf(definition)

				sendToUser(event.ReplyToken, definition)
			}
		}
	}
}

func main() {
	ChannelSecret := os.Getenv("CHANNEL_SECRET")
	ChannelAccessToken := os.Getenv("CHANNEL_ACCESS_TOKEN")

	bot, err := linebot.New(ChannelSecret, ChannelAccessToken)
	log.Println("Bot:", bot, " err:", err)

	http.HandleFunc("/", indexPage)
	http.HandleFunc("/callback", callbackHandler)

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	serverErr := http.ListenAndServe(addr, nil)
	if serverErr != nil {
		panic(serverErr)
	}

}
