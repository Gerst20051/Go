package main

// [$]> go run pointers.go

import (
  "fmt"
)

func updateName(x *string) { // accepts a pointer to a string
  *x = "wedge"
}

func main() {
  name := "tifa"

  fmt.Println("memory address of name is:", &name)

  m := &name // m is storing the pointer (memory address) to name

  fmt.Println("memory address:", m)
  fmt.Println("value at memory address:", *m) // dereference it to get the value of the memory address it points to

  fmt.Println(name) // prints "tifa"

  updateName(m)

  fmt.Println(name) // prints "wedge"
}
