package main

import (
    "fmt"
    "neel/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("neel")
    fmt.Println(message)
}
