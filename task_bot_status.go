package main

import "fmt"

// SlackBotのフラグを変更するタスク.
type botStatusTask struct {
	bot        *SlackBot
	statusName string
	status     bool
}

func newBotStatusTask(bot *SlackBot, statusName string, status bool) botStatusTask {
	return botStatusTask{
		bot:        bot,
		statusName: statusName,
		status:     status,
	}
}

func (botStatusTask) getName() string {
	return "bot_status_task"
}

func (t botStatusTask) run() error {
	switch t.statusName {
	case "debug":
		t.bot.SetDebug(t.status)
	case "status":
		t.bot.SetStatus(t.status)
	case "speech":
		t.bot.SetSpeech(t.status)
	default:
		return fmt.Errorf("Unknown status: %s", t.statusName)
	}
	return nil
}
