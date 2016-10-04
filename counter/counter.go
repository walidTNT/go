// Counter

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	var text string

	if len(os.Args) > 1 {
		text = strings.ToLower(os.Args[1])
	}

	var voyelles, consonnes int

	for _, c := range text {
		switch c {
		case 'a', 'e', 'i', 'o', 'u', 'y':
			voyelles++
		case ' ', ',', '.', ';':
		default:
			consonnes++
		}
	}

	fmt.Printf("%s", color.CyanString("%s", "Voyelles")+": ")
	fmt.Printf("%d\n", voyelles)
	fmt.Printf("%s", color.GreenString("%s", "Consonnes")+": ")
	fmt.Printf("%d", consonnes)
}
