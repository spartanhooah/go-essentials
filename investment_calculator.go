package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	// var investmentAmount, years, expectedReturnRate float64
	// var investmentAmount float64 = 1000
	// years := 10.0
	// expectedReturnRate := 5.5
	// Could combine the above three lines like the below; stylistic choice
	// investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	investmentAmount := prompt("Initial investment amount: ")
	expectedReturnRate := prompt("Expected return rate: ")
	years := prompt("Length of the investment in years: ")

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future value without inflation: %.2f\n", futureValue)
	fmt.Print(formattedFV)
	fmt.Printf("Future value after inflation: %.2f", futureRealValue)

	// Can use back-ticks for multi-line strings. In this case, \n is printed as \n rather than a line break
}

func prompt(prompt string) float64 {
	fmt.Print(prompt)

	var input float64
	fmt.Scan(&input)

	return input
}

func calculateFutureValues(investment, returnRate, years float64) (float64, float64) {
	fv := investment * math.Pow(1+returnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
}
