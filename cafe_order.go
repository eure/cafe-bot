package main

import (
	"fmt"
	"time"
)

// 注文履歴用の構造体.
type Order struct {
	Time time.Time
	User string
	Item string
}

func newOrder(user, item string) Order {
	return Order{
		Time: time.Now(),
		User: user,
		Item: item,
	}
}

func (o Order) String() string {
	return fmt.Sprintf("[%s] %s: %s", o.Time.Format("2006-01-02 15:04:05"), o.User, o.Item)
}
