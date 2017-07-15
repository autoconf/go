package main

import(
  "fmt"
)
//***değişken türü değişken adından sonra geliyor***
func add(x int, y int) int {
  return x + y
}

func main(){
  fmt.Println(add(42, 13))
}
