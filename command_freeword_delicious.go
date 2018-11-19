package main

import (
	"fmt"
	"regexp"
)

var reDelicious = regexp.MustCompile("(.*[^がを])(が|を)?(美味しい|おいしい|うまい|美味い|ウマイ|旨い)")

// 美味しさを表現するコマンド.
// GoogleHomeで美味しさを発話する.
func newCommandFreewordDelicious(d CommandData) Command {
	c := Command{}
	if !d.Bot.status {
		return emptyCommand
	}

	words := reDelicious.FindStringSubmatch(d.Text)
	if len(words) != 4 {
		return emptyCommand
	}

	item := words[1]
	text := fmt.Sprintf("%sさんが「%s」を美味しいって言っています。", d.User, item)
	c.Add(newCastPlayTask(d.Clients.castClient, text, d.Bot.isSpeech))
	c.Add(newSlackReplyThreadTask(d.Clients.slackRTM, d.SlackChannel, text, d.ThreadTimestamp))
	return c
}

// 美味しさ情報が含まれているかどうか判定する.
func containsDelicious(text string) bool {
	return reDelicious.MatchString(text)
}
