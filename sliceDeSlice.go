package main

import (
  "fmt"
  "strings"
)

func main() {
  tabuleiro := [][]string{
    []string{"X", "-", "-"},
    []string{"O", "X", "O"},
    []string{"-", "O", "X"},
  }

  // fmt.Println(tabuleiro)
  for i := 0; i < len(tabuleiro); i++ {
    // fmt.Println(tabuleiro[i])
    fmt.Printf("%s\n", strings.Join(tabuleiro[i], " "))
  }
}

// [
//  [X - -]
//  [O X O]
//  [- O X]
// ]
