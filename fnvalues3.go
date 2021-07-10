package main

import "fmt"

// nome: Joao ---> Ola, Joao!
func cumprimentar(nome string) func() string {
  return func() string {
    return fmt.Sprintf("Ola, %s!", nome)
  }
}

func main() {
  cumprimentarJoao := cumprimentar("Joao")

  fmt.Println(cumprimentarJoao())

  cumprimentarMaria := cumprimentar("Maria")

  fmt.Println(cumprimentarMaria())
}
