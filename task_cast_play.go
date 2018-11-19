package main

import (
	"errors"

	"github.com/evalphobia/google-home-client-go/googlehome"
)

// Google Homeで発話するタスク.
type castPlayTask struct {
	client  *googlehome.Client
	text    string
	enabled bool
}

func newCastPlayTask(client *googlehome.Client, text string, enabled bool) castPlayTask {
	return castPlayTask{
		client:  client,
		text:    text,
		enabled: enabled,
	}
}

func (castPlayTask) getName() string {
	return "cast_play_task"
}

func (t castPlayTask) run() error {
	if t.client == nil {
		return errors.New("Google Homeクライアントが設定されていません")
	}

	return t.client.Notify(t.text)
}
