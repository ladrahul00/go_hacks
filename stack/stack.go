package stack

import "fmt"

type Stack struct {
	top    int
	values []T
}

type T interface {
}

func (stack Stack) Push(value T) Stack {
	if len(stack.values) == 0 {
		stack.top = -1
	}
	stack.values = append(stack.values, value)
	stack.top++
	fmt.Println("Stack push: ", value)
	fmt.Println("Stack : ", stack)

	return stack
}

func (stack Stack) Pop() (Stack, T) {
	var pop_val T
	if len(stack.values) > 0 {
		pop_val = stack.values[len(stack.values)-1]
		stack.values = stack.values[:len(stack.values)-1]
		stack.top--
	} else {
		stack.top = -1
		return stack, nil
	}

	return stack, pop_val
}
