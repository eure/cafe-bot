package main

import (
	"strings"
)

// バリスタ アカウント用発話コマンド.
// GoogleHomeで発話する.
func newCommandBaristaSay(d CommandData) Command {
	c := Command{}
	if !d.Bot.status {
		return emptyCommand
	}

	text := createTextForBaristaSay(d.Text)
	if text == "" {
		return emptyCommand
	}

	c.Add(newCastPlayTask(d.Clients.castClient, text, d.Bot.isSpeech))
	return c
}

// バリスタアカウントかどうか判定する.
func isBarista(user string) bool {
	return user == "the-local"
}

// バリスタアカウント向けの発話用テキストを作成する..
func createTextForBaristaSay(text string) string {
	switch {
	case strings.Contains(text, "到着") && strings.Contains(text, "ドア"):
		return "バリスタさんが到着いたしました。ドアをお願いします。"
	case strings.Contains(text, "ラストオーダー"):
		return "間もなくラストオーダーです"
	}

	return ""
}
