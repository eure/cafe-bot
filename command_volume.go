package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// volume用コマンド.
// GoogleHomeの音量を設定する.
func newCommandVolume(d CommandData) Command {
	c := Command{}

	volume, err := getVolume(d.Text)
	if err != nil {
		c.Add(newSlackSendChannleTask(d.Clients.slackRTM, d.SlackChannel, err.Error()))
		return c
	}

	c.Add(newCastVolumeTask(d.Clients.castClient, volume))
	return c
}

func getVolume(text string) (float64, error) {
	words := strings.Split(text, " ")
	if len(words) < 2 {
		return 0, errors.New("音量を 0.0 - 1.0 の範囲で指定してください")
	}

	volText := words[2]
	volume, err := strconv.ParseFloat(volText, 64)
	switch {
	case err != nil,
		volume < 0.0,
		volume > 1.0:
		return 0, fmt.Errorf("音量は 0.0 - 1.0 の範囲で指定してください [%s]", volText)
	}
	return volume, nil
}
