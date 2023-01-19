package main

import "testing"

func TestHello(t *testing.T) {
  checkCorrectMessage := func (t *testing.T, result, expected string) {
    t.Helper()
    if result != expected {
      t.Errorf("Got \"%s\", expected \"%s\"", result, expected)
    }
  }

  t.Run("should print hello if no nome is given", func (t *testing.T) {
    result := Hello("", "")
    expected := "Hello"

    checkCorrectMessage(t, result, expected)
  })

  t.Run("should print hola for spanish", func (t *testing.T) {
    result := Hello("João", "es")
    expected := "Hola, João"

    checkCorrectMessage(t, result, expected)
  })

  t.Run("should print boujour for french", func (t *testing.T) {
    result := Hello("João", "fr")
    expected := "Bonjour, João"

    checkCorrectMessage(t, result, expected)
  })
}
