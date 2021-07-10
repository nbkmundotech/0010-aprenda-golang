package main

import (
  "fmt"
)

type Potencia interface {
  Quad() int
}

type Inteiro int

func (i Inteiro) Quad() int {
  return int(i * i)
}

type Outro string

func (o Outro) Quad() int { // "abcd"
  return len(o) * len(o)
}

func descrever(potencia Potencia) {
  fmt.Printf("%v / %T\n", potencia, potencia)
}

func main() {
  // var num Inteiro = 3

  var pot Potencia          // (valor, tipo)
  pot = Inteiro(3)

  var out Potencia
  out = Outro("abcd")

  fmt.Println(pot.Quad()) // => 9
  descrever(pot)

  fmt.Println(out.Quad())
  descrever(out)
}
