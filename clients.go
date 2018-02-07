package main

import (
	"fmt"
	"log"

	"github.com/evalphobia/google-home-client-go/googlehome"
	"github.com/nlopes/slack"
)

// 外部クライアント管理用の構造体.
type ClientManager struct {
	logger *log.Logger

	//Slackクライアント
	slackClient  *slack.Client
	slackRTM     *slack.RTM
	slackBotID   string
	slackBotName string

	// Google Homeクライアント
	castClient *googlehome.Client
}

func newClientManagerWithSlack(conf Config) (ClientManager, error) {
	cli := slack.New(conf.GetToken())

	logger := conf.GetLogger()
	if logger != nil {
		slack.SetLogger(conf.GetLogger())
	}

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

// ログを記録する.
func (c ClientManager) Logging(typ, msg string, v ...interface{}) {
	if c.logger == nil {
		return
	}
	c.logger.Printf("[%s] %s", typ, fmt.Sprintf(msg, v...))
}
