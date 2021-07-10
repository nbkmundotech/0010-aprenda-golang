package main

import "fmt"

func main() {
  // if    else        else if

  // var a int = 11

  // a > 10: "a é maior que 10"

  // a <= 10 && a > 5:  "a está entre 6 e 10"

  // a <= 5  "a é menor ou igual a 5"

  if a := 2; a > 10 {
    fmt.Println("a é maior que 10", a)
  } else if a > 5 {
    fmt.Println("a está entre 6 e 10", a)
  } else {
    fmt.Println("a é menor ou igual a 5", a)
  }

  fmt.Println(a)
}
