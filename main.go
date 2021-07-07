package piglatin

import (
	"errors"
	"strings"
)

const alpha = "abcdefghijklmnopqrstuvwxyz"
const vowels = "aeiou"

//ToPigLatin takes a string and returns its pig-latin translation.
func ToPigLatin(s string) (string, error) {
	words := strings.Fields(s)
	if !isSliceAlphabetic(words) {
		return "", errors.New("input must be strictly alphabetic")
	}
	var pl []string
	for _, word := range words {
		if isVowel(word[0]) {
			pl = append(pl, word+"yay")
		} else {
			consonants := ""
			for _, c := range word {
				if isVowel(byte(c)) {
					break
				}
				consonants += string(c)
			}
			pl = append(pl, strings.TrimPrefix(word, consonants)+consonants+"ay")
		}
	}
	return strings.Join(pl, " "), nil
}

//isVowel returns a boolean that states whether a byte is a vowel (excluding Y for weird pig latin reasons)
func isVowel(b byte) bool {
	return strings.Contains(vowels, string(b))
}

//isSliceAlphabetic works similar to isAlphabetic but for slices
func isSliceAlphabetic(s []string) bool {
	for _, e := range s {
		if !isAlphabetic(e) {
			return false
		}
	}
	return true
}

//isAlphabetic returns a boolean that states whether the strings contains strictly alphabetic characters (A-Z + a-z)
func isAlphabetic(s string) bool {
	for _, c := range s {
		if !strings.Contains(alpha, strings.ToLower(string(c))) {
			return false
		}
	}
	return true
}
