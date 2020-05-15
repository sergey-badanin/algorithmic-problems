package minstack

import "testing"

func TestStackOperations(t *testing.T) {
	stack := Constructor()

	stack.Push(-2)
	want := -2
	if got := stack.Top(); got != want {
		t.Errorf("Got %v, want %v", got, want)
	}

	stack.Push(0)
	want = 0
	if got := stack.Top(); got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
	want = -2
	if got := stack.GetMin(); got != want {
		t.Errorf("Got %v, want %v", got, want)
	}

	stack.Push(-3)
	want = -3
	if got := stack.Top(); got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
	want = -3
	if got := stack.GetMin(); got != want {
		t.Errorf("Got %v, want %v", got, want)
	}

	stack.Pop()
	want = 0
	if got := stack.Top(); got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
	want = -2
	if got := stack.GetMin(); got != want {
		t.Errorf("Got %v, want %v", got, want)
	}

	stack.Pop() //comming to empty stack doesn't panic
}
