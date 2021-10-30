package main

import (
	"flag"
	"fmt"

	"github.com/slack-go/slack"

	"slack-kintai/readconfig"
)

func main() {
	flag.Parse()
	args := flag.Args()
	valid := checkArgs(args)
	if !valid {
		return
	}

	conf, err := readconfig.ReadConfig("./config.json")
	if err != nil {
		fmt.Println("configuration error")
	}
	c := slack.New(conf.Token)

	if args[0] == "shukkin" {
		shukkin(c, conf)
	} else {
		taikin(c, conf)
	}
}

func checkArgs(args []string) bool {
	if len(args) != 1 {
		printArgErr()
		return false
	}
	if args[0] != "shukkin" && args[0] != "taikin" {
		printArgErr()
		return false
	}
	return true
}

func printArgErr() {
	fmt.Println("argument error!!")
	fmt.Println("Argument has to be 'shukkin' or 'taikin'")
}

func shukkin(c *slack.Client, conf readconfig.Config) {
	fmt.Println("post shukkin to slack")
	for _, channel := range conf.Channels {
		fmt.Println("channel:", channel)
		_, _, err := c.PostMessage(channel, slack.MsgOptionText(conf.Shukkin, true))
		if err != nil {
			panic(err)
		}
	}
}

func taikin(c *slack.Client, conf readconfig.Config) {
	fmt.Println("post taikin to slack")
	for _, channel := range conf.Channels {
		fmt.Println("channel:", channel)
		_, _, err := c.PostMessage(channel, slack.MsgOptionText(conf.Taikin, true))
		if err != nil {
			panic(err)
		}
	}
}
