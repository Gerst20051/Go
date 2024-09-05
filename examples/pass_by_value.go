package main

// [$]> go run pass_by_value.go

import (
  "fmt"
)

func updateName(x string) {
  x = "wedge" // only updates the copy of variable
}

func main() {
  name := "tifa"

  updateName(name)

  fmt.Println(name) // prints "tifa"
}
