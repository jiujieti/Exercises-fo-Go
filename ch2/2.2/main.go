//!+

// converts a numeric argument to Kilograms and Pounds
package main

import (
	"fmt"
	"os"
	"strconv"

	"./weightconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "KgP: %v\n", err)
			os.Exit(1)
		}
		k := weightconv.Kilograms(num)
		p := weightconv.Pounds(num)
		fmt.Printf("%s = %s, %s = %s\n",
			k, weightconv.KgToP(k), p, weightconv.PToKg(p))
	}
}

//!-
