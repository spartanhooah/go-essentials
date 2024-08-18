package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64
	// var investmentAmount float64 = 1000
	// years := 10.0
	// expectedReturnRate := 5.5
	// Could combine the above three lines like the below; stylistic choice
	// investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	fmt.Print("Initial investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Length of the investment in years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future value without inflation: %.2f\n", futureValue)
	fmt.Print(formattedFV)
	fmt.Printf("Future value after inflation: %.2f", futureRealValue)

	// Can use back-ticks for multi-line strings. In this case, \n is printed as \n rather than a line break
}
