package cafebot

import (
	"fmt"
	"strings"

	"github.com/eure/bobo-googlehome/googlehome"
	"github.com/eure/bobo/command"
)

// DebugSayCommand is a command to debug playing voice on google home.
// GoogleHomeで発話を行うコマンド.
var DebugSayCommand = command.BasicCommandTemplate{
	NoHelp:         true,
	MentionCommand: "debug-say",
	GenerateFn: func(d command.CommandData) command.Command {
		c := command.Command{}

		text := getDebugWord(d.Text)

		task, err := googlehome.NewCastPlayTask(text)
		if err != nil {
			errMessage := fmt.Sprintf("[ERROR]\t[NewCastPlayTask]\t`%s`", err.Error())
			task := command.NewReplyEngineTask(d.Engine, d.Channel, errMessage)
			c.Add(task)
			return c
		}

		c.Add(task)
		return c
	},
}

func getDebugWord(text string) string {
	words := strings.Split(text, " ")
	if len(words) < 2 {
		return "ちぇけらっちょ"
	}

	return strings.Join(words[2:], "")
}
