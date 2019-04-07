// basename1 removes directory components and a .suffix
// without the help of libraries
package main

import (
	"fmt"
	"os"
)

func basename(s string) string {
	// Discard last '/' and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// Preserve everything before last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func main() {
	fmt.Println(os.Args[1])
	val := os.Args[1:]
	result := "Provide a valid string to alter"
	if len(val) == 1 {
		result = basename(val[0])
	}
	fmt.Println(result)
}
