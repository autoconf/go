package main

import "fmt"

func main(){
  alpha := 42
  bravo := 3.1415
  charlie := false
  delta := "hello"

  fmt.Printf("alpha is of type %T\n", alpha)
  fmt.Printf("bravo is of type %T\n", bravo)
  fmt.Printf("charlie is of type %T\n", charlie)
  fmt.Printf("delta is of type %T\n", delta)
}
