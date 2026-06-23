package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64
	var userKg float64
	fmt.Print("--- Калькулятор індекса маси тіла ---\n")
	fmt.Print("Введіть свій ріст в метрах ")
	fmt.Scan(&userHeight)
	fmt.Print("Введіть свою вагу ")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight, IMTPower)
	fmt.Print("Ваш індекс маси тіла: ")
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
