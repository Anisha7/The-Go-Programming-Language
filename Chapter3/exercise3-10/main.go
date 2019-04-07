// a non-recursive version of comma, using bytes.Bugger instead of string concatenation

package main

import (
	"bytes"
	"fmt"
	"os"
)

func comma(s string) string {
	var buf bytes.Buffer
	for i := range s {
		if i != 0 && i%3 == 0 {
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
