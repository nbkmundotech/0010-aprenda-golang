package main

import "fmt"

func main() {
  var nomes = []string{
    "Ana",   // [0]
    "Jose",  // [1]
    "Maria", // [2]
  }

  nomes = nomes[1:]

  fmt.Println(nomes)
  fmt.Printf("len=%d, cap=%d\n", len(nomes), cap(nomes))

  // fmt.Println("len=", len(nomes))   // => 3
  // fmt.Printf("%T\n", nomes)
}
