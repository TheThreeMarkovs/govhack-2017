package main

import (
	"bytes"
	"math/rand"
	"strings"
)

type Markov struct {
	states map[[2]string][]string
}

func NewMarkov() *Markov {
	markov := Markov{}
	markov.states = make(map[[2]string][]string)
	return &markov
}

func (m *Markov) Parse(text string) {
	letters := strings.Split(text, "")

	for i := 0; i < len(letters)-2; i++ {
		// Initialise prefix with 2 letters as the key
		prefix := [2]string{letters[i], letters[i+1]}

		// Assign the third letter as value to the prefix
		if _, ok := m.states[prefix]; ok {
			m.states[prefix] = append(m.states[prefix], letters[i+2])
		} else {
			m.states[prefix] = []string{letters[i+2]}
		}
	}
}

func (m *Markov) Generate() string {
	var phrase bytes.Buffer

	// Initialise prefix with a random key
	prefix := m.getRandomPrefix([2]string{"", ""})
	phrase.WriteString(strings.Join(prefix[:], ""))
	limit := rand.Intn(15) + 5

	for i := 0; i < limit; i++ {
		suffix := getRandomLetter(m.states[prefix])
		phrase.WriteString(suffix)

		prefix = [2]string{prefix[1], suffix}
	}

	return phrase.String()
}

func (m *Markov) getRandomPrefix(prefix [2]string) [2]string {
	// By default, go orders randomly for maps
	for key := range m.states {
		if key != prefix {
			prefix = key
			break
		}
	}

	return prefix
}

func getRandomLetter(slice []string) string {
	if !(cap(slice) == 0) {
		return slice[rand.Intn(len(slice))]
	}
	return ""
}
