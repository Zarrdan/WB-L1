package main

import "fmt"

func main() {
	arr := []int{-5, -3, 0, 2, 10, 50}

	i, message := search(arr, 0)
	fmt.Printf("%s %d", message, i)

}
func search(arr []int, target int) (int, string) {
	var min int
	max := len(arr) - 1

	for min <= max {
		mid := (max + min) / 2

		switch {
		case arr[mid] == target:
			return mid, "Found in position:"
		case arr[mid] > target:
			max = mid - 1
		case arr[mid] < target:
			min = mid + 1
		}
	}
	return -1, "Not found"
}
