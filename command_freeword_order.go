package main

import (
	"fmt"
	"regexp"
)

var reOrder = regexp.MustCompile("(.*[^がを])(が|を)?(欲しい|ほしい|つくって|作って|頼む|たのむ|ください|下さい|のみたい|飲みたい|お願い)")

// 自由に注文するコマンド.
// GoogleHomeで注文を発話する.
func newCommandFreewordOrder(d CommandData) Command {
	c := Command{}
	if !d.Bot.status {
		return emptyCommand
	}

	words := reOrder.FindStringSubmatch(d.Text)
	if len(words) != 4 {
		return emptyCommand
	}

	item := words[1]
	c.Add(newBotAddHistoryTask(d.Bot, d.User, item))

	text := fmt.Sprintf("%sさんが「%s」を欲しいって言っています。", d.User, item)
	c.Add(newCastPlayTask(d.Clients.castClient, text))
	c.Add(newSlackReplyThreadTask(d.Clients.slackRTM, d.SlackChannel, text, d.ThreadTimestamp))
	return c
}

// 注文情報が含まれているかどうか判定する.
func containsOrder(text string) bool {
	return reOrder.MatchString(text)
}
