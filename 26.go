package main

import (
	"fmt"
	"strings"
)

func testUnique(s string) bool {
	s = strings.ToLower(s)
	m := make(map[string]int)
	for _, value := range s {
		if _, ok := m[string(value)]; ok {
			return false
		}
		m[string(value)]++
	}
	return true
}

func main() {
	str := "abcde"
	str1 := "abCdefAaff"
	str2 := "абвгдее"
	fmt.Println(str, "-", testUnique(str))
	fmt.Println(str1, "-", testUnique(str1))
	fmt.Println(str2, "-", testUnique(str2))
}
