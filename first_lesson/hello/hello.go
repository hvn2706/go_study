package main

import (
    "fmt"
    "log"
    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    people := []string{"hvn", "darlin", "HN17"}

    messages, err := greetings.Hellos(people)

    if (err != nil) {
        log.Fatal(err)
    }

    fmt.Println("Messages: \n", messages)
}
