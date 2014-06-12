package main

import (
  "fmt"
  "fatsecret"
)

func main() {
	foods, err := fatsecret.SearchFood("banana")
	if err != nil {
  }

  for _, food := range foods.FoodList {
    fmt.Printf("%s\n", food)
  }
}

