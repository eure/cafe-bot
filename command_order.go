package main

import (
	"fmt"
	"strings"
)

// 注文用コマンド.
// GoogleHomeで注文を発話する.
func newCommandOrder(d CommandData) Command {
	c := Command{}
	if !d.Bot.status {
		return emptyCommand
	}

	words := strings.Split(d.Text, " ")
	if len(words) < 2 {
		return emptyCommand
	}

	itemRaw := words[1]
	item, ok := getMenuName(itemRaw) // 名寄せ
	if !ok {
		return emptyCommand
	}

	// ホット or アイスの片方のみしか存在しないメニューの場合
	if !hasBothHeat(item) {
		c.Add(newBotAddHistoryTask(d.Bot, d.User, item))
		text := fmt.Sprintf("%sさんが「%s」が欲しいって言っています。", d.User, item)
		c.Add(newCastPlayTask(d.Clients.castClient, text))
		c.Add(newSlackReplyThreadTask(d.Clients.slackRTM, d.SlackChannel, text, d.ThreadTimestamp))
		return c
	}

	// 以下、ホットとアイスの両方が存在するメニューの場合
	errText := "hot か ice か選択して下さい"
	if len(words) < 3 {
		c.Add(newSlackSendChannleTask(d.Clients.slackRTM, d.SlackChannel, errText))
		return c
	}
	heat := getHeat(words[2])
	if !hasHeat(heat) {
		c.Add(newSlackSendChannleTask(d.Clients.slackRTM, d.SlackChannel, errText))
		return c
	}

	c.Add(newBotAddHistoryTask(d.Bot, d.User, fmt.Sprintf("%s (%s)", item, heat)))
	text := fmt.Sprintf("%sさんが「%sの%s」が欲しいって言っています。", d.User, heat, item)
	c.Add(newCastPlayTask(d.Clients.castClient, text))
	c.Add(newSlackReplyThreadTask(d.Clients.slackRTM, d.SlackChannel, text, d.ThreadTimestamp))
	return c
}
