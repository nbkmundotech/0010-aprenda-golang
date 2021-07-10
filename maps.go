package main

import "fmt"

func main() {
  var notas map[string]int
  notas = make(map[string]int)

  // "Ana" ---> 9
  // "Maria" ---> 10
  // ...

  notas["Ana"] = 9
  notas["Maria"] = 10

  valor, existe := notas["Joao"]

  if existe {
    fmt.Println(valor)
  }

  // fmt.Println(existe)
}
