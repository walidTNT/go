// poo/poo.go
package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name            string // ""
	Age             uint8  // 0
	Email, Password string
	Connected       bool // false
}

// SetName... à toi de jouer, je crois en toi :D
func (u *User) SetName(n string) {
	u.Name = n
}

func (u User) IsConnected() bool {
	return u.Connected
}

type Admin struct {
	User
	Root bool
}

func (a Admin) IsRoot() bool {
	return a.Root
}

// *User == j'attend en paramètre une adresse mémoire
func SayHelloToUser(u *User) {
	u.Name = strings.ToUpper(u.Name)
	fmt.Println("Hello", u.Name)
}

func main() {
	user := User{
		Name:     "toto",
		Age:      0,
		Email:    "toto@gmail.com",
		Password: "****",
	}
	user.Connected = true
	SayHelloToUser(&user)

	//fmt.Println(user)
	//fmt.Printf("%v", user.IsConnected())
	//fmt.Printf("%p, %p", &user, &user.Name)

	admin := Admin{user, true}
	fmt.Println(admin)

	user.SetName("Steven Seagle")
	fmt.Println(user.Name)
}
