package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	OrderID    string
	ExpireTime time.Time
}

type DelayQueue struct {
	mu    sync.Mutex
	queue []Order
}

func (dq *DelayQueue) AddOrder(order Order) {
	dq.mu.Lock()
	defer dq.mu.Unlock()
	dq.queue = append(dq.queue, order)
}

func (dq *DelayQueue) StartProcessing() {
	for {
		dq.mu.Lock()
		now := time.Now()
		indexesToRemove := []int{}
		for i, order := range dq.queue {
			if order.ExpireTime.Before(now) {
				fmt.Printf("Order %s has expired\n", order.OrderID)
				indexesToRemove = append(indexesToRemove, i)
			}
		}
		// Remove expired orders
		for i := len(indexesToRemove) - 1; i >= 0; i-- {
			index := indexesToRemove[i]
			dq.queue = append(dq.queue[:index], dq.queue[index+1:]...)
		}
		dq.mu.Unlock()
		time.Sleep(1 * time.Second) // Check every second
	}
}

func main() {
	delayQueue := DelayQueue{}

	// Start processing the delay queue
	go delayQueue.StartProcessing()

	// Add orders to the delay queue
	delayQueue.AddOrder(Order{OrderID: "123", ExpireTime: time.Now().Add(10 * time.Second)})
	delayQueue.AddOrder(Order{OrderID: "456", ExpireTime: time.Now().Add(20 * time.Second)})

	// Wait for some time to let orders expire
	time.Sleep(30 * time.Second)
}
