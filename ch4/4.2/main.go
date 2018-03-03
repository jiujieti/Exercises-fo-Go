// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var s384 = flag.Bool("s384", false, "print out the hash value calculated by SHA384")
var s512 = flag.Bool("s512", false, "print out the hash value calculated by SHA512")
var s = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	var bytes = []byte(flag.Args()[0])
	if *s384 {
		hash := sha512.Sum384(bytes)
		fmt.Printf("%x\n", hash)
		return
	}
	if *s512 {
		hash := sha512.Sum512(bytes)
		fmt.Printf("%x\n", hash)
		return
	}
	hash := sha256.Sum256(bytes)
	fmt.Printf("%x\n", hash)
}

//!-
