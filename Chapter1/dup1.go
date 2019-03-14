// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.

package main
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// makes map with key string and value int
	// make: built-in function, creates a new empty map
	counts := make(map[string]int)
	// reads input and breaks it into lines or words
	input := bufio.NewScanner(os.Stdin)
	
	// reads the next line each time, and removes new line character
	// returns false when no more input, true otherwise
	for input.Scan() {
		// adds line as key to map with count 1 
		// or increments its count if it exists
		// retrieves result from input.Scan() before moving to next line
		counts[input.Text()]++
	}
	// Note: ignoring potential errors from input.Err()
	// for key, value in map counts
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
