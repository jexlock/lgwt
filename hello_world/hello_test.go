package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello("Trinity")
    want := "Hello, Trinity"
    
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
