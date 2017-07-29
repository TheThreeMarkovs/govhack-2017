package main

import (
	"math/rand"
	"strings"
)

type Markov struct {
	chain map[string][]rune
	n     int
}

func NewMarkov(n int) *Markov {
	return &Markov{
		chain: map[string][]rune{},
		n:     n,
	}
}

func (m *Markov) ParseWord(word string) {
	runes := []rune(word)
	if m.n > len(runes) {
		return
	}

	end := len(runes) - m.n
	for i := 0; i < end; i++ {
		key := string(runes[i : i+m.n])
		val := runes[i+m.n]
		if _, ok := m.chain[key]; ok {
			m.chain[key] = append(m.chain[key], val)
		} else {
			m.chain[key] = []rune{val}
		}
	}
}

func (m *Markov) GenerateWord(length int) string {
	letters := []rune(m.getRandomPrefix())

	for i := m.n; i <= length; i++ {
		lastLetters := letters[i-m.n : i]
		// fmt.Printf("Looking for key %v\n", string(lastLetters))
		if runes, ok := m.chain[string(lastLetters)]; ok {
			letters = append(letters, getRandomLetter(runes))
		} else {
			return string(letters)
		}
	}

	return string(letters)
}

func (m *Markov) GenerateBusinessName() string {
	numWords := rand.Intn(3) + 1
	words := []string{}

	for i := 0; i < numWords; i++ {
		wordLength := rand.Intn(7) + 4
		word := m.GenerateWord(wordLength)
		word = strings.Title(strings.ToLower(word))
		words = append(words, word)
	}

	return strings.Join(words, " ")
}

func (m *Markov) getRandomPrefix() string {
	// Go randomises the order of maps every time
	for key, _ := range m.chain {
		return key
	}

	return "ab"
}

func getRandomLetter(slice []rune) rune {
	return slice[rand.Intn(len(slice))]
}
