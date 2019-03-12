// 4. Modify dup2 to print the names of all files in which each duplicated line occurs
// Dup2 prints the count and text of lines that appear more than once in the input.
// It reads from stdin or from a list of named files.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// stores line:file
	locations := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		var s string = "command-line"
		countLines(os.Stdin, counts, locations, s)
	} else {
		for _, arg := range files {
			// os.Open returns the open file (*os.File) 
			// and err (==nil if file is opened successfully)
			f, err := os.Open(arg)
			if err != nil {
				// %v displays values of any type
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, locations, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t\t from file: %s\n", n, line, locations[line])
		}
	}
}

// reads the file and counts its lines into a map line:count
// map is a reference to the data structure created by make, 
// so the changes can be seen by the caller's map reference too (in this case, by main())
func countLines(f *os.File, counts map[string]int, locations map[string]string, arg string) {
	input := bufio.NewScanner(f)
	var line string
	for input.Scan() {
		line = input.Text()
		counts[line]++
		locations[line] = arg
	}
	// Note: ignoring potential errors from input.Err()
}