package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	i := 0
	z := 1.0
	oz := z
	//for i := 0; i < 10; i++ {
	//	z -= (z*z - x) / (2 * z)
	//}
	for {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-oz) < 0.00000000000001 {
			break
		}
		oz = z
		i++
	}
	return z, i
}

func main() {
	fmt.Println(Sqrt(2))
}
