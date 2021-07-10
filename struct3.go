package main

import "fmt"

func main() {
  type Posicao struct {
    X int
    Y int
  }

  var pos Posicao = Posicao{47, 81}
  var k *Posicao = &pos
  k.Y = 33
  fmt.Println(     k.Y    )
}
