package main

import "fmt"

const french = "French"
const japanese = "Japanese"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const japaneseHelloPrefix = "Kenichiwa, "

func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    prefix := englishHelloPrefix

    switch language {
    case french:
        prefix = frenchHelloPrefix
    case spanish:
        prefix = spanishHelloPrefix
    case japanese:
        prefix = japaneseHelloPrefix
    }

    return prefix + name
}

func main() {
    fmt.Println(Hello("world", "English"))
}
