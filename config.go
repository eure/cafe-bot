package main

import (
	"log"
	"os"
	"strconv"
)

// 設定値用の構造体
type Config struct {
	Token  string
	Logger *log.Logger
	Debug  bool
	Speech bool
}

// Slackトークンを取得する.
func (c Config) GetToken() string {
	if c.Token != "" {
		return c.Token
	}
	return getToken()
}

// ロガーを取得する.
func (c Config) GetLogger() *log.Logger {
	if c.Logger != nil {
		return c.Logger
	}

	return log.New(os.Stdout, "[SlackBot Log]: ", log.Lshortfile|log.LstdFlags)
}

// デバッグフラグがオンかどうか判定する.
func (c Config) IsDebug() bool {
	if c.Debug {
		return true
	}

	b, _ := strconv.ParseBool(os.Getenv("SLACK_BOT_DEBUG"))
	return b
}

// 音声フラグがオンかどうか判定する.
func (c Config) IsSpeech() bool {
	if c.Speech {
		return true
	}

	b, _ := strconv.ParseBool(os.Getenv("SLACK_BOT_SPEECH"))
	return b
}

// Slackトークン用の環境変数リスト
var envTokens = []string{
	"SLACK_RTM_TOKEN",
	"SLACK_BOT_TOKEN",
	"SLACK_TOKEN",
}

// Slackトークンを取得する.
func getToken() string {
	for _, envName := range envTokens {
		token := os.Getenv(envName)
		if token != "" {
			return token
		}
	}
	return ""
}
