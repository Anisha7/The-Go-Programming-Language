// returns the number of set bits, that is, bits whose value is 1, in a uint64 value (population count)

package popcount

//pc[i] in the population count of i.
var pc [256]byte

func init() {
	// i, _ would also do this, where is index and _ ignores value
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCount2 - Exercise 2.3: Rewrite Popcount to use a loop. Compare performance with v1.
func PopCount2(x uint64) int {
	var val byte
	var i uint
	for i < 8 {
		val += pc[byte(x>>(i*8))]
		i++
	}
	return int(val)
}

// PopCount3 - Exercise 2.4: Write a version of PopCount that counts bits
// by shifting its argument through 64 bit positions, testing the rightmost
// bit each time. Compare performance with v1
func PopCount3(x uint64) int {
	var val byte
	var i uint
	for i < 64 {
		val += pc[byte(x>>(i*64))]
		i++
	}
	return int(val)
}

// PopCount4 - Exercise 2.5: The expression x&(x-1) clears the rightmost
// non-zero bit of x. Write a version of PopCount that counts bits by
// using this fact, and assess its performance
func PopCount4(x uint64) int {
	//TODO: FIX THIS FUNCTION: placeholder code right now
	x = x & (x - 1)
	var i = PopCount3(x)

	return i + 1
}
