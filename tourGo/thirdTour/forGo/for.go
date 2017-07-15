package main

import "fmt"

func main()  {
  sum := 0
//!!! for deyimi için () gerekmez ama {} her zaman gereklidir !!!
  for i := 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)
}
