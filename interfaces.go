package main

import (
  "fmt"
  "math"
)

type Geometrica interface {
  area() float64
}

type Quadrado struct {
  lado float64
} // area = lado ^ 2         (lado * lado)

func (q Quadrado) area() float64 {
  return q.lado * q.lado
}

type Circulo struct {
  raio float64
} // area = pi * (raio ^ 2)     (raio * raio)

func (c Circulo) area() float64 {
  return math.Pi * c.raio * c.raio
}

func main() {
  var g Geometrica
  g = Quadrado{3}

  fmt.Printf("A area do quadrado é %v\n", g.area())

  g = Circulo{5}

  fmt.Printf("A area do circulo é %v\n", g.area())
}
