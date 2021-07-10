package main

import "fmt"

func main() {
  var notas map[string]int
  notas = map[string]int{
    "Ana": 9,
    "Maria": 10,
  }

  // notas = make(map[string]int)
  // notas["Ana"] = 9
  // notas["Maria"] = 10

  fmt.Println(notas)
}
