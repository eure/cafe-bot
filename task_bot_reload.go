package main

// SlackBotをリロードするタスク.
type botReloadTask struct {
	bot *SlackBot
}

func newBotReloadTask(bot *SlackBot) botReloadTask {
	return botReloadTask{
		bot: bot,
	}
}

func (botReloadTask) getName() string {
	return "bot_reload_task"
}

func (t botReloadTask) run() error {
	t.bot.reloadChan <- struct{}{}
	return nil
}
