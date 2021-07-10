package main

import "fmt"

func computar(fn func(float64, float64) float64, fn2 func(float64, float64) float64) float64 {
  // 11  +  12
  return fn(5, 6) + fn2(3,4)
}

func main() {
  somar := func(a, b float64) float64 {
    return a + b
  }
  multiplicar := func(a, b float64) float64 {
    return a * b
  }

  fmt.Println(computar(somar, multiplicar))
}
