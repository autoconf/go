package main

import (
  "fmt"
  "time"
)
func main() {
  fmt.Printf("Hello world.\n")
  fmt.Println("this is exercise alpha")
  now := time.Now()

  fmt.Println(now)
  fmt.Println(now.Second())
}
