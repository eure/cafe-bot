package main

// ping用コマンド.
// Slack上でPONGと発言する.
func newCommandPing(d CommandData) Command {
	c := Command{}

	task := newSlackSendChannleTask(d.Clients.slackRTM, d.SlackChannel, "PONG")
	c.Add(task)
	return c
}
