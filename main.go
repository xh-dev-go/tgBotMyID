package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/xh-dev-go/tgBotMyID/entities/update"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

const VERSION = "1.0.0"

func maxUpdateId(update2 []update.Update) int64 {
	var max int64 = 0
	for _, item := range update2 {
		if max < item.UpdateId {
			max = item.UpdateId
		}
	}
	return max
}

func getUpdateUrlWithOffset(token string, offset int64) string {
	return fmt.Sprintf(`https://api.telegram.org/bot%s/getUpdates?offset=%d`, token, offset)
}
func getUpdateUrl(token string) string {
	return fmt.Sprintf(`https://api.telegram.org/bot%s/getUpdates`, token)
}
func getMaxId(token string) int64 {
	var result update.Response
	if response, err := http.Get(getUpdateUrl(token)); err != nil {
		panic(err)
	} else if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		panic(err)
	} else if !result.Ok {
		panic("error")
	} else {
		return maxUpdateId(result.Result)
	}
}

func sendMessage(token string, to int64, msg string) {
	data := url.Values{
		"parse_mode": {"MarkdownV2"},
		"chat_id":    {strconv.FormatInt(to, 10)},
		"text":       {msg},
	}
	if resp, err := http.PostForm(fmt.Sprintf(`https://api.telegram.org/bot%s/sendMessage`, token), data); err != nil {
		panic(err)
	} else if byte, err := ioutil.ReadAll(resp.Body); err != nil {
		panic(err)
	} else {
		println(string(byte))
	}

}

func main() {
	const cmd_token = "token"
	const cmd_version = "version"
	var token string
	var checkVersion bool
	flag.BoolVar(&checkVersion, cmd_version, false, "Show version of application")
	flag.StringVar(&token, cmd_token, "", "token of the bot")
	flag.Parse()

	if checkVersion {
		print(VERSION)
		os.Exit(0)
	}

	if token == "" {
		print("No token provided")
		flag.Usage()
		os.Exit(1)
	}

	curMessageId := getMaxId(token)

	for {
		var result update.Response
		if response, err := http.Get(getUpdateUrlWithOffset(token, curMessageId+1)); err != nil {
			panic(err)
		} else if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
			panic(err)
		} else if !result.Ok {
			panic("error")
		} else {
			for _, item := range result.Result {
				fromId := item.Message.From.Id
				isBot := item.Message.From.IsBot
				firstName := item.Message.From.FirstName
				lastName := item.Message.From.LastName

				msg := fmt.Sprintf("*Name*:%s, %s\n*ID*: %d\n*Is bot*: %t\n", firstName, lastName, fromId, isBot)
				sendMessage(token, fromId, msg)
				curMessageId = item.UpdateId
			}

			time.Sleep(time.Second)
		}

	}

}
