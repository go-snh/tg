package tg

import (
	"fmt"
	"net/http"
	"time"
	"errors"
	"encoding/json"
	"runtime"
	//	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	//"bytes"
	"strings"
)

/** Global Variables **/
var myClient = &http.Client{Timeout: 10 * time.Second}

var debug = false

/** Structures **/
type TelegramBot struct {
	token string
	baseURL string
}

type Response struct {
	Status     bool         `json:"ok"`
	Result     string       `json:"result"`
}

type status struct {
	Status     bool         `json:"ok"`
}

/** Utility Functions **/
func funcName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func generateMethodUrl(methodStr string, params map[string]interface{}) string {

	methodStr = methodStr + "?"
	for key, value := range params {
		methodStr = fmt.Sprintf("%s%s=%s&", methodStr, key, value)
	}
	return methodStr
}

/** Structure Functions **/
// Function to initialize the Bot and set Token
func (bot *TelegramBot) SetToken(token []byte) bool {
	bot.token = string(token)
	bot.token = strings.TrimSpace(string(token))
	if bot.token == "" {
		return false
	}
	bot.baseURL = fmt.Sprintf("https://api.telegram.org/bot%s", bot.token)
	return true
}

func (bot *TelegramBot) validate() bool {
	return (bot.token != "")
}

// Sends  GET/POST requests for the Bot API
func (bot *TelegramBot) callAPI(methodName string, params map[string]interface{}) ([]byte, error) {
	if !bot.validate() {
		return nil, errors.New(fmt.Sprint("Set the token using bot_name.SetToken() before using "))
	}
	methodStr := methodName
	if params != nil {
		methodStr = generateMethodUrl(methodName, params)
	}
	fullUrl := fmt.Sprintf("%s/%s", bot.baseURL, methodStr)
	if debug {
		println("DEBUG: "+ fullUrl)
	}
	resp, _ := myClient.Get(fullUrl)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(resp.Status))
	}
	response, _ := ioutil.ReadAll(resp.Body)
	if !strings.Contains(string(response), "{\"ok\":true") {
		return nil, errors.New(fmt.Sprint("%s returned failure.\nFull URL: %s", methodName, fullUrl))
	}
	return response, nil
}

// https://core.telegram.org/bots/api#getme
// Return Type: User
func (bot *TelegramBot) GetMe() (*User, error) {

	response := new(struct {status, User *User `json:"result"`})
	str, _ := bot.callAPI("getMe", nil)
	err := json.Unmarshal(str, &response)
	if err != nil {
		return nil, err
	}
	return response.User, nil;
}

// https://core.telegram.org/bots/api#sendmessage
// Return Type: Message
func (bot *TelegramBot) SendMessage(params map[string]interface{}) (*Message, error) {
	response := new (struct {status, Message *Message `json:"result"`})
	str, _ := bot.callAPI("getMe", params)
	err := json.Unmarshal(str, &response)
	if err != nil {
		return nil, err
	}
	return response.Message, nil
}
