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
	words = applyCmds(words)
	words = mergePunct(words)
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

	var b strings.Builder
	needSpace := false

	for _, w := range words {
		if isSpaces(w) {
			if needSpace {
				b.WriteByte(' ')
			}
			b.WriteString(w)
			needSpace = false
			continue
		}
		if w == "\n" {
			b.WriteByte('\n')
			needSpace = false
			continue
		}
		if w == "\t" {
			b.WriteByte('\t')
			needSpace = false
			continue
		}

		if needSpace {
			b.WriteByte(' ')
		}
		b.WriteString(w)
		needSpace = true
	}

	file.WriteString(b.String())
}
func detectTextf(words []string) []string {
	for i := 0; i < len(words); i++ {
		tok := words[i]
		low := strings.ToLower(tok)
		if (low == "a" || low == "an") && i < len(words)-1 {
			checkAAn(words, i)
		}

		if isPunctTok(tok) {
			for i > 0 && isSpaces(words[i-1]) {
				words = removeIndex(words, i-1)
				i--
			}
			if i > 0 && words[i-1] != "\n" && words[i-1] != "\t" {
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
