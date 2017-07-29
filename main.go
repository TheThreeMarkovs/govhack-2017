package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	markov := NewMarkov()

	file, err := os.Open("companynames")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		words := getGoodWords(str)
		for _, w := range words {
			fmt.Println(w)
			markov.Parse(w)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("===============")

	for i := 0; i < 100; i++ {
		fmt.Println(markov.Generate())
	}

	// fmt.Printf("%+v\n", markov.states)
}

func getGoodWords(text string) []string {
	words := strings.Split(text, " ")
	var newWords []string
	for _, w := range words {
		if !isBadWord(w) {
			newWords = append(newWords, w)
		}
	}

	return newWords
}

func isBadWord(text string) bool {
	if strings.ContainsAny(text, "1234567890()'.-&") {
		return true
	}

	for _, w := range []string{
		"PTY",
		"LTD",
		"LIMITED",
		"HOLDINGS",
	} {
		if text == w {
			return true
		}
	}

	return false
}
