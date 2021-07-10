package main

import (
  "fmt"
  "strings"
)

type Pessoa struct {
  Nome string
  Sobrenome string
}

func (p Pessoa) NomeCompleto() string {
  return p.Nome + " " + p.Sobrenome
}

func (p *Pessoa) CapitalizarNome() {
  p.Nome = strings.ToUpper(p.Nome)
  fmt.Printf("Dentro de CapitalizarNome: %s\n", p.Nome)
}

func main() {
  alguem := Pessoa{"José", "de Alencar"}

  alguem.CapitalizarNome()
  fmt.Println(alguem.NomeCompleto())
}
