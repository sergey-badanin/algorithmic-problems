package minstack

import "container/list"

//Problem:
//	Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
//		push(x) -- Push element x onto stack.
//		pop() -- Removes the element on top of the stack.
//		top() -- Get the top element.
//		getMin() -- Retrieve the minimum element in the stack.

//Constraints:
//	Methods pop, top and getMin operations will always be called on non-empty stacks.

//MinStack represents stack tracking minimal element in it
type MinStack struct {
	stack list.List
	min   list.List
}

//Constructor returns initialized instance of MinStack
func Constructor() MinStack {
	return MinStack{*list.New(), *list.New()}
}

//Push to stack
func (ms *MinStack) Push(x int) {
	ms.stack.PushFront(x)
	if ms.min.Front() == nil || x <= ms.min.Front().Value.(int) {
		ms.min.PushFront(x)
	}
}

//Pop from stack
func (ms *MinStack) Pop() {
	if ms.min.Front() != nil && ms.stack.Front().Value.(int) == ms.min.Front().Value.(int) {
		ms.min.Remove(ms.min.Front())
	}
	ms.stack.Remove(ms.stack.Front())
}

//Top returns last pushed elem
func (ms *MinStack) Top() int {
	return ms.stack.Front().Value.(int)
}

//GetMin returns min elem present in stack
func (ms *MinStack) GetMin() int {
	return ms.min.Front().Value.(int)
}
