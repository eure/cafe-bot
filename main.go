package main

import (
	"os"
)

// エントリーポイント
func main() {
	bot, err := NewSlackBot()
	if err != nil {
		panic(err)
	}
	os.Exit(bot.runRTM())
}
