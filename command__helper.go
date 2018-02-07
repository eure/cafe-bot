package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var reSpace = regexp.MustCompile(`(\s|　)+`)

// 連続したホワイトスペースを1つのホワイトスペースにする.
// また前後のホワイトスペースを除去する.
func trimSpaces(text string) string {
	return strings.TrimSpace(reSpace.ReplaceAllString(text, " "))
}

// ボットメンションが含まれるかどうか判定する.
func containsMention(clients ClientManager, text string) bool {
	words := strings.Split(text, " ")
	if len(words) == 0 {
		return false
	}

	switch words[0] {
	case clients.slackBotID,
		clients.slackBotName,
		fmt.Sprintf("@%s", clients.slackBotID),
		fmt.Sprintf("<@%s>", clients.slackBotID):
		return true
	}
	return false
}

// ひらがなからカタカナへの変換器
var kanaConv = unicode.SpecialCase{
	unicode.CaseRange{
		Lo: 0x3041, // ぁ
		Hi: 0x3093, // ん
		Delta: [unicode.MaxCase]rune{
			0x30a1 - 0x3041, // UpperCase でカタカナに変換
			0,               // LowerCase では変換しない
			0x30a1 - 0x3041, // TitleCase でカタカナに変換
		},
	},
	// カタカナをひらがなに変換
	unicode.CaseRange{
		Lo: 0x30a1, // ァ
		Hi: 0x30f3, // ン
		Delta: [unicode.MaxCase]rune{
			0,               // UpperCase では変換しない
			0x3041 - 0x30a1, // LowerCase でひらがなに変換
			0,               // TitleCase では変換しない
		},
	},
}

// テキストをカタカナへ変換する.
func toKatakana(str string) string {
	org := []rune(strings.ToLower(str))
	dst := make([]rune, len(org))
	for i, r := range org {
		if r <= 0x7f {
			// ASCIIはそのまま
			dst[i] = r
		} else {
			dst[i] = kanaConv.ToUpper(r)
		}
	}
	return string(dst)
}
