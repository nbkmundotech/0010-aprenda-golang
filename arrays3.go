package main

import "fmt"

func main() {
  // Arrays
  // 1, 1, 2, 3, 5, 8, 13
  // var numeros = [7]int{1, 1, 2, 3, 5, 8}
  numeros := []int{1, 1, 2, 3, 5, 8, 13}

  for indice := 0; indice < len(numeros); indice++ {
    fmt.Println(numeros[indice])
  }
}
