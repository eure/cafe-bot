package main

// GoogleHomeへ再接続するタスク.
type castReconnectTask struct {
	bot *SlackBot
}

func newCastReconnectTask(bot *SlackBot) castReconnectTask {
	return castReconnectTask{
		bot: bot,
	}
}

func (castReconnectTask) getName() string {
	return "cast_reconnect_task"
}

func (t castReconnectTask) run() error {
	return t.bot.clients.SetCastClient()
}
