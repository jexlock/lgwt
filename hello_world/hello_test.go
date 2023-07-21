package main

import "testing"

func TestHello(t *testing.T) {
    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Trinity", "English")
        want := "Hello, Trinity"
        assertCorrectMessage(t, got, want)   
    })

    t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
        got := Hello("", "English")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)   
    })

    t.Run("in spanish", func(t *testing.T) {
        got := Hello("Bruno", "Spanish")
        want := "Hola, Bruno"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in french", func(t *testing.T) {
        got := Hello("Gigi", "French")
        want := "Bonjour, Gigi"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in japanese", func(t *testing.T) {
        got := Hello("Issei", "Japanese")
        want := "Kenichiwa, Issei"
        assertCorrectMessage(t, got, want)
    })

}

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
