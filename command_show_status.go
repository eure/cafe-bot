package main

import (
	"fmt"
	"strings"
)

// status用コマンド.
// Slack上でステータス値を表示する.
func newCommandShowStatus(d CommandData) Command {
	c := Command{}

	text := fmt.Sprintf("```%s```", showStatus(d.Bot))
	task := newSlackSendChannleTask(d.Clients.slackRTM, d.SlackChannel, text)
	c.Add(task)
	return c
}

// 現在のステータスを表示する.
func showStatus(bot *SlackBot) string {
	var result []string
	result = append(result, fmt.Sprintf("- 注文ステータス:\t%v", bot.status))
	result = append(result, fmt.Sprintf("- 音声ステータス:\t%v", bot.isSpeech))
	result = append(result, fmt.Sprintf("- デバッグフラグ:\t%v", bot.isDebug))
	if bot.clients.castClient != nil {
		result = append(result, "======================")
		result = append(result, fmt.Sprintf("- Google Home IP:\t%v", bot.clients.castClient.GetIPv4()))
	}
	return strings.Join(result, "\n")
}
