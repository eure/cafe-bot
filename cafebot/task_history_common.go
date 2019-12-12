package cafebot

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var histOnce sync.Once
var globalHistory *History

func getGlobalHistory() *History {
	histOnce.Do(func() {
		globalHistory = &History{}
	})
	return globalHistory
}

// 注文履歴用の構造体.
type History struct {
	Max    int
	muList sync.Mutex
	List   []Order
	Number int64
}

func (h *History) GetMax() int {
	if h.Max > 0 {
		return h.Max
	}
	const defaultMax = 100
	return defaultMax
}

func (h *History) Append(user, item string) {
	h.appendOrder(newOrder(user, item))
}

func (h *History) appendOrder(o Order) {
	o.Number = atomic.AddInt64(&h.Number, 1)

	h.muList.Lock()
	defer h.muList.Unlock()

	max := h.GetMax()
	if len(h.List) < max {
		h.List = append([]Order{o}, h.List...)
		return
	}

	h.List = append([]Order{o}, h.List[0:max-1]...)
	return
}

func (h *History) ListOrdersByPage(page int, items ...int) []Order {
	const defaultShowSize = 10
	showSize := defaultShowSize
	if len(items) != 0 {
		showSize = items[0]
	}
	pageFirst := (page) * showSize
	pageLast := (page + 1) * showSize

	size := len(h.List)
	if size <= pageFirst {
		return nil
	}

	if size < pageLast {
		pageLast = size
	}
	return h.List[pageFirst:pageLast]
}

// 単一の注文履歴用の構造体.
type Order struct {
	Time   time.Time
	Number int64
	User   string
	Item   string
}

func newOrder(user, item string) Order {
	return Order{
		Time: time.Now(),
		User: user,
		Item: item,
	}
}

func (o Order) String() string {
	return fmt.Sprintf("[#%d] [%s]\t%s: %s", o.Number, o.Time.Format("2006-01-02 15:04:05"), o.User, o.Item)
}
