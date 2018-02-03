package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please give two strings.")
		return
	}
	isAna := isAnagram(os.Args[1], os.Args[2])
	fmt.Printf("%t\n", isAna)
}

func isAnagram(s1, s2 string) bool {
	m := make(map[rune]int)
	for _, s := range s1 {
		m[s]++
	}
	for _, s := range s2 {
		m[s]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
