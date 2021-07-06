package main

import "testing"

func TestToPigLatin(t *testing.T) {
	t.Run("starts with singular consonant", func(t *testing.T) {
		want := "ellohay orldway"
		got, err := toPigLatin("hello world")
		if err != nil {
			t.Error(err)
		}
		if got != want {
			t.Errorf("Got %v, Want %v\n", got, want)
		}
	})
	t.Run("starts with consonant cluster", func(t *testing.T) {
		want := "ellsmay outshay ancefray isphyllay"
		got, err := toPigLatin("smell shout france phyllis")
		if err != nil {
			t.Error(err)
		}
		if got != want {
			t.Errorf("Got %v, Want %v\n", got, want)
		}
	})
	t.Run("starts with vowel", func(t *testing.T) {
		want := "antyay eatersyay indulgeyay overyay umbrellasyay"
		got, err := toPigLatin("ant eaters indulge over umbrellas")
		if err != nil {
			t.Error(err)
		}
		if got != want {
			t.Errorf("Got %v, Want %v\n", got, want)
		}
	})
}

func TestIsVowel(t *testing.T) {
	t.Run("true case", func(t *testing.T) {
		want := true
		got := isVowel('a')
		if got != want {
			t.Errorf("Got %v, Want %v\n", got, want)
		}
	})
	t.Run("false case", func(t *testing.T) {
		want := false
		got := isVowel('b')
		if got != want {
			t.Errorf("Got %v, Want %v\n", got, want)
		}

	})
}

func TestIsAlphabetic(t *testing.T) {
	t.Run("true case", func(t *testing.T) {
		want := true
		got := isAlphabetic("abcdefghijklmnopqrstuvwxyz")
		if got != want {
			t.Errorf("Got %v, Want %v\n", got, want)
		}
	})
	t.Run("false case", func(t *testing.T) {
		want := false
		got := isAlphabetic("1234567890")
		if got != want {
			t.Errorf("Got %v, Want %v\n", got, want)
		}
	})
}
