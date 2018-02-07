package main

// 単一タスク用インターフェース.
type task interface {
	// タスク名を取得
	getName() string
	// タスクを実行
	run() error
}
