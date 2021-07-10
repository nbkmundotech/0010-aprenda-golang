package main

import "fmt"

func main() {
  var nomes = [3]string{
    "Ana",  // [0]
    "Jose", // [1]
    "Maria", // [2]
  }

  var p1 []string = nomes[0:2]
  p1[0] = "Rogerio"
  fmt.Println(p1)
  fmt.Println("Array original:", nomes)
}
