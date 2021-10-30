package main

import (
	"fmt"

	"github.com/slack-go/slack"

	"slack-kintai/readconfig"
)

func main() {
	conf, err := readconfig.ReadConfig("./config.json")
	if err != nil {
		fmt.Println("error")
	}
	c := slack.New(conf.Token)

	for _, channel := range conf.Channels {
		fmt.Println("channel:", channel)
		_, _, err = c.PostMessage(channel, slack.MsgOptionText(conf.Comment, true))
		if err != nil {
			panic(err)
		}
	}
}
