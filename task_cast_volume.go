package main

import (
	"errors"

	"github.com/evalphobia/google-home-client-go/googlehome"
)

// Google Homeで発話するタスク.
type castVolumeTask struct {
	client *googlehome.Client
	volume float64
}

func newCastVolumeTask(client *googlehome.Client, volume float64) castVolumeTask {
	return castVolumeTask{
		client: client,
		volume: volume,
	}
}

func (castVolumeTask) getName() string {
	return "cast_volume_task"
}

func (t castVolumeTask) run() error {
	if t.client == nil {
		return errors.New("Google Homeクライアントが設定されていません")
	}

	return t.client.SetVolume(t.volume)
}
