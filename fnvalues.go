package main

import "fmt"

func main() {
  somar := func(a, b float64) float64 {
    return a + b
  }
  multiplicar := func(a, b float64) float64 {
    return a * b
  }

  fmt.Println(somar(3, 4))
  fmt.Println(multiplicar(3, 4))
}
