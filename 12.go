package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	var res []string
	newArr := make(map[string]int)

	for _, val := range arr {
		if newArr[val] != 1 {
			newArr[val] = 1
			res = append(res, val)
		}
	}
	fmt.Println(res)
}
