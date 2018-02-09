package main

import (
	"fmt"
	"os"
)

// エントリーポイント
func main() {
	fmt.Printf("pid: %d\n", os.Getpid())
	for {
		errCode := run()
		if errCode != errorCodeNone {
			os.Exit(errCode)
		}
	}
}

func run() int {
	bot, err := NewSlackBot()
	if err != nil {
		panic(err)
	}

	return bot.runRTM()
}
