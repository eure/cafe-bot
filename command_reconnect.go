package main

// reconnect用コマンド.
// GoogleHomeに再接続する.
func newCommandReconnect(d CommandData) Command {
	c := Command{}

	c.Add(newCastReconnectTask(d.Bot))
	return c
}
