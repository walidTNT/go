// type/type.go
package main

import (
	"fmt"
	"reflect"
)

// string
// bool
// int

// Str is a string
type Str string

// Print is a printer of Str
func (s Str) Printer() {
	fmt.Println(s)
}

// Les entiers, byte, flottants, ...
// power mode
func main() {
	var s Str
	var s2 string
	s = "My Str"
	s2 = "Test string"
	fmt.Println(s + Str(s2))

	s.Printer()

	var a uint8
	var b int

	a = 10
	b = 20

	fmt.Println("Les entiers")
	fmt.Println(a + uint8(b))

	fmt.Println("Les Floats")
	var f float32
	f = 5.556
	fmt.Println(f + float32(b))

	fmt.Println("Les Booleans")
	var Boolean bool
	Boolean = true
	fmt.Println(Boolean)

	fmt.Println("Les Bytes")
	MyByte := byte('a')
	fmt.Println(MyByte)

	fmt.Println("Les Strings")
	MyBytes := []byte("hello")
	fmt.Println(MyBytes)

	MyBytes2 := []byte{'a', 'b'}
	fmt.Println(MyBytes2)

	byte2 := byte(108)
	fmt.Println(string(byte2))

	message := "à bientôt"
	fmt.Println(len([]rune(message)))

	fmt.Println("Show type")

	// Première méthode
	fmt.Printf("Type: %T\n", message)

	// Deuxième méthode
	fmt.Println(reflect.TypeOf(message))
}
