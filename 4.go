package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/* Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

func work(id int, ch chan int) { // функция печати воркеров, value передаются из главной функции по каналу ch
	for {
		value := <-ch
		fmt.Printf("Worker #%d is doing %d \n", id, value)
	}
}

func main() {

	var N int
	fmt.Print("Введите количество воркеров: ")
	//	fmt.Fscan(os.Stdin, &N)
	fmt.Scan(&N)
	ch := make(chan int)

	for i := 1; i < N+1; i++ { // запускаю горутины воркеров передаю номер воркера и канал ch
		go work(i, ch)
	}

	sigChan := make(chan os.Signal, 1)     //инициализация канала, принимающего сигналы от ОС
	signal.Notify(sigChan, syscall.SIGINT) // уведомление о прерывании в канал

	func() {
		var i int
		for {
			i++
			select {
			case <-sigChan:
				return
			default:
				time.Sleep(time.Second)
				ch <- i
			}
		}
	}()
	close(ch)
	close(sigChan)
	fmt.Print("stop")
}
