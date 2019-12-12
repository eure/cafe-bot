package cafebot

import (
	"fmt"

	"github.com/eure/bobo-googlehome/googlehome"
	"github.com/eure/bobo/command"
)

// ReloadCommand is a command to reload bot.
// ボットを初期化するコマンド.
var ReloadCommand = command.BasicCommandTemplate{
	Help:           "Reload bot",
	MentionCommand: "reload",
	GenerateFn: func(d command.CommandData) command.Command {
		c := command.Command{}
		c.Add(command.NewReplyEngineTask(d.Engine, d.Channel, "Reloading..."))
		c.Add(command.NewReloadEngineTask(d.Engine))

		task, err := googlehome.NewReloadClientTask()
		if err != nil {
			errMessage := fmt.Sprintf("[ERROR]\t[NewReloadClientTask]\t`%s`", err.Error())
			c.Add(command.NewReplyEngineTask(d.Engine, d.Channel, errMessage))
		}
		c.Add(task)
		return c
	},
}
