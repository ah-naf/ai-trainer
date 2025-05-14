package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

type WordLists struct {
	Nouns      []string  `json:"Nouns"`
	Numbers    []float64 `json:"Numbers"`
	Adjectives []string  `json:"Adjectives"`
}

func getRandomIndex(length int) int {
	return rand.Intn(length)
}

func cleanWord(word string) string {
	return strings.TrimSuffix(word, ",")
}

func main() {

	jsonData, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var lists WordLists
	err = json.Unmarshal(jsonData, &lists)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	for i, word := range lists.Nouns {
		lists.Nouns[i] = cleanWord(word)
	}
	for i, word := range lists.Adjectives {
		lists.Adjectives[i] = cleanWord(word)
	}

	madlibText, err := os.ReadFile("madlib.txt")
	if err != nil {
		fmt.Println("Error reading madlib file:", err)
		return
	}

	madlib := string(madlibText)

	nounRegex := regexp.MustCompile(`\(noun\)`)
	adjectiveRegex := regexp.MustCompile(`\(adjective\)`)
	numberRegex := regexp.MustCompile(`\(number\)`)

	for nounRegex.MatchString(madlib) {
		randomNoun := lists.Nouns[getRandomIndex(len(lists.Nouns))]
		madlib = nounRegex.ReplaceAllLiteralString(madlib, randomNoun)
	}

	for adjectiveRegex.MatchString(madlib) {
		randomAdjective := lists.Adjectives[getRandomIndex(len(lists.Adjectives))]
		madlib = adjectiveRegex.ReplaceAllLiteralString(madlib, randomAdjective)
	}

	for numberRegex.MatchString(madlib) {
		randomNumber := lists.Numbers[getRandomIndex(len(lists.Numbers))]
		madlib = numberRegex.ReplaceAllLiteralString(madlib, fmt.Sprintf("%.1f", randomNumber))
	}

	fmt.Println("Completed Madlib:")
	fmt.Println(madlib)
}
