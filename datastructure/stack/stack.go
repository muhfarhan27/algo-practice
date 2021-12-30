package stack

import (
	"container/list"
)

// Source = https://golangbyexample.com/stack-in-golang/
// Stack custom implementation from Linked List

type CustomStack struct {
	stack *list.List
}

func (c *CustomStack) New() *CustomStack {
	return &CustomStack{list.New()}
}

func (c *CustomStack) Push(value interface{}) {
	c.stack.PushFront(value)
}

func (c *CustomStack) Pop() (interface{}, bool) {
	ele := c.stack.Front()
	if ele != nil {
		c.stack.Remove(ele)
		return ele.Value, true
	}
	return nil, false
}

func (c *CustomStack) Top() (interface{}, bool) {
	ele := c.stack.Front()
	if ele != nil {
		return ele.Value, true
	}
	return nil, false
}

func (c *CustomStack) Size() int {
	return c.stack.Len()
}

func (c *CustomStack) Empty() bool {
	return c.stack.Len() == 0
}
