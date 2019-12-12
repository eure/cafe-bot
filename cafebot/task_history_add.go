package cafebot

type addHistoryTask struct {
	history *History
	user    string
	item    string
}

// NewAddHistoryTask is a task to add order history.
func NewAddHistoryTask(user, item string) *addHistoryTask {
	return NewAddHistoryTaskWithHistory(getGlobalHistory(), user, item)
}

// NewAddHistoryTaskWithHistory is a task to add order history.
func NewAddHistoryTaskWithHistory(history *History, user, item string) *addHistoryTask {
	return &addHistoryTask{
		history: history,
		user:    user,
		item:    item,
	}
}

func (addHistoryTask) GetName() string {
	return "add_history_task"
}

func (t addHistoryTask) Run() error {
	t.history.Append(t.user, t.item)
	return nil
}
