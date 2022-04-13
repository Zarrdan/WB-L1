package main

import (
	"fmt"
	"sync"
)

func main() {
	array := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, value := range array {
		wg.Add(1)
		go func(a int) { // передаю значения из массива в внонимную функцию, умножаю каждое значение и вывожу в терминал
			fmt.Println(a * a)
			wg.Done() // уменьшает внутренний счетчик активных элементов на единицу.
		}(value) // передавыемый аргумент

	}
	wg.Wait() // жду завершение горутин
}
