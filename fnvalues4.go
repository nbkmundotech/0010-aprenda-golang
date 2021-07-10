package main

import "fmt"

func adicionador() func(int) int {
  var soma int = 0
  //closure
  return func(a int) int {
    soma = soma + a
    return soma
  }
}

func main() {
  var ad1 = adicionador()
  fmt.Println(ad1(21))
  fmt.Println(ad1(13))
  fmt.Println("-----")
  var ad2 = adicionador()
  fmt.Println(ad2(7))
  fmt.Println(ad1(0))
}
