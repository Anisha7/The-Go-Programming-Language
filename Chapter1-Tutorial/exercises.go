package main
import (
	"fmt"
	"os"
	"strings"
	"time"
)
// 1. Write echo program that also prints os.Args[0] (the name of the command)
func printing() {
	fmt.Println(strings.Join(os.Args, " "))
}
// 2. Modify to print index and value of each arg one per line
func printAll() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}
}
// 3. Measure the difference in run time with echo2.go and an implementation that uses strings.Join
func slowPrinting() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func main() {
	fmt.Println("Exercise 1: ")
	printing()
	fmt.Println("")

	fmt.Println("Exercise 2: ")
	printAll()
	fmt.Println("")

	fmt.Println("Exercise 3: ")
	fmt.Println("SLOW TIME:")
	start := time.Now()
	slowPrinting()
	fmt.Printf("%.9fs elapsed\n", time.Since(start).Seconds())
	fmt.Println("")
	fmt.Println("FAST TIME:")
	start = time.Now()
	printing()
	fmt.Printf("%.9fs elapsed\n", time.Since(start).Seconds())

}