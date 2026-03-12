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
	words = extractEmbeddedCommands(words)
	words = applyAllCommands(words)
	words = detectDQuotes(words)
	words = detectQuotes(words)
	words = detectTextf(words)
	words = mergeSpecialPunctuation(words)
	words = checkword(words)
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

func extractEmbeddedCommands(words []string) []string {
	var result []string
	for _, w := range words {
		// Ищем последнее вхождение '(' и проверяем, что после него есть ')'
		openIdx := strings.LastIndex(w, "(")
		closeIdx := strings.LastIndex(w, ")")
		if openIdx != -1 && closeIdx != -1 && closeIdx > openIdx {
			// Возможная команда внутри слова
			before := w[:openIdx]
			cmdPart := w[openIdx : closeIdx+1]
			after := w[closeIdx+1:]

			// Добавляем часть до команды (если не пустая)
			if before != "" {
				result = append(result, before)
			}
			// Добавляем команду как отдельный токен
			result = append(result, cmdPart)
			// Добавляем часть после команды (если не пустая)
			if after != "" {
				result = append(result, after)
			}
		} else {
			result = append(result, w)
		}
	}
	return result
}

//func detect(words []string) []string {
//	for i := 0; i < len(words); i++ {
//		cmd, ok := parseCommand(words[i])
//		if !ok {
//			continue
//		}
//
//		applyCommand(words, i, cmd)
//		words = removeIndex(words, i)
//		i--
//	}
//	return words
//}
//

func detectTextf(words []string) []string {
	for i := 0; i < len(words); i++ {
		word := words[i]

		switch word {
		case "a", "an":
			if i > 0 && i != len(words) {
				checkAAn(words, i)
			}
		case ",", ".", "!", "?", ":", ";", "...":
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
