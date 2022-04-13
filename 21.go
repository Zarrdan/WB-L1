package main

import "fmt"

type dog struct{}
type cat struct{}

type animal interface {
	say()
}

type sleepcat interface {
	sleep()
}

func (d *dog) say() {
	fmt.Println(" Dog : gav gav")
}

func (c *cat) sleep() {
	fmt.Println(" Cat : sleep sleep")
}

type animaladapter struct {
	sleepcat sleepcat
}

func newAdapter(s sleepcat) *animaladapter {
	return &animaladapter{
		sleepcat: s,
	}
}

func (w *animaladapter) say() {
	w.sleepcat.sleep()
}

func test(a animal) {
	a.say()
}

func main() {
	dog := &dog{}
	cat := &cat{}

	dogAdapter := newAdapter(cat)
	test(dog)
	//test(cat)
	test(dogAdapter)
}
