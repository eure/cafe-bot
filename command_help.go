package main

import "fmt"

// help用コマンド.
// Slack上で使い方を発言する.
func newCommandHelp(d CommandData) Command {
	c := Command{}

	text := fmt.Sprintf("```%s```", getUsage())
	task := newSlackSendChannleTask(d.Clients.slackRTM, d.SlackChannel, text)
	c.Add(task)
	return c
}

// ヘルプ用テキストを返却する.
func getUsage() string {
	return `
/* コマンド */
help    // ヘルプ表示
menu    // メニュー表示
history // 注文履歴表示

/* オーダー */
<ドリンク名> <hot/ice>  // hot or iceはドリンクによっては省略可


/* （その他） */
mute      // 音声をオフに変更
speech    // 音声をオンに変更
on        // Slackでの反応をオンに変更
off       // Slackでの反応をオフに変更
reload    // ボットをリロードする
status    // ステータスを表示する
`
}
