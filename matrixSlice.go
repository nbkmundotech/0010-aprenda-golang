package main

import (
  "fmt"
)

func main() {
  tabuleiro := [][]string{
    []string{"X", "-", "-"},
    []string{"O", "X", "O"},
    []string{"-", "O", "X"},
  }

  for i := 0; i < len(tabuleiro); i++ {
    //tabuleiro[i]
    for j := 0; j < len(tabuleiro[i]); j++ {
      fmt.Println(tabuleiro[i][j])
    }
  }
}
