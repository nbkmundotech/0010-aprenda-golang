package main

import (
  "fmt"
)

func main() {
  var nome string = "Joao"

  switch nome {
  case "Ana":
    fmt.Println("É a Ana")
  case "Joao":
    fmt.Println("É o João")
  default:
    fmt.Println("Não conheço")
  }

  fmt.Println(nome)
}
