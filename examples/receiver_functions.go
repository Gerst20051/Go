package main

// [$]> go run receiver_functions.go

import "fmt"

type bill struct {
  name string
  items map[string]float64
  tip float64
}

func newBill(name string) bill {
  b := bill{
    name: name,
    items: map[string]float64{"pie": 5.99, "cake": 3.99},
    tip: 0,
  }

  return b
}

func main() {
  mybill := newBill("mario's bill")

  mybill.addItem("onion soup", 4.50)
  mybill.addItem("coffee", 3.25)
  mybill.updateTip(10)

  fmt.Println(mybill.format())
}

// receiver function (limit function to bill objects)
func (b *bill) format() string {
  fs := "Bill breakdown: \n"
  var total float64 = 0

  for k, v := range b.items {
    fs += fmt.Sprintf("%-25v ...$%v\n", k + ":", v)
    total += v
  }

  fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

  fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total + b.tip)

  return fs
}

func (b *bill) updateTip(tip float64) {
  // (*b).tip = tip
  b.tip = tip // automatically deferences pointers for structs
}

func (b *bill) addItem(name string, price float64) {
  b.items[name] = price
}
