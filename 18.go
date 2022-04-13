package main

import (
	"fmt"
	"sync"
)

type countSruct struct {
	mx sync.Mutex
	wg sync.WaitGroup
	i  int
}

func newStruct() countSruct {
	return countSruct{}
}

func count(one *countSruct) {
	one.mx.Lock()
	one.i++
	one.mx.Unlock()
	one.wg.Done()
}

func main() {
	one := newStruct()
	fmt.Println("Значение счетчика до вычислений", one.i)

	for i := 0; i < 10; i++ {
		one.wg.Add(1)
		go count(&one)
	}
	one.wg.Wait()
	fmt.Println("Значение счетчика после", one.i)
}
