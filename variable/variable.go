package main

// les variables
import (
	"fmt"
	"log"
)

// Definition d'une variable statique
var prenom string

var (
	version  string
	username string
)

const tva = 20

const (
	x = 52
	y = 23
)

const (
	a1 = 1 << (1 * iota)
	a2
	a3
	a4
	a5
	a6
)

func main() {
	fmt.Println("Les variables...")
	fmt.Println(a1, a2, a3, a4, a5, a6)

	prenom = "toto"
	fmt.Println("Hello", prenom)

	name := "aston"
	fmt.Println(name)

	a, b, c := 1, 2, "hello"
	fmt.Println(a, b, c)

	// Gestion des erreurs en retour de fonction
	n, err := fmt.Print("Super Toto\n")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n)

	nb, _ := fmt.Println("test")
	fmt.Println(nb)
}
