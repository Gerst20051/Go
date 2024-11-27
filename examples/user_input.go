package main

// [$]> go run user_input.go

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

type bill struct {
  name string
  items map[string]float64
  tip float64
}

func newBill(name string) bill {
  b := bill{
    name: name,
    items: map[string]float64{},
    tip: 0,
  }

  return b
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
  fmt.Print(prompt)
  input, err := r.ReadString('\n')

  return strings.TrimSpace(input), err
}

func createBill() bill {
  reader := bufio.NewReader(os.Stdin)

  name, _ := getInput("Create a new bill name: ", reader)

  b := newBill(name)
  fmt.Println("Created the bill -", b.name)

  return b
}

func promptOptions(b bill) {
  reader := bufio.NewReader(os.Stdin)

  opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

  switch opt {
  case "a":
    name, _ := getInput("Item name: ", reader)
    price, _ := getInput("Item price: ", reader)

    fmt.Println(name, price)
  case "s":
    fmt.Println("you chose s")
  case "t":
    tip, _ := getInput("Enter tip amount ($): ", reader)

    fmt.Println(tip)
  default:
    fmt.Println("that was not a valid option...")
    promptOptions(b)
  }
}

func main() {
  mybill := createBill()
  promptOptions(mybill)
}
