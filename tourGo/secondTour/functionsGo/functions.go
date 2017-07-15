package main

import(
  "fmt"
)
/*
***değişken türü değişken adından sonra geliyor***
**ardışık olarak aynı türden parametre alıyorsa, en sonda bir kere tanımlamak yeterli**
*/
func add(x, y int) int {
  return x + y
}

func main(){
  fmt.Println(add(42, 13))
}
