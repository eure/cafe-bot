package main

import (
	"github.com/nlopes/slack"
)

// Slackのチャンネルに発言するタスク.
type slackSendChannelTask struct {
	rtm     *slack.RTM
	channel string
	text    string
}

func newSlackSendChannleTask(rtm *slack.RTM, channel, text string) slackSendChannelTask {
	return slackSendChannelTask{
		rtm:     rtm,
		channel: channel,
		text:    text,
	}
}

func (slackSendChannelTask) getName() string {
	return "slack_send_channel_task"
}

func (t slackSendChannelTask) run() error {
	reply := t.rtm.NewOutgoingMessage(t.text, t.channel)
	t.rtm.SendMessage(reply)
	return nil
}
