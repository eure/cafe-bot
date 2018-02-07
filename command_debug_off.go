package main

// debug-off用コマンド.
// デバッグフラグをオフにする.
func newCommandDebugOff(d CommandData) Command {
	c := Command{}
	c.Add(newBotStatusTask(d.Bot, "debug", false))
	return c
}
