// array = tableau à taille fixe
package main

import "fmt"

func main() {
	a := [3]string{"a", "b", "c"}
	//a[4] = "test" pas bien
	a[2] = "Joe"
	fmt.Println(a)

	// La limite se définie en fonction du nombre
	// d'éléments
	b := [...]int{50, 60, 52, 68}
	fmt.Println(b)

	c := [5]bool{}
	fmt.Println(c, len(c), cap(c))

	d := make([]byte, 4)
	d[2] = 'k'
	fmt.Println(d)
}
