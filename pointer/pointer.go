// les pointers

package main

import "fmt"

func main() {
	x := 52
	y := &x // y == 0xc42000a350
	z := &y // z == 0xc42002a020

	fmt.Println(*y)
	fmt.Println(**z)
	fmt.Printf("%p, %p, %p", &x, &y, &z)
}
