// loop.go
package main

import "fmt"

func main() {

	// Boucle classique (for)
	for i := 0; i < 5; i++ {
		fmt.Printf("#%d\n", i)
	}

	// for {
	// 	fmt.Println("test")
	// }

	// for == while
	counter := 0
	for counter < 10 {
		fmt.Println("For: ", counter)
		counter++
	}

	m := map[string]string{
		"ville":       "Paris",
		"pays":        "France",
		"departement": "Ile de France",
	}

	// Equivalent du foreach
	for k, v := range m {
		fmt.Println(k, v)
	}
}
