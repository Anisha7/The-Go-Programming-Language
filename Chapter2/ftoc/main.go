// Ftoc prints two Fahrenheit-to-Celsius conversions.
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g degree F = %g degree C\n", freezingF, fToC(freezingF)) // "32 degree F = 0 degree C"
	fmt.Printf("%g degree F = %g degree C\n", boilingF, fToC(boilingF))   // "212 degree F = 100 degree C"
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
