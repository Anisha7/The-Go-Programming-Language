package main

import (
	"fmt"
	"time"

	"./popcount"
)

func main() {
	var x uint64 = 100

	fmt.Println("Testing Version 1")
	start := time.Now()
	fmt.Println(popcount.PopCount(x))
	fmt.Printf("%.6fs elapsed\n", time.Since(start).Seconds())

	fmt.Println("Testing Version 2")
	start = time.Now()
	fmt.Println(popcount.PopCount2(x))
	fmt.Printf("%.6fs elapsed\n", time.Since(start).Seconds())

	fmt.Println("Testing Version 3")
	start = time.Now()
	fmt.Println(popcount.PopCount3(x))
	fmt.Printf("%.6fs elapsed\n", time.Since(start).Seconds())

	fmt.Println("Testing Version 4")
	start = time.Now()
	fmt.Println(popcount.PopCount4(x))
	fmt.Printf("%.6fs elapsed\n", time.Since(start).Seconds())

}
