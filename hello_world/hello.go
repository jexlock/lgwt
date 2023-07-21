package main

const (
    french = "French"
    japanese = "Japanese"
    spanish = "Spanish"

    englishHelloPrefix = "Hello, " 
    spanishHelloPrefix = "Hola, " 
    frenchHelloPrefix = "Bonjour, " 
    japaneseHelloPrefix = "Kenichiwa, "
)


func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }
    
    return getGreetingPrefix(language) + name
}

func getGreetingPrefix(language string) (prefix string) {
    switch language {
    case french:
        prefix = frenchHelloPrefix
    case spanish:
        prefix = spanishHelloPrefix
    case japanese:
        prefix = japaneseHelloPrefix
    default:
        prefix = englishHelloPrefix
    }
    return 
}
