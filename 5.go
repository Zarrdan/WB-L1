package main

/* Разработать программу, которая будет последовательно
отправлять значения в канал,а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться. */

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var N int
	fmt.Print("Введите время работы в секундах: ")
	fmt.Fscan(os.Stdin, &N)
	ch := make(chan int)
	timer := make(chan int)

	var t int

	go func(ch chan int) {
		for range ch {
			read := <-ch
			fmt.Println(read)
			t += 1
		}
	}(ch)

	go func(timer chan int) {
		time.Sleep(time.Duration(N) * time.Second)
		timer <- 1
	}(timer)

	for {
		time.Sleep(time.Second)
		select {
		case <-timer:
			close(ch)
			return
		default:
			ch <- t
		}
	}
}
