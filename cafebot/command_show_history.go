package cafebot

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/eure/bobo/command"
)

// ShowOrderHistoryCommand is a command to show order history.
// 注文履歴を確認するコマンド.
var ShowOrderHistoryCommand = command.BasicCommandTemplate{
	Help:           "Show order history",
	MentionCommand: "history",
	GenerateFn: func(d command.CommandData) command.Command {
		c := command.Command{}

		page := 0
		words := strings.Split(d.Text, " ")
		if len(words) == 3 {
			var err error
			page, err = strconv.Atoi(words[2])
			if err != nil {
				errMessage := fmt.Sprintf("ページ番号は整数で指定して下さい: page=%s", words[2])
				task := command.NewReplyThreadEngineTask(d.Engine, d.Channel, errMessage, d.ThreadTimestamp)
				c.Add(task)
				return c
			}
		}

		task := NewShowHistoryTask(page, func(history string) error {
			return command.NewReplyEngineTask(d.Engine, d.Channel, fmt.Sprintf("```%s```", history)).Run()
		})
		c.Add(task)
		return c
	},
}
