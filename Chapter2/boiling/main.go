// Boiling prints the boiling point of water.
package main

import "fmt"

// global constant variable, aailable to all files of the package
const boilingF = 212.0

func main() {
	var f = boilingF // variables local to function
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g degree F or %g degree C\n", f, c)
	// Output:
	// boiling point = 212 degree F or 100 degree C
}
