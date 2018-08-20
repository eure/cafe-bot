package main

import (
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/nlopes/slack"
)

// エラーコード
const (
	errorCodeNone        = 0
	errorCodeExit        = 2
	errorCodeGeneral     = 3
	errorCodeInvalidAuth = 4
)

// Slack Bot用の構造体
type SlackBot struct {
	//各種クライアント
	clients ClientManager

	// ボット自身のSlack情報
	botID   string
	botName string

	users map[string]string // ユーザー名のキャッシュ

	// 各種フラグ
	isDebug  bool
	isSpeech bool
	status   bool

	maxOrderHistory int
	orderHistory    []Order

	closeChan  chan int
	reloadChan chan struct{}
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

	bot := &SlackBot{
		clients:         cli,
		botID:           cli.slackBotID,
		botName:         cli.slackBotID,
		users:           make(map[string]string),
		isDebug:         conf.IsDebug(),
		isSpeech:        conf.IsSpeech(),
		status:          true,
		closeChan:       make(chan int, 1),
		reloadChan:      make(chan struct{}, 1),
		maxOrderHistory: conf.GetMaxHistory(),
	}
	go bot.catchSignal()

	return bot, nil
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

	for {
		select {
		case errCode := <-b.closeChan:
			return errCode
		case <-b.reloadChan:
			b.exitNoError()
			return errorCodeNone
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				// ユーザーの投稿に反応するイベント
				b.processMessage(ev)
			case *slack.DisconnectedEvent:
				b.Logging("DisconnectedEvent", "intentional=%t", ev.Intentional)
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
			}
		}
	}
	return errorCodeGeneral
}

// ユーザーからの投稿を内容に応じて処理する.
func (b *SlackBot) processMessage(ev *slack.MessageEvent) {
	switch {
	case ev.Text == "",
		ev.Hidden,
		b.botID == ev.Msg.User:
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

// .注文履歴を追加する.
func (b *SlackBot) getOrderHistory(page int) (orders []Order, start, last int) {
	const showSize = 10
	pageFirst := (page) * showSize
	pageLast := (page + 1) * showSize

	size := len(b.orderHistory)
	if size <= pageFirst {
		return nil, 0, 0
	}

	if size < pageLast {
		pageLast = size
	}
	return b.orderHistory[pageFirst:pageLast], pageFirst + 1, pageLast
}

func (b *SlackBot) getOrderHistorySize() int {
	return len(b.orderHistory)
}

// .注文履歴を追加する.
func (b *SlackBot) addOrderHistory(o Order) {
	max := b.maxOrderHistory
	if len(b.orderHistory) < max {
		b.orderHistory = append([]Order{o}, b.orderHistory...)
		return
	}

	b.orderHistory = append([]Order{o}, b.orderHistory[0:max-1]...)
	return
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

// OSシグナルを処理する.
func (b SlackBot) catchSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	for {
		s := <-ch
		switch s {
		case syscall.SIGHUP:
			b.Logging("catchSignal", "syscall.SIGHUP")
			b.reloadChan <- struct{}{}
			return
		case syscall.SIGINT:
			b.Logging("catchSignal", "syscall.SIGINT")
			b.exit()
			return
		case syscall.SIGTERM:
			b.Logging("catchSignal", "syscall.SIGTERM")
			b.exit()
			return
		case syscall.SIGQUIT:
			b.Logging("catchSignal", "syscall.SIGQUIT")
			b.exit()
			return
		}
	}
}

func (b SlackBot) exit() {
	b.clients.CloseAll()
	b.closeChan <- errorCodeExit
}

func (b SlackBot) exitNoError() {
	b.clients.CloseAll()
	b.closeChan <- errorCodeNone
}
