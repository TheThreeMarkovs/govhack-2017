package main

import (
	"math/rand"
	"strings"
)

var businessEndWords = []string{
	"Limited",
	"Pty Ltd",
	"Systems",
	"Solutions",
	"Enterprises",
	"Holdings",
	"International",
	"Group",
}

type stringCounts struct {
	countMap   map[string]int
	totalCount int
}

func (counts *stringCounts) addString(str string) {
	if _, ok := counts.countMap[str]; ok {
		counts.countMap[str]++
	} else {
		counts.countMap[str] = 1
	}
	counts.totalCount++
}

func (counts *stringCounts) getRandomString() string {
	// Use count weightings to determine
	var probSum float64 = 0
	randProb := rand.Float64()

	var chosenStr string
	for chosenStr, count := range counts.countMap {
		thisProbRange := float64(count) / float64(counts.totalCount)
		if probSum < randProb && randProb < probSum+thisProbRange {
			return chosenStr
		}
		probSum += thisProbRange
	}

	return chosenStr
}

type Markov struct {
	chain            map[string]stringCounts
	order            int
	secondaryOrder   int
	strictWordStarts bool
	wordStarts       []string
}

func NewMarkov(order int, secondaryOrder int, strictWordStarts bool) *Markov {
	return &Markov{
		chain:            map[string]stringCounts{},
		order:            order,
		secondaryOrder:   secondaryOrder,
		strictWordStarts: strictWordStarts,
		wordStarts:       []string{},
	}
}

func (m *Markov) ParseWord(word string) {
	runes := []rune(word)
	if m.order > len(runes) {
		return
	}

	end := len(runes) - m.order - m.secondaryOrder
	for i := 0; i < end; i++ {
		key := string(runes[i : i+m.order])
		if i == 0 {
			m.wordStarts = append(m.wordStarts, key)
		}
		val := string(runes[i+m.order : i+m.order+m.secondaryOrder])
		runeCount, ok := m.chain[key]
		if ok {
			runeCount.addString(val)
		} else {
			m.chain[key] = stringCounts{
				countMap:   map[string]int{val: 1},
				totalCount: 1,
			}
		}
	}
}

func (m *Markov) GenerateWord(length int) string {
	letters := []rune(m.getRandomPrefix())

	for i := m.order; i <= length; i += m.secondaryOrder {
		lastLetters := letters[i-m.order : i]
		// fmt.Printf("Looking for key %v\n", string(lastLetters))
		if runes, ok := m.chain[string(lastLetters)]; ok {
			letters = append(letters, []rune(runes.getRandomString())...)
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

	if rand.Float32() > 0.5 {
		words = append(words, businessEndWords[rand.Intn(len(businessEndWords))])
	}

	return strings.Join(words, " ")
}

func (m *Markov) getRandomPrefix() string {
	if m.strictWordStarts {
		return m.wordStarts[rand.Intn(len(m.wordStarts))]
	} else {
		// Go randomises the order of maps every time
		for key, _ := range m.chain {
			return key
		}
	}

	return ""
}

func getRandomLetter(slice []rune) rune {
	return slice[rand.Intn(len(slice))]
}
