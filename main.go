package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.Open("business_names.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	markov := NewMarkov(4, 4, true)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		words := getGoodWords(str)
		for _, word := range words {
			markov.ParseWord(word)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 20; i++ {
		fmt.Println(markov.GenerateBusinessName())
	}

	// fmt.Printf("%+v\n", markov.chain)
	fmt.Printf("\nExecution Time: %s\n", time.Since(start))
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
	if strings.ContainsAny(text, "1234567890()'.-&/") {
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
