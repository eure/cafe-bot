package cafebot

import (
	"fmt"

	"github.com/eure/bobo/command"
)

// ShowMenuCommand is a command to show menu.
// メニューを表示するコマンド.
var ShowMenuCommand = command.BasicCommandTemplate{
	Help:           "Show menu",
	MentionCommand: "menu",
	GenerateFn: func(d command.CommandData) command.Command {
		c := command.Command{}

		text := fmt.Sprintf("```%s\n--------\n* グランデサイズあり```", getAllMenu())

		task := command.NewReplyEngineTask(d.Engine, d.Channel, text)
		c.Add(task)
		return c
	},
}
