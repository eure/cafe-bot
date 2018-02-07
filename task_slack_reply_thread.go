package main

import (
	"github.com/nlopes/slack"
)

// Slackのスレッドに返信するタスク.
type slackReplyThreadTask struct {
	rtm     *slack.RTM
	channel string
	text    string
	thread  string
}

func newSlackReplyThreadTask(rtm *slack.RTM, channel, text, thread string) slackReplyThreadTask {
	return slackReplyThreadTask{
		rtm:     rtm,
		channel: channel,
		text:    text,
		thread:  thread,
	}
}

func (slackReplyThreadTask) getName() string {
	return "slack_reply_thread_task"
}

func (t slackReplyThreadTask) run() error {
	reply := t.rtm.NewOutgoingMessage(t.text, t.channel)
	reply.ThreadTimestamp = t.thread

	t.rtm.SendMessage(reply)
	return nil
}
