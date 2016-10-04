// les slices == tableau à taille dynamique

package main

import "fmt"

func main() {
	s := "Lorem ipsum lores ecatum"

	fmt.Printf("%s\n", s[0:5])  // Lorem
	fmt.Printf("%s\n", s[6:12]) // ipsum
	fmt.Printf("%s\n", s[6:])   // ipsum lores ecatum
	fmt.Printf("%s\n", s[:12])  // Lorem ipsum
	fmt.Printf("%s\n", s)

	a := []float64{
		52.36,
		89.54,
		78.15,
	}

	a[1] = 59.545
	a = append(a, 598.64)

	fmt.Println(a)

}
