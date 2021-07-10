package main

import "fmt"

func main() {
  defer fmt.Println("Oi")
  defer fmt.Println("Tudo bem?")
}
