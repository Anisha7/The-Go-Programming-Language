// this program formats a non-negative decimal integer string by inserting commas
package main

import (
	"fmt"
	"os"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	fmt.Printf("Given: %s\n", os.Args[1])
	val := os.Args[1:]
	result := "Provide a valid integer string to alter"
	if len(val) == 1 {
		result = comma(val[0])
	}
	fmt.Printf("Result: %s\n", result)
}
