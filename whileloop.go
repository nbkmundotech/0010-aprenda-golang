package main

import "fmt"
// import "time"

func main() {
  // queremos somar: 2 + 4 + 8 + 16 + ...
  // 2^1 + 2^2 + 2^3 + 2^4 + ...

  // 2^9 = 512
  // 2^10 = 1024

  var soma int = 2

  for true { // while
    soma += soma
    fmt.Println(soma)
    // time.Sleep(100 * time.Millisecond)
  }

  // 1024
}
