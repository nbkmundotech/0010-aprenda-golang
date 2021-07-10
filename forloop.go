package main

import "fmt"

func main() {
  // somar numeros inteiros de 1 a 10
  var soma int = 0

  for i := 1; i <= 10; i++ {
    soma += i
  }

  fmt.Println(soma)
}

// i : 11
// soma : 55
