package cafebot

import (
	"strings"
)

type showHistoryTask struct {
	history *History
	page    int
	showFn  func(history string) error
}

// NewShowHistoryTask is a task to add order history.
func NewShowHistoryTask(page int, showFn func(history string) error) *showHistoryTask {
	return NewShowHistoryTaskWithHistory(getGlobalHistory(), page, showFn)
}

// NewShowHistoryTaskWithHistory is a task to add order history.
func NewShowHistoryTaskWithHistory(history *History, page int, showFn func(history string) error) *showHistoryTask {
	return &showHistoryTask{
		history: history,
		page:    page,
		showFn:  showFn,
	}
}

func (showHistoryTask) GetName() string {
	return "show_history_task"
}

func (t showHistoryTask) Run() error {
	orders := t.history.ListOrdersByPage(t.page)
	result := make([]string, len(orders))
	for i, hist := range orders {
		result[i] = hist.String()
	}

	return t.showFn(strings.Join(result, "\n"))
}
