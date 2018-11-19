package main

import "strings"

// debug-say用コマンド.
// GoogleHomeで発話する.
func newCommandDebugSay(d CommandData) Command {
	c := Command{}

	text := getDebugWord(d.Text)

	c.Add(newCastPlayTask(d.Clients.castClient, text, d.Bot.isSpeech))
	return c
}

func getDebugWord(text string) string {
	words := strings.Split(text, " ")
	if len(words) < 2 {
		return "ちぇけらっちょ"
	}

	return strings.Join(words[2:], "")
}
