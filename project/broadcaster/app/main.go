package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nats-io/nats.go"
)

var bot *tgbotapi.BotAPI
var chatId int64

func main() {
	nc, err := nats.Connect(os.Getenv("NATS_URL"))

	if err != nil {
		log.Fatalf("unable to connect to nats, error: %s", err.Error())
	}

	log.Printf("succesfully connected to %s\n", nc.ConnectedAddr())

	_, err = nc.QueueSubscribe(os.Getenv("NATS_SUBJECT"), os.Getenv("NATS_GROUP"), sendToTelegram)
	if err != nil {
		log.Fatalf("unable to create subscription, error: %s", err.Error())
	}
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on %s, group %s\n", os.Getenv("NATS_SUBJECT"), os.Getenv("NATS_GROUP"))

	bot, err = tgbotapi.NewBotAPI(os.Getenv("TG_BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Failed to create bot api instance. %s", err.Error())
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	if chatId, err = strconv.ParseInt(os.Getenv("TG_CHAT_ID"), 10, 64); err != nil {
		log.Fatalf("Failed to parse chat id. %s", err.Error())
	}

	runtime.Goexit()
}

func sendToTelegram(msg *nats.Msg) {
	fmt.Printf("Received message: %s", msg.Data)
	message := tgbotapi.NewMessage(chatId, string(msg.Data))
	_, err := bot.Send(message)

	if err != nil {
		log.Printf("Failed to send message to telegram chat with id %d", chatId)
	}
}
