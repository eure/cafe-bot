package main

// mute用コマンド.
// 音声フラグをオフにする.
func newCommandSpeechOff(d CommandData) Command {
	c := Command{}
	c.Add(newBotStatusTask(d.Bot, "speech", false))

	task := newSlackSendChannleTask(d.Clients.slackRTM, d.SlackChannel, "音声をオフにしました")
	c.Add(task)
	return c
}
