package main

import "fmt"

type Perfil struct {
  Altura float64
  Ativo bool
  Profissao string
}

func main() {
  var perfis map[string]Perfil = make(map[string]Perfil)

  perfis["Joao"] = Perfil{
    1.74, true, "Medico",
  }
  perfis["Maria"] = Perfil{
    1.63, false, "Engenheira",
  }

  fmt.Println(perfis)
}
