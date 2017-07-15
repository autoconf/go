package main

import "fmt"
// ** herhangi bir sayıda sonuç döndürülebilir **
func swap (x, y string) (string, string){
  return y, x
}

func main()  {
  a, b := swap("hello", "world")
  fmt.Println(a, b)
}
