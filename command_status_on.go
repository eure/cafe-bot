package main

// on用コマンド.
// 反応ステータスをオンにする.
func newCommandStatusOn(d CommandData) Command {
	c := Command{}
	c.Add(newBotStatusTask(d.Bot, "status", true))

	task := newSlackSendChannleTask(d.Clients.slackRTM, d.SlackChannel, "ステータスをオンにしました")
	c.Add(task)
	return c
}
