// Package tempconv performs Celsius and Fahrenheit temparature computations
// We turn different temparature scales into different types
// All variables that start with a uppercase letter is exported

package tempconv

// Celsius is a type declaration
type Celsius float64

// Fahrenheit is a type declaration
type Fahrenheit float64

// differentiating these types prevent a Celsius or Fahrenheit type
// to be compared with each other even though both are float64's
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// CToF converts a Celsius value to a Fahrenheit value
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit value to a Celsius value
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
