package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	d := Dictionary{
		"test": "it's just a test",
	}

	t.Run("should return the value of a known word in a dictionary", func(t *testing.T) {
		got, err := d.Search("test")
		want := "it's just a test"
		AssertErrorNil(t, err)
		AssertStrings(t, want, got)
	})

	t.Run("should return an error if a key is not found", func(t *testing.T) {
		_, err := d.Search("unknown word")
		AssertErrorNotNil(t, err)
		AssertError(t, ErrNotFound, err)
	})
}

func TestAdd(t *testing.T) {
	d := Dictionary{}
	t.Run("should add a new word to the dictionary", func(t *testing.T) {
		want := "this is a new word added"
		d.Add("new", want)
		got, err := d.Search("new")
		AssertErrorNil(t, err)
		AssertStrings(t, want, got)
	})

	t.Run("should return an error when add a word that already exists in the dictionary", func(t *testing.T) {
		word := "word"
		definition := "already exists"
		d := Dictionary{word: definition}
		err := d.Add("word", "is this a new definition?")
		AssertErrorNotNil(t, err)
		AssertError(t, ErrExistentWord, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("should be able to update an existent word", func(t *testing.T) {
		word := "word"
		definition := "already exists"
		new_definition := "this a new definition"
		d := Dictionary{word: definition}

		err := d.Update(word, new_definition)
		AssertErrorNil(t, err)

		got, _ := d.Search(word)
		AssertStrings(t, new_definition, got)
	})

	t.Run("should be able to update an existent word", func(t *testing.T) {
		word := "word"
		new_definition := "this a new definition"

		d := Dictionary{}
		err := d.Update(word, new_definition)

		AssertError(t, ErrWordDoesNotExist, err)
	})
}

func TestDelete(t *testing.T) {
	word := "word"
	definition := "already exists"
	d := Dictionary{word: definition}

	d.Delete(word)

	got, err := d.Search(word)
	AssertStrings(t, got, "")
	AssertError(t, ErrNotFound, err)
}

func AssertStrings(t testing.TB, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func AssertErrorNil(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("expected error to be nil")
	}
}

func AssertErrorNotNil(t testing.TB, got error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected an error, got nil")
	}
}

func AssertError(t testing.TB, want, got error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, expected %d", got, want)
	}
}
