package main

import (
	"fmt"
	"math"
	"math/rand"
)

func randtape() {
	var alpha int64
	var beta int64
	var gama int64
	//var A, B int64
	//A = int64(i*7) % 89
	//B = int64(i+13) % 13
	alpha = int64(rand.Intn(10) + rand.Intn(100))
	beta = ((alpha * 3) + 7) / 13
	gama = beta % 20
	gama = int64(math.Abs(float64(gama)))
	//fmt.Println(gama)
	if gama < 5 {
		fmt.Printf("A")
	} else {
		if gama > 5 && gama < 10 {
			fmt.Printf("C")
		} else {
			if gama > 10 && gama < 15 {
				fmt.Printf("G")
			} else {
				fmt.Printf("T")
			}
		}
	}

}
func main() {
	for i := 0; i < 64; i++ {
		for j := 0; j < 64; j++ {
			randtape()
		}
		fmt.Println()
	}

}
