package main

import (
  "fmt"
  "math"
)

func main() {
  var total float64 = 1;
  if math.Mod(total, 100) == 1 {
    fmt.Println(total)
  }
}