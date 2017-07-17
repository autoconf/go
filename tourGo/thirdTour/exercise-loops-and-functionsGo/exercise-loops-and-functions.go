package main

import (
	"fmt"
  "math"
)

func sqrt(x float64) float64 {
	var z float64 = 10
	end := 0.0
	for{
		z = z - ((z * z) - x) / (2 * z)
		if math.Abs(end - z) < 0.000001{
			end = z
			break
		}
		end = z
	}
	return end
}

func main() {
	fmt.Println(sqrt(2))
}
