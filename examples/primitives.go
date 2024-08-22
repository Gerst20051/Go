package main

// [$]> go run primitives.go

import (
  "fmt"
  "time"
)

func main() {
  firstName, lastName := name()
  fmt.Println(firstName == "andrew")
  fmt.Println(lastName == "gerst")
  myAge := age(1993)
  fmt.Println(myAge == 31)
  hobbies()
  foods()
}

func name() (string, string) {
  var firstName string = "andrew"
  lastName := "gerst"
  fmt.Println(fmt.Sprintf("Hello %s %s", firstName, lastName))
  return firstName, lastName
}

func age(dob int) (int) {
  year, _, _ := time.Now().Date()
  age := year - dob
  fmt.Println("You are", age, "years old")
  return age
}

func hobbies() {
  var hobbies [5]string = [5]string{"Bodybuilding", "Running", "Hiking", "Nutrition", "Beach"}
  top2 := [2]string{"Bodybuilding", "Beach"}
  for i := 0; i < len(hobbies); i++ {
    fmt.Println("hobby:", hobbies[i])
  }
  for i := 0; i < len(top2); i++ {
    fmt.Println("top hobby:", top2[i])
  }
}

func foods() {
  var foods = []string{"Steak", ""}
  foods[1] = "Chicken"
  foods = append(foods, "Eggs", "Milk", "Butter", "Salt")
  top3 := foods[0:3]
  for _, food := range foods {
    fmt.Println("food:", food)
  }
  for index, value := range top3 {
    fmt.Printf("top food %v is %v\n", index + 1, value)
  }
}
