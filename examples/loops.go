package main

// [$]> go run loops.go

import (
  "fmt"
)

func main() {
  x := 0
  for x < 5 {
    fmt.Println("value of x is:", x)
    x++
  }

  for i := 0; i < 5; i++ {
    fmt.Println("value of i is:", i)
  }
}
