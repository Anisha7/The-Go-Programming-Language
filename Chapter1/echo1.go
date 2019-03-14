package main
import (
	"fmt"
	"os"
)

func main() {
	// two variables s, sep and type string for both
	// if not initialized, numerics are 0 and strings are "" initially
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}