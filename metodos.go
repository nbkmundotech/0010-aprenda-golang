package main

import (
  "fmt"
)

type Pessoa struct {
  Nome string
  Sobrenome string
}

func (p Pessoa) NomeCompleto() string {
  return p.Nome + " " + p.Sobrenome
}

func main() {
  alguem := Pessoa{"José", "de Alencar"}

  fmt.Println(alguem.NomeCompleto())
}
