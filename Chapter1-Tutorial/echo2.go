package main
import (
	"fmt"
	"os"
)

func main() {
	// two variables s, sep and type string for both
	// if not initialized, numerics are 0 and strings are "" initially
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}