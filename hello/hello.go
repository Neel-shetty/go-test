package main

import (
    "fmt"
    "neel/greetings"
		"lib/test"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("neel")
    fmt.Println(message)
		num := Add(1)
}
