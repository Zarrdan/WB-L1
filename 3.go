package main

/* Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений. */

import (
	"fmt"
	"sync"
)

func main() {
	array := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	sum := 0
	for _, value := range array { // запускаю цикл, в value передаю значения из массива array
		wg.Add(1)
		go func(a int) { // в анонимной функции считаю сумму квадратов каждого array[value]
			sum += a * a
			wg.Done()
		}(value)
	}
	wg.Wait() // жду завершение горутин
	fmt.Println(sum)
}
