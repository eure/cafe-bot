package main

// コマンドの作成・実行時に使うデータセット
type CommandData struct {
	Bot             *SlackBot
	Clients         ClientManager
	Text            string
	User            string
	SlackChannel    string
	ThreadTimestamp string
}
