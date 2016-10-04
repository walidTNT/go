// map == tableau associatif
package main

import "fmt"

func main() {

	m := map[string]string{
		"prenom": "Chuck",
		"nom":    "Norris",
		"age":    "inconu",
	}

	fmt.Println(m["prenom"])

	m2 := map[string]bool{
		"a": true,
		"b": true,
		"c": false,
	}

	fmt.Println(m2)

	m3 := map[string]interface{}{
		"prenom":   "John",
		"nom":      "Doe",
		"age":      50,
		"isActive": true,
	}

	// interface assertion
	var prenom string
	prenom, ok := m3["prenom"].(string)
	if !ok {
		panic("age is not string")
	}
	fmt.Println(prenom)

	var version interface{}
	version = "v0.1.0"

	var getVersion string
	getVersion = version.(string)

	fmt.Println(getVersion)
}
