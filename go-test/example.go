package main

import (
	"fmt"
	"github.com/go-snh/tg"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

var strict bool = false

func fatal(err error) {
	if strict {
		log.Fatal(err)
	}
}

func main() {
	println(fmt.Sprint("Starting Program..."))
	bot1 := new(tg.TelegramBot)

	token, _ := ioutil.ReadFile("go-test/MyBot.token")
	if !bot1.SetToken(token) {
		log.Fatal("Unable to use bot token.\nMake sure you have saved correct token in MyBot.token file")
	}
	bot_user, err := bot1.GetMe()
	if err != nil { log.Fatal(err) }
	println("Hello! I am " + bot_user.FirstName)

	message := make(map[string]interface{})
	message["chat_id"] = 3899337
	message["text"] = "Hello! I am " + bot_user.FirstName
	//_, err = bot1.SendMessage(message)
	//if err != nil { log.Fatal(err) }


	for {
		updates, err := bot1.GetNewUpdates(nil)
		if err != nil { log.Fatal(err) }
//		println(fmt.Sprintf("%+v", updates))
		for _, update := range updates {
			replyText := ""
			if (strings.HasPrefix(update.Message.Text, "/start") || strings.HasPrefix(update.Message.Text, "/help")) {
				tempTxt := "As of now, I only support /start and /help commands."
				replyText = fmt.Sprintf("Hello _%s_,\n\nI am still under development.\n\n%s\n\nI was developed using [Telegram BotAPI GoLang](https://github.com/go-snh/tg) Library by @NiSaR.\n\nI am only a demo bot. If you like to make your own Telegram bot in GoLang, you can install the [library](https://github.com/go-snh/tg) using command `go get github.com/go-snh/tg`. You can also find the documentation for the library [here](https://godoc.org/github.com/go-snh/tg).\n\nFor more details, contact my [Boss](t.me/NiSaR)\n\nCheers,\n%s", update.Message.From.FirstName, tempTxt, bot_user.FirstName)
			}
			if replyText != "" {
				replyParams := make(map[string]interface{})
				replyParams["chat_id"] = update.Message.Chat.Id
				replyParams["text"] = replyText
				replyParams["parse_mode"] = "Markdown"
				_, err = bot1.SendMessage(replyParams)
				if err != nil { log.Fatal(err) }
			}
		}
		time.Sleep(1 * time.Second)
	}

	body := make(map[string]interface{})
	body["limit"] = 100
	updates, err := bot1.GetUpdates(body)
	println(fmt.Sprintf(">>> %d", len(updates)))
	if err != nil { log.Fatal(err) }
	updates, err = bot1.GetUpdates(body)
	println(fmt.Sprintf(">>> %d", len(updates)))
	if err != nil { log.Fatal(err) }
}

