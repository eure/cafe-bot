package main

// reload用コマンド.
// Botをリロードする.
func newCommandReload(d CommandData) Command {
	c := Command{}

	c.Add(newSlackSendChannleTask(d.Clients.slackRTM, d.SlackChannel, "リロードします"))
	c.Add(newBotReloadTask(d.Bot))
	return c
}
