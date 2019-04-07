// a non-recursive version of comma, using bytes.Bugger instead of string concatenation
// modified to deal with floating-point numbers and an optional sign

package main

import (
	"bytes"
	"fmt"
	"os"
)

func comma(s string) string {
	// account for optional sign +/-
	temp := 0
	if s[0] == '+' || s[0] == '-' {
		temp = 1
	}

	var buf bytes.Buffer
	found := false
	for i := range s {
		if s[i] == '.' {
			found = true
		}
		if i != temp && (len(s)-i-temp)%3 == 0 && !found {
			buf.WriteString(",")
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
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
