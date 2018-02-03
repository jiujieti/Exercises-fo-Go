// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var sign, intPart, fracPart, str string
	if s[0] == '+' || s[0] == '-' {
		sign = s[:1]
		s = s[1:]
	}
	temp := strings.Split(s, ".")
	intPart = temp[0]
	if len(temp) > 1 {
		fracPart = temp[1]
	}
	for n := len(intPart); n >= 0; n -= 3 {
		if n-3 <= 0 {
			str = intPart[0:n] + str
			break
		}
		str = "," + intPart[n-3:n] + str
	}
	if fracPart != "" {
		str += "."
		for n := 0; n < len(fracPart); n += 3 {
			if n+3 >= len(fracPart) {
				str += fracPart[n:len(fracPart)]
				break
			}
			str += fracPart[n:n+3] + ","
		}
	}
	if sign != "" {
		str = sign + str
	}
	return str
}

//!-
