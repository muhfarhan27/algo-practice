package queue

import (
	"container/list"
)

// Source = https://golangbyexample.com/queue-in-golang/
// Queue custom implementation from Linked List

type CustomQueue struct {
	queue *list.List
}

func New() *CustomQueue {
	return &CustomQueue{list.New()}
}

func (c *CustomQueue) Enqueue(value interface{}) {
	c.queue.PushFront(value)
}

func (c *CustomQueue) Dequeue() (interface{}, bool) {
	ele := c.queue.Back()
	if ele != nil {
		c.queue.Remove(ele)
		return ele.Value, true
	}
	return nil, false
}

func (c *CustomQueue) Front() (interface{}, bool) {
	ele := c.queue.Back()
	if ele != nil {
		return ele.Value, true
	}
	return nil, false
}

func (c *CustomQueue) Size() int {
	return c.queue.Len()
}

func (c *CustomQueue) Empty() bool {
	return c.queue.Len() == 0
}
