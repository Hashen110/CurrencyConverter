package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("------------------")
	fmt.Println("Currency Converter")
	fmt.Println("------------------")
	usd := 0.0051
	euro := 0.0042
	fmt.Println("\nUSD Rate: ", usd)
	fmt.Println("Euro Rate: ", euro)
	fmt.Println("\nHow do you need to convert")
	fmt.Println("\t1. LKR to (USD, EURO)")
	fmt.Println("\t2. (USD, EURO) to LKR")
	fmt.Print("\nChoice: ")
	var covertChoice int
	_, choiceError := fmt.Scanf("%d", &covertChoice)
	if choiceError != nil {
		fmt.Println("Error: ", choiceError)
		fmt.Println("Try Again Later!")
	} else {
		if covertChoice == 1 {
			fmt.Print("\nEnter amount to convert = ")
			var num float64
			_, err := fmt.Scanf("%f", &num)
			if err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Println("\nConverted amount ( LKR to USD ) = ", math.Floor(usd*num*10000)/10000)
				fmt.Println("Converted amount ( LKR to EURO ) = ", math.Floor(euro*num*10000)/10000)
			}
		} else if covertChoice == 2 {
			fmt.Print("\nEnter amount to convert = ")
			var num float64
			_, err := fmt.Scanf("%f", &num)
			if err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Println("\nConverted amount ( USD to LKR ) = ", num*195.08)
				fmt.Println("Converted amount ( EURO to LKR ) = ", num*237.22)
			}
		} else {
			fmt.Println("Invalid Choice")
		}
	}
}
