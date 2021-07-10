package main

import "fmt"

func main() {
  type Posicao struct {
    X int
    Y int
  }

  var pos Posicao = Posicao{Y: 47, X: 21}
  fmt.Println(pos)
}
