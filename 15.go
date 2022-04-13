package main

import "fmt"

// Го кодирует строку как массив байтов. Но не различает строки ASCII  и Unicode.
// Так как некоторые символы занимают 2 или 3 байта(нужно переводить символы в руны).

func createHugeString(size int) (s string) {
	for i := 0; i < size; i++ {
		s += "А"
	}
	return
}

// Вывод первых 10 символов
func func1() {
	v := createHugeString(1 << 10)
	a := []rune(v)

	justString := string(v[:10])
	justStringRune := string(a[:10])

	fmt.Println(justString)
	fmt.Println(justStringRune)
}

func main() {
	func1()
}
