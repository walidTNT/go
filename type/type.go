// type/type.go
package main

import "fmt"

// Les entiers, byte, flottants, ...
// power mode
func main() {
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

}
