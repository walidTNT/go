package main

import (
	"fmt"
	"strings"
	"time"
)

// carriage return
func main() {
	for i := 0; i <= 50; i++ {
		fmt.Printf("\rLoading... %d/50%s ", i, "%")
		time.Sleep(time.Millisecond * 20)
	}

	fmt.Println()

	for i := 0; i <= 50; i++ {
		str := strings.Repeat("=", 50-i)
		fmt.Printf("\r[%s] %d/50%s ", str, i, "%")
		time.Sleep(time.Millisecond * 20)
	}

	fmt.Println("\nDone.")
}
