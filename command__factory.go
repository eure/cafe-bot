package main

import (
	"log"
	"strings"
)

var emptyCommand = Command{}

// コマンドのファクトリ関数
func CreateCommand(d CommandData) Command {
	switch {
	case containsMention(d.Clients, d.Text):
		// メンションを含む場合
		return createCommandForMention(d)
	case containsOrder(d.Text):
		// メンション以外で注文情報を含む場合
		return newCommandFreewordOrder(d)
	case containsDelicious(d.Text):
		// メンション以外で美味しさ情報を含む場合
		return newCommandFreewordDelicious(d)
	default:
		// 空のコマンドを返却する
		return emptyCommand
	}
}

// コマンドのファクトリ関数（@メンション有りの場合）
func createCommandForMention(d CommandData) Command {
	words := strings.Split(d.Text, " ")
	if len(words) < 2 {
		return emptyCommand
	}

	item := words[1]
	switch item {
	case "ping":
		return newCommandPing(d)
	case "help":
		return newCommandHelp(d)
	case "menu":
		return newCommandMenu(d)
	case "mute":
		return newCommandSpeechOff(d)
	case "speech":
		return newCommandSpeechOn(d)
	case "on":
		return newCommandStatusOn(d)
	case "off":
		return newCommandStatusOff(d)
	case "debug-on":
		return newCommandDebugOn(d)
	case "debug-off":
		return newCommandDebugOn(d)
	case "status":
		return newCommandShowStatus(d)
	case "debug-say":
		return newCommandDebugSay(d)
	case "reconnect":
		return newCommandReconnect(d)
	default:
		return newCommandOrder(d)
	}
}

// コマンド用構造体.
// 複数のタスクを保持し、順次実行する.
type Command struct {
	logger *log.Logger
	tasks  []task
}

// タスクを追加する.
func (c *Command) Add(t task) {
	c.tasks = append(c.tasks, t)
}

// タスクを全て実行する.
func (c Command) Exec() {
	for _, task := range c.tasks {
		err := task.run()
		if err != nil {
			c.logger.Printf("[ERROR] task=[%s], error=%s", task.getName(), err.Error())
			return
		}
	}
}
