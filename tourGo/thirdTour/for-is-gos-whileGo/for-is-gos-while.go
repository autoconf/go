package main

import "fmt"

func main() {
  sum := 1
/*
!!! go'da for harici döngü yapısı yok. !!!
    for deyimindeki ';' kullanmazsak go için bu while deyimine karşılık gelir
*/
  for sum < 1000 {
    sum += sum
  }
  fmt.Println(sum)
}
