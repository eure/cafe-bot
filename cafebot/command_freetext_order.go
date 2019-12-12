package cafebot

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/eure/bobo-googlehome/googlehome"
	"github.com/eure/bobo/command"
)

var reOrder = regexp.MustCompile(fmt.Sprintf("(.*[^がを])(が|を)?(%s)", strings.Join(orderWords, "|")))

var orderWords = []string{
	"欲しい",
	"ほしい",
	"作って",
	"つくって",
	"頼む",
	"たのむ",
	"下さい",
	"ください",
	"飲みたい",
	"のみたい",
	"お願い",
	"おねがい",
	"ぷりーず",
	"please",
}

// FreeTextOrderCommand is a command to order drink by regex.
// 正規表現でフリーテキストから注文を行うコマンド.
var FreeTextOrderCommand = command.BasicCommandTemplate{
	Help:   "Order from free text",
	Regexp: reOrder,
	GenerateFn: func(d command.CommandData) command.Command {
		c := command.Command{}

		words := reOrder.FindStringSubmatch(strings.ToLower(d.RawText))
		if len(words) != 4 {
			return c
		}

		item := words[1]
		c.Add(NewAddHistoryTask(d.SenderName, item))

		text := fmt.Sprintf("%sさんが「%s」を欲しいって言っています。", d.SenderName, item)
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
