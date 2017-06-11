package main

import (
	"github.com/go-snh/tg"
	"fmt"
	"log"
	"io/ioutil"
)
func main() {
	println(fmt.Sprint("Starting Program..."))
	bot1 := new(tg.TelegramBot)

	token, _ := ioutil.ReadFile("go-test/MyBot.token")
	if !bot1.SetToken(token) {
		log.Fatal("Unable to use bot token.\nMake sure you have saved correct token in MyBot.token file")
	}
	bot_user, err := bot1.GetMe()
	if err != nil {
		log.Fatal(err)
	}
	println("Hello! I am "+bot_user.FirstName)

	message := make(map[string]interface{})
	message["chat_id"] = 3899337
	message["text"] = "Hello! I am "+bot_user.FirstName
	_, err = bot1.SendMessage(message)
	if err != nil {
		log.Fatal(err)
	}
}

