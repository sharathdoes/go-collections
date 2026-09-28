package stack

import "testing"


func TestPeek(t *testing.T) {
	stack := NewStack()
	stack.Push(10)
	value, err := stack.Peek()

	if err != nil {
		t.Fatal("unexpected error")
	}

	if value != 10 {
		t.Error("expected 10")
	}
}

func TestPush(t *testing.T) {
	stack := NewStack()

	stack.Push(10)

	if len(stack.items) != 1 {
		t.Errorf("expected 1 item, got %d", len(stack.items))
	}
}

func TestPopOrder(t *testing.T) {
	stack := NewStack()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	expected := []int{30, 20, 10}

	for _, want := range expected {
		got, err := stack.Pop()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	}
}