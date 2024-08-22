package main

// [$]> go run stdlib.go

import (
  "fmt"
  "sort"
  "strings"
)

func main() {
  strings_lib()
  sort_lib()
}

func strings_lib() {
  greeting := "hello there friends!"

  fmt.Println(strings.Contains(greeting, "hello!"))
  fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
  fmt.Println(strings.ToUpper(greeting))
  fmt.Println(strings.Index(greeting, "ll")) // 2
  fmt.Println(strings.Split(greeting, " "))
}

func sort_lib() {
  ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

  sort.Ints(ages)
  fmt.Println(ages) // [20 25 30 35 45 50 60 75]

  index := sort.SearchInts(ages, 30)
  fmt.Println(index) // 2

  names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

  sort.Strings(names)
  fmt.Println(names) // [bowser luigi mario peach yoshi]
  fmt.Println(sort.SearchStrings(names, "bowser")) // 0
}
