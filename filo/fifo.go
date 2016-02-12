package main

import "fmt"

type Queue struct {
    queue []int
}

func (q *Queue) enqueue(value int) bool {
    q.queue = append(q.queue, value)
    return true
}

func (q *Queue) dequeue() int {
    if len(q.queue) < 1 {
        return -1
    }
    val := q.queue[0]
    q.queue = q.queue[1:]
    return val
}

func (q Queue) peek() int {
    if len(q.queue) < 1 {
        return -1
    }
    return q.queue[0]
}

func main() {
    test := Queue{[]int{1,2,3,4}}
    fmt.Println(test)
    test.enqueue(29)
    fmt.Println(test)
    val := test.dequeue()
    fmt.Println(test)
    fmt.Println(val)
    fmt.Println(test.peek())
}
