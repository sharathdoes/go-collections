package queue

import "errors"


type Queue struct {
	items []int
}

func NewQueue() *Queue {
	return &Queue{
		items: []int{},
	}
}

func(q *Queue) Push(x int ) {
	q.items=append(q.items, x)
}

func(q *Queue) Peek() (int, error) {
	if len(q.items)<1 {
		return 0,errors.New("no values in queue")
	}
	return q.items[0], nil
}

func(q *Queue) Pop() (int, error ) {
	if len(q.items)<1 {
		return 0,errors.New("no values in queue")
	}
	value:=q.items[0]
	q.items=q.items[1:]
	return value, nil
}