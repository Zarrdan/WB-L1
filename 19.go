package main

import "fmt"

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
	/*arr := []rune(str)
	res := make([]rune, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		res = append(res, arr[i])
	}
	return string(res) */
}

func main() {
	str := "главрыба"
	strRev := reverse(str)
	fmt.Println(str)
	fmt.Println(strRev)
}
