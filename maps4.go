package main

import "fmt"

type Perfil struct {
  Altura float64
  Ativo bool
  Profissao string
}

func main() {
  var perfis map[string]Perfil = map[string]Perfil{
    "Joao": {
      1.74, true, "Medico",
    },
    "Maria": {
      1.63, false, "Engenheira",
    },
  }

  fmt.Println(perfis)
}
