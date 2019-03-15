// Package tempconv performs Celsius and Fahrenheit conversions.

package tempconv

import (
	"fmt"

	"gopl.io/ch2/tempconv"
)

// Celsius is a type declaration
type Celsius float64

// Fahrenheit is a type declaration
type Fahrenheit float64

// Kelvin is a type declaration
type Kelvin float64

// differentiating these types prevent a Celsius or Fahrenheit type
// to be compared with each other even though both are float64's
const (
	AbsoluteZeroC Celsius = -273.15 // = 0 Kelvin
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// prints Celsius value as a formatted string
func (c Celsius) String() string {
	// formats according to a format specifier
	return fmt.Sprintf("%g degree C", c)
}

// prints Fahrenheit value as a formatted string
func (f Fahrenheit) String() string {
	// formats according to a format specifier
	return fmt.Sprintf("%g degree F", f)
}

// prints Kelvin value as a formatted string
func (k Kelvin) String() string {
	// formats according to a format specifier
	return fmt.Sprintf("%g degree Kelvin", k)
}

func main() {
	// using functions from conv.go (we added to the same tempconv package)
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC) // "Brrrr! -273.15 degree C"
	fmt.Println(tempconv.CToF(tempconv.BoilingC))     // "212 degree F"
}
