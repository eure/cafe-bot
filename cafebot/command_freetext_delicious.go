package cafebot

import (
	"fmt"
	"regexp"

	"github.com/eure/bobo-googlehome/googlehome"
	"github.com/eure/bobo/command"
)

var reDelicious = regexp.MustCompile("(.*[^がを])(が|を)?(美味しい|おいしい|うまい|美味い|ウマイ|旨い)")

// FreeTextDeliciousCommand is a command to express delicios feelings.
// 正規表現でフリーテキストから美味しさを表現するコマンド.
var FreeTextDeliciousCommand = command.BasicCommandTemplate{
	Help:   "Express feeling from free text",
	Regexp: reDelicious,
	GenerateFn: func(d command.CommandData) command.Command {
		c := command.Command{}

		words := reDelicious.FindStringSubmatch(d.RawText)
		if len(words) != 4 {
			return c
		}

		item := words[1]

		text := fmt.Sprintf("%sさんが「%s」を美味しいって言っています。", d.SenderName, item)
		c.Add(command.NewReplyThreadEngineTask(d.Engine, d.Channel, text, d.ThreadTimestamp))

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
