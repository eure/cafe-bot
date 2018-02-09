package main

// SlackBotの注文履歴を追加するタスク.
type botAddHistoryTask struct {
	bot  *SlackBot
	user string
	item string
}

func newBotAddHistoryTask(bot *SlackBot, user, item string) botAddHistoryTask {
	return botAddHistoryTask{
		bot:  bot,
		user: user,
		item: item,
	}
}

func (botAddHistoryTask) getName() string {
	return "bot_add_history_task"
}

func (t botAddHistoryTask) run() error {
	t.bot.addOrderHistory(newOrder(t.user, t.item))
	return nil
}
