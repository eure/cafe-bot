package cafebot

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

// メニュー一覧
const (
	menuCoffee    = "コーヒー"
	menuLatte     = "カフェラテ"
	menuMocha     = "カフェモカ"
	menuSoy       = "ソイラテ"
	menuChai      = "チャイラテ"
	menuHerb      = "ハーブティー"
	menuMelon     = "メロンソーダ"
	menuGinger    = "ジンジャーソーダ"
	menuChoco     = "ホットチョコレート"
	menuCaramel   = "キャラメルラテ"
	menuEspresso  = "エスプレッソ"
	menuAmericano = "アメリカーノ"
	menuMatcha    = "抹茶ラテ"
)

// 実際のメニューとユーザー入力メニューのマッピング
var menuMap = map[string]string{
	menuCoffee:    menuCoffee,
	"コーシー":        menuCoffee,
	"today":       menuCoffee,
	menuLatte:     menuLatte,
	"ラテ":          menuLatte,
	"latte":       menuLatte,
	"late":        menuLatte,
	"rate":        menuLatte,
	"ratte":       menuLatte,
	"cafelatte":   menuLatte,
	"cafelate":    menuLatte,
	menuMocha:     menuMocha,
	"モカ":          menuMocha,
	"mocha":       menuMocha,
	"moca":        menuMocha,
	"moka":        menuMocha,
	"mocka":       menuMocha,
	"cafemocha":   menuMocha,
	"cafemoca":    menuMocha,
	menuSoy:       menuSoy,
	"ソイ":          menuSoy,
	"soy":         menuSoy,
	"soi":         menuSoy,
	"soylatte":    menuSoy,
	"soylate":     menuSoy,
	menuChai:      menuChai,
	"チャイ":         menuChai,
	"chai":        menuChai,
	"tyai":        menuChai,
	menuHerb:      menuHerb,
	"ハーブ":         menuHerb,
	"herb":        menuHerb,
	"herv":        menuHerb,
	menuMelon:     menuMelon,
	"メロン":         menuMelon,
	"melon":       menuMelon,
	"meron":       menuMelon,
	menuGinger:    menuGinger,
	"ジンジャー":       menuGinger,
	"ジンジャ":        menuGinger,
	"ginger":      menuGinger,
	menuChoco:     menuChoco,
	"チョコ":         menuChoco,
	"チョコレート":      menuChoco,
	"choco":       menuChoco,
	"choko":       menuChoco,
	"chocolate":   menuChoco,
	menuCaramel:   menuCaramel,
	"キャラメル":       menuCaramel,
	"カラメル":        menuCaramel,
	"caramel":     menuCaramel,
	menuEspresso:  menuEspresso,
	"エスプレ":        menuEspresso,
	"espresso":    menuEspresso,
	menuAmericano: menuAmericano,
	"アメリカン":       menuAmericano,
	"america":     menuAmericano,
	"americano":   menuAmericano,
	"american":    menuAmericano,
	menuMatcha:    menuMatcha,
	"抹茶":          menuMatcha,
	"茶":           menuMatcha,
	"matcha":      menuMatcha,
	"macha":       menuMatcha,
}

// ホットメニュー
var hotMap = map[string]struct{}{
	menuCoffee:    {},
	menuLatte:     {},
	menuMocha:     {},
	menuSoy:       {},
	menuChai:      {},
	menuHerb:      {},
	menuChoco:     {},
	menuCaramel:   {},
	menuEspresso:  {},
	menuAmericano: {},
	menuMatcha:    {},
}

// アイスメニュー
var iceMap = map[string]struct{}{
	menuCoffee:  {},
	menuLatte:   {},
	menuMocha:   {},
	menuSoy:     {},
	menuChai:    {},
	menuCaramel: {},
	menuMelon:   {},
	menuGinger:  {},
	menuMatcha:  {},
}

// グランデサイズメニュー
var largeSizeMap = map[string]struct{}{
	menuCoffee: {},
	menuLatte:  {},
	menuMatcha: {},
	menuMelon:  {},
	menuSoy:    {},
}

var (
	// ホットとアイスの両方が存在するメニューをinit時に設定
	bothHeatMap = map[string]struct{}{}
	// ホット・アイスを含めた全メニューをinit時に設定
	menus = []string{}
)

func init() {
	// メニューの重複防止用キャッシュ
	itemMap := make(map[string]struct{})
	for _, name := range menuMap {
		if _, ok := itemMap[name]; ok {
			continue
		}

		itemMap[name] = struct{}{}

		// グランデサイズ対応
		menuName := name
		if _, ok := largeSizeMap[name]; ok {
			menuName = fmt.Sprintf("* %s", name)
		}

		// ホットとアイスの両方が存在するメニューを判定
		if _, ok := hotMap[name]; ok {
			if _, ok := iceMap[name]; ok {
				menus = append(menus, fmt.Sprintf("%s (hot)", menuName))
				menus = append(menus, fmt.Sprintf("%s (ice)", menuName))
				bothHeatMap[name] = struct{}{}
				continue
			}
		}
		menus = append(menus, menuName)
	}
	sort.Strings(menus)
}

// ユーザーが入力したメニュー名から統一されたメニュー名を返却する
func getMenuName(item string) (string, bool) {
	item = toKatakana(item)
	if name, ok := menuMap[item]; ok {
		return name, true
	}
	return item, false
}

// 全メニューを返却する
func getAllMenu() string {
	return strings.Join(menus, "\n")
}

// ユーザーが入力した「ホット・アイス」のマッピング
var heatMap = map[string]string{
	"hot":  "ホット",
	"熱い":   "ホット",
	"熱":    "ホット",
	"ice":  "アイス",
	"冷たい":  "アイス",
	"冷":    "アイス",
	"コールド": "アイス",
}

// ユーザーが入力した値から ホット or アイス を返却する
func getHeat(heat string) string {
	heat = toKatakana(heat)
	if name, ok := heatMap[heat]; ok {
		return name
	}
	return heat
}

// ホット or アイス が含まれているかどうか判定する
func hasHeat(heat string) bool {
	if heat == "ホット" {
		return true
	}
	return heat == "アイス"
}

// 単一の商品でホットとアイスの2種類が存在するかどうか判定する
func hasBothHeat(item string) bool {
	_, ok := bothHeatMap[item]
	return ok
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
