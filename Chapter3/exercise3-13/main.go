package main

import "fmt"

// in bytes
const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB
	TiB //overflow
	PiB
	EiB
	ZiB //overflow
	YiB
)

// in KB
const (
	_  = 1 << (10 * iota) // 1024
	KB                    // 1
	MB                    // 1024
	GB
	TB // overflow
	PB
	EB
	ZB // overflow
	YB
)

func main() {
	fmt.Printf("%s %s %s %s %s %s %s %s", KiB, MiB, GiB, TiB, PiB, EiB, ZiB, YiB)
	fmt.Printf("%s %s %s %s %s %s %s %s", KB, MB, GB, TB, PB, EB, ZB, YB)
}
