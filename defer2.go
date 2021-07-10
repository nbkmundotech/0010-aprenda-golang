package main

import "fmt"

func imprimir() string {
  fmt.Println("Imprimindo...")
  return "VALOR de IMPRIMIR"
}

func main() {
  defer fmt.Println(imprimir())
  fmt.Println("2")
  fmt.Println("3")
}


//stack
// topo
fmt.Println(imprimir())
fmt.Println(imprimir())
fmt.Println(imprimir())

// fundo

// execucao:
// fmt.Println("3")
// fmt.Println("2")
// fmt.Println("1")
