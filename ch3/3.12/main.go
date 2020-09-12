package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please enter two strings!")
		return
	}
	var isAna = ""
	if !isAnagram(os.Args[1], os.Args[2]) {
		isAna = " not"
	}
	fmt.Printf("%s and %s are%s anagrams of each other.\n", os.Args[1], os.Args[2], isAna)
}

func isAnagram(s1, s2 string) bool {
	m := make(map[rune]int)
	for _, r := range s1 {
		m[r]++
	}
	for _, r := range s2 {
		m[r]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
