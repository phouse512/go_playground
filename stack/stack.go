package main

import "errors"
import "fmt"

// This is a simple array-based implementation of an int stack in Go,
//  written for practice and for fun. All methods won't panic if
//  the stack is full, or if there are no empty values to return,
//  it will simply return an error

type Stack struct {
	array []int
	size  int
	front int
}

func NewStack(maxSize int) (*Stack, error) {
	if maxSize < 1 {
		return nil, errors.New("must be at least size 1")
	}

	stack := new(Stack)
	stack.array = make([]int, maxSize)
	stack.size = maxSize
	stack.front = 0
	return stack, nil
}

func main() {
	test, _ := NewStack(5)

	fmt.Println(test)
}
