package main

import "fmt"

// menu用コマンド.
// Slack上でメニュー一覧を発言する.
func newCommandMenu(d CommandData) Command {
	c := Command{}

	text := fmt.Sprintf("```%s```", getAllMenu())
	task := newSlackSendChannleTask(d.Clients.slackRTM, d.SlackChannel, text)
	c.Add(task)
	return c
}
