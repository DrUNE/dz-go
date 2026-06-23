package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Привет Андрей!")
	userHeight := 1.92
	userKg := 92.0
	IMT := userKg / math.Pow(userHeight, 2)
	fmt.Println(IMT)
}
