package main

import "fmt"

func main() {
  // var nomeDaVariavel tipoDaVariavel
  // nomeDaVariavel = ValorDaVariavel
  //
  // var altura, idade int = 166, 35
  //
  // var (
  //   idade = 35
  //   altura = 166
  //   nome string = ""
  // )

  idade, altura, nome := 35, 166, "MeuNome"

  fmt.Println("O meu nome é", nome, "e tenho", idade, "anos")
  fmt.Println("Minha altura é", altura)
}
