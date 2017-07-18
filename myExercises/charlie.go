package main

import "fmt"

func main() {
	var i int = 1
	for i < 50 {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				break
			}
			if j == i-1 {
				fmt.Println(i)
				break
			}
		}
		i++
	}
}
