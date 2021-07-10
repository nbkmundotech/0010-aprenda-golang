package main

import "fmt"

func calcular(a int) (quadrado int, cubo int) {
  quadrado = a * a
  cubo = a * a * a

  return
}

func main() {
  fmt.Println(calcular(2))
}
