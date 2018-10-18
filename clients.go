package main

import (
	"fmt"
	"log"
	"time"

	"github.com/evalphobia/google-home-client-go/googlehome"
	"github.com/nlopes/slack"
)

// 外部クライアント管理用の構造体.
type ClientManager struct {
	logger *log.Logger

	//Slackクライアント
	slackClose   chan struct{}
	slackClient  *slack.Client
	slackRTM     *slack.RTM
	slackBotID   string
	slackBotName string

	// Google Homeクライアント
	castClient *googlehome.Client
}

func newClientManagerWithSlack(conf Config) (ClientManager, error) {
	var opts []slack.Option
	logger := conf.GetLogger()
	if logger != nil {
		opts = append(opts, slack.OptionLog(conf.GetLogger()))
	}

	cli := slack.New(conf.GetToken(), opts...)

	// Tokenを用いたSlackとの疎通テスト
	resp, err := cli.AuthTest()
	if err != nil {
		return ClientManager{}, err
	}
	return ClientManager{
		slackClient:  cli,
		slackRTM:     cli.NewRTM(),
		slackBotID:   resp.UserID,
		slackBotName: resp.User,
		logger:       logger,
		slackClose:   make(chan struct{}, 1),
	}, nil
}

// Google Homeクライアントをセットする.
func (c *ClientManager) SetCastClient(config ...googlehome.Config) error {
	var conf googlehome.Config
	switch {
	case len(config) != 0:
		conf = config[0]
	default:
		// 空の場合はデフォルトコンフィグを使う
		conf = googlehome.Config{}
	}
	castClient, err := googlehome.NewClientWithConfig(conf)
	if err != nil {
		return err
	}
	castClient.SetLang("ja")
	c.castClient = castClient
	return nil
}

// 長期間起動時に接続が切れてしまうため
func (c *ClientManager) KeepAlive() {
	go c.keepAliveSlack()
}

func (c *ClientManager) keepAliveSlack() {
	cli := c.slackRTM
	ticker := time.NewTicker(time.Second * 300)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// send empty message.
			cli.SendMessage(cli.NewTypingMessage(""))
		case <-c.slackClose:
			cli.Disconnect()
			return
		}
	}
}

func (c *ClientManager) CloseAll() {
	c.slackClose <- struct{}{}
}

// ログを記録する.
func (c ClientManager) Logging(typ, msg string, v ...interface{}) {
	if c.logger == nil {
		return
	}
	c.logger.Printf("[%s] %s", typ, fmt.Sprintf(msg, v...))
}
