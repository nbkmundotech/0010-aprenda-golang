package main

import "fmt"

func main() {
  // && || !

  //    3 < x < 7
  // x > 3 && x < 7

  var x int = 9;
  fmt.Println("x > 3 && x < 7", x > 3 && x < 7)

  //    -------(3)---(7)----
  fmt.Println("x < 3 || x > 7", x < 3 || x > 7)

  fmt.Println(!true)
}
