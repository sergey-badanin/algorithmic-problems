package linkedlistmiddle

import (
	"testing"
)

func TestMiddleNode(t *testing.T) {
	given := buildList(3)
	want := 2
	if got := middleNode(&given).Val; want != got {
		t.Errorf("For %v got %v, want %v", given, got, want)
	}

	given = buildList(4)
	want = 3
	if got := middleNode(&given).Val; want != got {
		t.Errorf("For %v got %v, want %v", given, got, want)
	}

	given = buildList(5)
	want = 3
	if got := middleNode(&given).Val; want != got {
		t.Errorf("For %v got %v, want %v", given, got, want)
	}

	given = buildList(1)
	want = 1
	if got := middleNode(&given).Val; want != got {
		t.Errorf("For %v got %v, want %v", given, got, want)
	}

	given = buildList(6)
	want = 4
	if got := middleNode(&given).Val; want != got {
		t.Errorf("For %v got %v, want %v", given, got, want)
	}

	given = buildList(2)
	want = 2
	if got := middleNode(&given).Val; want != got {
		t.Errorf("For %v got %v, want %v", given, got, want)
	}
}

func buildList(length int) ListNode {
	head := ListNode{1, nil}
	next := &head
	for i := 2; i <= length; i++ {
		next.Next = &ListNode{i, nil}
		next = next.Next
	}
	return head
}
