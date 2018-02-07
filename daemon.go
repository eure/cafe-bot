package main

import (
	"errors"

	"github.com/nlopes/slack"
)

// エラーコード
const (
	errorCodeGeneral     = 2
	errorCodeInvalidAuth = 3
)

// Slack Bot用の構造体
type SlackBot struct {
	//各種クライアント
	clients ClientManager
	// client *slack.Client
	// rtm    *slack.RTM
	// castClient *notifier.Client

	// ボット自身のSlack情報
	botID   string
	botName string

	users map[string]string // ユーザー名のキャッシュ

	// 各種フラグ
	isDebug  bool
	isSpeech bool
	status   bool
	orderLog []string
}

// 初期化済みSlackBotを返却する.
func NewSlackBot() (*SlackBot, error) {
	return NewSlackBotWithConfig(Config{})
}

// Configから、初期化済みSlackBotを返却する.
func NewSlackBotWithConfig(conf Config) (*SlackBot, error) {
	cli, err := newClientManagerWithSlack(conf)
	if err != nil {
		return nil, err
	}

	err = cli.SetCastClient()
	if err != nil {
		// Google Homeに接続できなくてもSlackBotはそのまま起動する
		// return nil, err
	}

	return &SlackBot{
		clients:  cli,
		botID:    cli.slackBotID,
		botName:  cli.slackBotID,
		users:    make(map[string]string),
		isDebug:  conf.IsDebug(),
		isSpeech: conf.IsSpeech(),
		status:   true,
	}, nil
}

// デバッグフラグ（ログ用）
func (b *SlackBot) SetDebug(f bool) {
	b.isDebug = f
}

// Google Homeで喋るかどうかのフラグ
func (b *SlackBot) SetSpeech(f bool) {
	b.isSpeech = f
}

// Slackの投稿に反応するかどうかのステータスフラグ
func (b *SlackBot) SetStatus(f bool) {
	b.status = f
}

// Slack Bot動作のメインループ
func (b *SlackBot) runRTM() int {
	rtm := b.clients.slackRTM
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			// ユーザーの投稿に反応するイベント
			b.processMessage(ev)
		case *slack.HelloEvent:
			// b.Logging("HelloEvent", "%v", ev)
		case *slack.ConnectedEvent:
			// b.Logging("ConnectedEvent", "info=%v, counter=%d", ev.Info, ev.ConnectionCount)
		case *slack.LatencyReport:
			// b.LogDebug("LatencyReport", "Current latency: %v", ev.Value)
		case *slack.RTMError:
			b.Logging("RTMError", "Error:%s", ev.Error())
		case *slack.ConnectionErrorEvent:
			if ev.Error() == "not_authed" && ev.Attempt > 2 {
				b.Logging("ConnectionErrorEvent", "Auth Error")
				return errorCodeInvalidAuth
			}
		case *slack.InvalidAuthEvent:
			b.Logging("InvalidAuthEvent", "Invalid credentials")
			return errorCodeInvalidAuth
		case *slack.FileCommentAddedEvent,
			*slack.FilePublicEvent,
			*slack.FileSharedEvent,
			*slack.FileChangeEvent,
			*slack.UserChangeEvent,
			*slack.UserTypingEvent,
			*slack.BotAddedEvent,
			*slack.ReactionAddedEvent,
			*slack.EmojiChangedEvent,
			*slack.ConnectingEvent,
			*slack.AckMessage:
			// 何もしない
		default:
			b.LogDebug("Unexpected", "Type=%T, data=%+v", ev, msg.Data)
		}
	}
	return errorCodeGeneral
}

// ユーザーからの投稿を内容に応じて処理する.
func (b *SlackBot) processMessage(ev *slack.MessageEvent) {
	if ev.Text == "" {
		return
	}

	userName := ev.User
	if !b.hasUser(userName) {
		// Slack APIを使い、SlackユーザーIDからユーザー名を取得する
		err := b.fetchUser(userName)
		if err != nil {
			return
		}
	}

	// ユーザー名キャッシュからユーザー名を取得する
	user := b.getUser(userName)
	text := trimSpaces(ev.Text)
	b.LogDebug("processMessage", "[DEBUG] text=%s", text)

	// テキスト内容に応じてコマンドを作成し実行する
	c := CreateCommand(CommandData{
		Bot:             b,
		Clients:         b.clients,
		User:            user,
		Text:            text,
		SlackChannel:    ev.Channel,
		ThreadTimestamp: ev.Timestamp,
	})
	c.Exec()
}

// ユーザーキャッシュからユーザー名を取得する
func (b *SlackBot) getUser(user string) string {
	return b.users[user]
}

// ユーザーキャッシュに該当ユーザーが存在するかどうか判定する.
func (b *SlackBot) hasUser(user string) bool {
	_, ok := b.users[user]
	return ok
}

// Slack APIを用いて、ユーザー名をキャッシュする.
func (b *SlackBot) fetchUser(user string) error {
	resp, err := b.clients.slackClient.GetUserInfo(user)
	switch {
	case err != nil:
		b.Logging("processMessage", "Error on `GetUserInfo`: %s", err.Error())
		return err
	case resp == nil:
		err := errors.New("response is nil")
		b.Logging("processMessage", "Error on `GetUserInfo`: %s", err.Error())
		return err
	}

	b.users[user] = resp.Name
	return nil
}

// ログを記録する.
func (b SlackBot) Logging(typ, msg string, v ...interface{}) {
	b.clients.Logging(typ, msg, v...)
}

// デバッグログを記録する.
func (b SlackBot) LogDebug(typ, msg string, v ...interface{}) {
	if b.isDebug {
		b.clients.Logging(typ, msg, v...)
	}
}
