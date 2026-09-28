package queue

import (
	"testing"
)

func TestInsertIntoQueue(t *testing.T){
	queue:=NewQueue()
	queue.Push(10)
	queue.Push(20)
	queue.Push(30)

	x, err:=queue.Peek()
	if err!=nil {
		t.Fatalf("unable to fetch peek")
	}

	if x!=10{
		t.Error("doesnt work")
	}
}