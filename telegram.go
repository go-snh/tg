package tg

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"time"
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
	token         string
	baseURL       string
	updatesOffset int
}

type Response struct {
	Status bool   `json:"ok"`
	Result string `json:"result"`
}

type status struct {
	Status bool `json:"ok"`
}

type apiError struct {
	Status      bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

/** Utility Functions **/

func funcName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
func debug_println(message string) {
	if debug {
		println("DEBUG: " + message)
	}
}
func generateMethodUrl(methodStr string, params map[string]interface{}) string {

	methodStr = methodStr + "?"
	for key, value := range params {
		data := ""
		switch value.(type) {
		case bool:
			data = fmt.Sprintf("%t", value.(bool))
		case int:
			data = fmt.Sprintf("%d", value.(int))
		case float32:
			data = fmt.Sprintf("%g", value.(float32))
		case string:
			data = url.QueryEscape(value.(string))
			//data = value.(string)
		}
		methodStr = fmt.Sprintf("%s%s=%s&", methodStr, key, data)
		debug_println("Key  = " + key)
		debug_println(fmt.Sprint("Value = " + data + "\n"))
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
	debug_println(fullUrl)
	resp, _ := myClient.Get(fullUrl)
	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		err_obj := new(apiError)
		json.Unmarshal(response, err_obj)
		debug_println(fmt.Sprintf("%+v", err_obj))
		return nil, errors.New(fmt.Sprintf("error_code: %d Description: %s", err_obj.ErrorCode, err_obj.Description))
	}
	/*
	if !strings.Contains(string(response), "{\"ok\":true") {
		err_obj := new(apiError);
		json.Unmarshal(response, err_obj)
		debug_println(fmt.Sprintf("%+v", err_obj))
		return nil, errors.New(fmt.Sprint("%s returned failure.\nFull URL: %s", methodName, fullUrl))
	}
	*/
	return response, nil
}

// https://core.telegram.org/bots/api#getme
// Return Type: User
func (bot *TelegramBot) GetMe() (*User, error) {

	response := new(struct {
		status, User *User `json:"result"`
	})
	str, err := bot.callAPI("getMe", nil)
	if err != nil { return nil, err }
	err = json.Unmarshal(str, &response)
	if err != nil { return nil, err }
	return response.User, nil
}

// https://core.telegram.org/bots/api#sendmessage
// Return Type: Message
func (bot *TelegramBot) SendMessage(params map[string]interface{}) (*Message, error) {
	response := new(struct {
		status, Message *Message `json:"result"`
	})
	if (params == nil) {
		return nil, errors.New(fmt.Sprint("params not passed in SendMessage"))
	}
	//params["text"] = url.QueryEscape(params["text"])
	str, err := bot.callAPI("sendMessage", params)
	if err != nil { return nil, err }
	err = json.Unmarshal(str, &response)
	if err != nil { return nil, err }
	return response.Message, nil
}

// https://core.telegram.org/bots/api#getupdates
// Return Type: []Updates
func (bot* TelegramBot) GetUpdates(params map[string]interface{}) ([]Update, error) {
	response := new(struct {
		status, Updates []Update `json:"result"`
	})
	str, err := bot.callAPI("getUpdates", params)
	params = nil
	if err != nil { return nil, err }
	err = json.Unmarshal(str, &response)
	if err != nil { return nil, err }
	if len(response.Updates) > 0 {
		bot.updatesOffset = response.Updates[len(response.Updates)-1].UpdateId + 1
	}
	return response.Updates, nil
}

// Return Type: []Updates
func (bot* TelegramBot) GetNewUpdates(params map[string]interface{}) ([]Update, error) {
	if params != nil {
		_, present := params["offset"]
		if !present || params["offset"] == 0 {
			params["offset"] = bot.updatesOffset
		}
	} else {
		params = make(map[string]interface{})
		params["offset"] = bot.updatesOffset
	}
	return bot.GetUpdates(params)
}
