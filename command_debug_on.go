package main

// debug-on用コマンド.
// デバッグフラグをオンにする.
func newCommandDebugOn(d CommandData) Command {
	c := Command{}
	c.Add(newBotStatusTask(d.Bot, "debug", true))
	return c
}
