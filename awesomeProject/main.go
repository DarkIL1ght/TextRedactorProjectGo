package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("Read error:", err)
		os.Exit(1)
	}

	words := razbiv(data)
	words = splitPunct(words)
	words = applyAllCommands(words)
	words = mergeSpecialPunctuation(words)
	words = detectQuotes(words)
	words = detectDQuotes(words)
	words = detectTextf(words)
	writeText("result.txt", words)
}

func writeText(filename string, words []string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Create error:", err)
		return
	}
	defer file.Close()

	file.WriteString(strings.Join(words, " "))
}

func detectTextf(words []string) []string {
	for i := 0; i < len(words); i++ {
		word := words[i]
		word = strings.ToLower(word)
		if (word == "a" || word == "an") && i < len(words)-1 {
			checkAAn(words, i)
		}

		switch word {
		case ",", ".", "!", "?", ":", ";", "...", "?!", "!?":
			if i > 0 {
				words = combineDot(words, i)
				i--
			}
		}
	}
	return words
}

func removeIndex(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}
