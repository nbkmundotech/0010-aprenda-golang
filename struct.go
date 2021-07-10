package main

import "fmt"

func main() {
  type Posicao struct {
    X int
    Y int
  }
  pos := Posicao{40, 67}
  pos.Y = 51

  fmt.Println(pos.X)
}
