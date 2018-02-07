package main

// off用コマンド.
// 反応ステータスをオフにする.
func newCommandStatusOff(d CommandData) Command {
	c := Command{}
	c.Add(newBotStatusTask(d.Bot, "status", false))

	task := newSlackSendChannleTask(d.Clients.slackRTM, d.SlackChannel, "ステータスをオフにしました")
	c.Add(task)
	return c
}
