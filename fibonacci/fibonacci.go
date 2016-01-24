package main

import (
  "fmt"
  "os"
  "strconv"
)

func main() {

    arg := os.Args[1]
    i, _ := strconv.Atoi(arg)
    fmt.Println(fibonacci(i))
}

func fibonacci(count int) int {
    if count == 0 {
        return 0
    } 
    if count == 1 {
        return 1
    }
    return fibonacci(count - 1) + fibonacci(count - 2)
}

