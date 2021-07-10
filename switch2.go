package main

import "fmt"

func main() {
  var nota int = 7

  // if (nota > 9) {
  //
  // } else if (nota > 7) {
  //
  // } else if (nota > 6) {
  //
  // } else {
  //
  // }

  switch {
  case nota > 9:
    fmt.Println("Ótimo")
  case nota > 7:
    fmt.Println("Muito bem")
  case nota > 6:
    fmt.Println("Bom")
  default:
    fmt.Println("Péssimo")
  }
}
