package main

// [$]> go run structs.go

import "fmt"

type bill struct {
  name string
  items map[string]float64
  tip float64
}

func newBill(name string) bill {
  b := bill{
    name: name,
    items: map[string]float64{}, // {} means empty
    tip: 0,
  }

  return b
}

func main() {
  mybill := newBill("mario's bill")

  fmt.Println(mybill) // prints {mario's bill map[] 0}
}
