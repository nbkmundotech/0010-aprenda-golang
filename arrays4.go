package main

import "fmt"

func main() {
  var numeros = [7]int{1, 1, 2, 3, 5, 8, 13}
  var soma = 0

  for indice := 0; indice < len(numeros); indice++ {
    soma = soma + numeros[indice]
  }

  fmt.Println(soma)
}
