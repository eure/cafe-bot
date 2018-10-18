package main

import (
	"fmt"
	"os"
	"time"
)

// エントリーポイント
func main() {
	fmt.Printf("pid: %d\n", os.Getpid())
	for {
		errCode := run()
		if errCode != errorCodeNone {
			os.Exit(errCode)
		}
		time.Sleep(time.Second * 5)
		fmt.Printf("run again...\n")
	}
}

func run() (errCode int) {
	bot, err := NewSlackBot()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := recover(); err != nil {
			bot.Logging("PANIC", fmt.Sprintf("%v", err))
			bot.exit()
			errCode = errorCodeNone
		}
	}()

	return bot.runRTM()
}
