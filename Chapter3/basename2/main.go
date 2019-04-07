// basename1 removes directory components and a .suffix
// without the help of libraries
package main

import (
	"fmt"
	"os"
	"strings"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
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
