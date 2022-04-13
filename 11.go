package main

import "fmt"

/* Реализовать пересечение двух неупорядоченных множеств. */

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6}
	arr2 := []int{4, 5, 6}
	newArr := make([]int, 0)

	for _, val1 := range arr1 { // запускаю цикл по первому массиву
		for _, val2 := range arr2 { //  в цикле новый цикл по второму массиву
			if val1 == val2 { // если есть совпадение чисел первого и второго массива, добавляю это число в новый срез
				newArr = append(newArr, val1)
			}
		}
	}
	//for _, val1 := range arr1 {
	//	newArr = append(newArr, val1)
	//}
	//for _, val2 := range arr2 {
	//	newArr = append(newArr, val2)
	//}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(newArr)
}
