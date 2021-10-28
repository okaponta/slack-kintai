package main

import (
	"fmt"

	"github.com/slack-go/slack"

	"slack-kintai/readconfig"
)

func main() {
	fmt.Println("出勤！")
	conf, err := readconfig.ReadConfig("./config.json")
	if err != nil {
		fmt.Println("error ")
	}
	fmt.Println(conf)
	tkn := conf.Token
	c := slack.New(tkn)

	// MsgOptionText() の第二引数に true を設定すると特殊文字をエスケープする
	_, _, err = c.PostMessage(conf.Channels[0], slack.MsgOptionText("Hello World", true))
	if err != nil {
		panic(err)
	}
}
