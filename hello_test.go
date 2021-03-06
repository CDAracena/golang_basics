package main

import "testing"

func TestHello(t *testing.T) {

  assertCorrectMessage := func(t *testing.T, got, want string) {
    t.Helper()
    if got != want {
      t.Errorf("got %q want %q", got, want)
    }
  }

  t.Run("Saying hello to ppl", func(t *testing.T) {
    got := Hello("Christian", "")
    want := "Hello, Christian"
    assertCorrectMessage(t, got, want)
  })

  t.Run("Say 'Hello, World' when an empty string is provided", func(t *testing.T) {
    got := Hello("", "")
    want := "Hello, World"
    assertCorrectMessage(t, got, want)
  })

  t.Run("in Spanish", func(t *testing.T) {
    got := Hello("Elodie", "Spanish")
    want := "Hola, Elodie"
    assertCorrectMessage(t, got, want)
  })
  t.Run("in French", func(t *testing.T) {
    got := Hello("Pierre", "French")
    want := "Bonjour, Pierre"
    assertCorrectMessage(t, got, want)
  })
}
