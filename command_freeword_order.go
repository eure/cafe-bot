package main

import (
	"fmt"
	"regexp"
	"strings"
)

var orderWords = []string{
	"欲しい",
	"ほしい",
	"作って",
	"つくって",
	"頼む",
	"たのむ",
	"下さい",
	"ください",
	"飲みたい",
	"のみたい",
	"お願い",
	"おねがい",
	"please",
	"ぷりーず",
}

var reOrder = regexp.MustCompile(fmt.Sprintf("(.*[^がを])(が|を)?(%s)", strings.Join(orderWords, "|")))

// 自由に注文するコマンド.
// GoogleHomeで注文を発話する.
func newCommandFreewordOrder(d CommandData) Command {
	c := Command{}
	if !d.Bot.status {
		return emptyCommand
	}

	words := reOrder.FindStringSubmatch(strings.ToLower(d.Text))
	if len(words) != 4 {
		return emptyCommand
	}

	item := words[1]
	c.Add(newBotAddHistoryTask(d.Bot, d.User, item))

	text := fmt.Sprintf("%sさんが「%s」を欲しいって言っています。", d.User, item)
	c.Add(newCastPlayTask(d.Clients.castClient, text, d.Bot.isSpeech))
	c.Add(newSlackReplyThreadTask(d.Clients.slackRTM, d.SlackChannel, text, d.ThreadTimestamp))
	return c
}

// 注文情報が含まれているかどうか判定する.
func containsOrder(text string) bool {
	return reOrder.MatchString(text)
}
