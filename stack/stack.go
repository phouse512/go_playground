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

func New(maxSize int) (*Stack, error) {
	if maxSize < 1 {
		return nil, errors.New("must be at least size 1")
	}

	stack := new(Stack)
	stack.array = make([]int, maxSize)
	stack.size = maxSize
	stack.front = -1
	return stack, nil
}

// push takes a value to add to the stack. it will return an error
//  as its second return value if it is already full, otherwise it
//  returns the pointer to the stack location
func (s *Stack) push(value int) (int, error) {
	if s.front == len(s.array) {
		return s.front, errors.New("stack is at max size")
	}

	s.front = s.front + 1
	s.array[s.front] = value

	return s.front, nil
}

// pop returns a value off the top of the stack, if it doesn't exist,
//  it will return an error.
func (s *Stack) pop() (int, error) {
	if s.front < 0 {
		return -1, errors.New("underflow error, no data")
	}

	to_return := s.array[s.front]
	s.front = s.front - 1
	return to_return, nil
}

func main() {
	test, _ := New(5)
	pointer, _ := test.push(21)
	pointer, _ = test.push(31)

	should_be, _ := test.pop()
	should_be_2, _ := test.pop()
	fmt.Println(test, pointer)
	fmt.Println("should be a 31 but it is a..", should_be)
	fmt.Println("should be a 21 but it is a..", should_be_2)

	pointer, _ = test.push(28)
	pointer, _ = test.push(29)
	pointer, _ = test.push(30)
	pointer, _ = test.push(31)
	pointer, _ = test.push(32)

	should_be, _ = test.pop()
	should_be_2, _ = test.pop()
	should_be_3, _ := test.pop()
	should_be_4, _ := test.pop()
	should_be_5, _ := test.pop()

	fmt.Println("should be a 28 but it is a..", should_be_5)
	fmt.Println("should be a 29 but it is a..", should_be_4)
	fmt.Println("should be a 30 but it is a..", should_be_3)
	fmt.Println("should be a 31 but it is a..", should_be_2)
	fmt.Println("should be a 32 but it is a..", should_be)
}
