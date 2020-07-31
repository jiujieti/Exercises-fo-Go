// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func popCount1(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func popCount2(x uint64) int {
	var count uint64
	for i := 0; i < 64; i++ {
		count += (x & 1)
		x = x >> 1
	}
	return int(count)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("please provide a number")
		os.Exit(1)
	}
	s := os.Args[1]
	n, e := strconv.ParseUint(s, 10, 64)
	if e != nil {
		fmt.Printf("error %v ocrrured when parsing int\n", e)
		os.Exit(1)
	}
	meature(popCount1, n)
	meature(popCount2, n)
}

func meature(f func(x uint64) int, x uint64) {
	start := time.Now()
	var n int
	for i := 0; i < 10000; i++ {
		n = f(x)
	}
	fmt.Printf("popcount=%d takes %v\n", n, time.Since(start)/10000)
}
