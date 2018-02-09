package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 注文用履歴コマンド.
// 注文履歴を表示する.
func newCommandHistory(d CommandData) Command {
	c := Command{}

	page := 0
	words := strings.Split(d.Text, " ")
	if len(words) == 3 {
		var err error
		page, err = strconv.Atoi(words[2])
		if err != nil {
			text := fmt.Sprintf("ページ番号は整数で指定して下さい: page=%s", words[2])
			c.Add(newSlackReplyThreadTask(d.Clients.slackRTM, d.SlackChannel, text, d.ThreadTimestamp))
			return c
		}
	}

	orders, start, last := d.Bot.getOrderHistory(page)
	result := make([]string, len(orders))
	for i, hist := range orders {
		result[i] = hist.String()
	}

	size := d.Bot.getOrderHistorySize()
	text := fmt.Sprintf("```logs:%d page:%d (%d-%d)\n\n%s```", size, page, start, last, strings.Join(result, "\n"))
	c.Add(newSlackSendChannleTask(d.Clients.slackRTM, d.SlackChannel, text))
	return c
}
