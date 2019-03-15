// Echo4 prints its command-line arguments

// flag package is used to alert users with a
// message if invalid arguments, flags, or -h/-help are used.

package main

import (
	"flag"
	"fmt"
	"strings"
)

// creates a flag variable of type bool
var n = flag.Bool("n", false, "omit trailing newline") // name of flag, default, message for incorrect val
// string variable
var sep = flag.String("s", " ", "separator") // name, default, and message

// n, sep are pointers to flag variables and must be accessed as *n and *sep

func main() {
	flag.Parse() // to update flag variables from default before using, program exits if flag finds error
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
