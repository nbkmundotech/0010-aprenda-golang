package main

import (
  "fmt"
)

func main() {
  var animes [3]string;
  animes = [3]string{"Dragon Ball", "Sailor Moon", "Yuyu Hakusho"}

  var doisPrimeiros = animes[:2]
  var doisUltimos = animes[1:]

  fmt.Println(animes)
  fmt.Println(doisPrimeiros)
  fmt.Println(doisUltimos)
  fmt.Printf("len = %d     cap = %d\n", len(doisPrimeiros), cap(doisPrimeiros))
  fmt.Printf("len = %d     cap = %d\n", len(doisUltimos), cap(doisUltimos))
}
