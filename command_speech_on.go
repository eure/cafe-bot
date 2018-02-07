package main

// speech用コマンド.
// 音声フラグをオンにする.
func newCommandSpeechOn(d CommandData) Command {
	c := Command{}
	c.Add(newBotStatusTask(d.Bot, "speech", true))

	task := newSlackSendChannleTask(d.Clients.slackRTM, d.SlackChannel, "音声をオンにしました")
	c.Add(task)
	return c
}
