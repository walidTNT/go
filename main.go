package main

// /usr/local/go/src
import "aston/test"

func main() {
    test.Name = "John"
    test.SayHello()

    test.Name = "toto"
}
