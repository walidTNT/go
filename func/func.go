// func.go
package main

import "fmt"

func Swap(a, b int) (int, int) {
	return b, a
}

func Sum(a ...int) int {
	var sum int
	for _, v := range a {
		sum += v
	}
	return sum
}

type FunkHandler interface {
	ImFunky()
}

type Funky func(string) string

func (f Funky) ImFunky() {
	s := f("toto")
	fmt.Println(s)
}

func GetFunky(f FunkHandler) {
	f.ImFunky()
}

func main() {
	f := func(name string) string {
		return "Hi " + name
	}
	GetFunky(Funky(f))

	x, y := 20, 30
	fmt.Println(x, y)

	y, x = Swap(y, x)
	fmt.Println(x, y)

	s := Sum(5, 10, 5, 20)
	fmt.Println(s)
}
