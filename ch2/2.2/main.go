//!+

// converts a numeric argument to Kilograms and Pounds
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"./weightconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			convert(arg)
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			convert(input.Text())
		}
	}
}

func convert(v string) {
	n, err := strconv.ParseFloat(v, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "KgP: %v\n", err)
		os.Exit(1)
	}
	kg := weightconv.Kilograms(n)
	p := weightconv.Pounds(n)
	fmt.Printf("%s = %s, %s = %s\n", kg, weightconv.KgToP(kg), p, weightconv.PToKg(p))
}

//!-
