package main

import (
	"fmt"
	"math"
)

func main() {
	userHeight := 1.8
	var userKg float64
	userKg = 70
	var IMT = userKg / math.Pow(userHeight, 2)
	fmt.Print(IMT)
	if IMT > 30 {
		fmt.Print(" You're fat")
	}
	if IMT < 18 {
		fmt.Print(" You're slime")
	}
	if IMT > 18.5 {
		fmt.Print("You're a nice guy")
	}
}
