package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var testArr = generate(100)
	measure(myAppend, testArr)
	measure(myJoin, testArr)
}

func measure(f func(args []string), testedStrings []string) {
	var start = time.Now()
	f(testedStrings)
	fmt.Println("The test finished in ", time.Since(start).Seconds())
}

func generate(size int) []string {
	var strArr = make([]string, size)
	for i := range strArr {
		strArr[i] = "test"
	}
	return strArr
}

func myAppend(strArr []string) {
	var s string
	for _, v := range strArr {
		s += " " + v
	}
}

func myJoin(strArr []string) {
	strings.Join(strArr, " ")
}
