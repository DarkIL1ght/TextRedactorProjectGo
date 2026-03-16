package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func hex(words []string, idx int) {
	num, err := strconv.ParseInt(words[idx], 16, 64)
	if err == nil {
		words[idx] = fmt.Sprintf("%d", num)
	}
}

func bin(words []string, idx int) {
	num, err := strconv.ParseInt(words[idx], 2, 64)
	if err == nil {
		words[idx] = fmt.Sprintf("%d", num)
	}
}

func capWord(words []string, idx int) {
	if len(words[idx]) == 0 {
		return
	}
	runes := []rune(words[idx])
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	words[idx] = string(runes)
}

func checkAAn(words []string, index int) {
	if index+1 >= len(words) {
		return
	}
	if len(words[index+1]) == 0 {
		return
	}
	runes := []rune(words[index+1])
	run := unicode.ToLower(runes[0])
	wantAn := isVowel(run)

	switch words[index] {
	case "a":
		if wantAn {
			words[index] = "an"
		}
	case "A":
		if wantAn {
			words[index] = "An"
		}
	case "an":
		if !wantAn {
			words[index] = "a"
		}
	case "An":
		if !wantAn {
			words[index] = "A"
		}
	}

}
func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'h':
		return true
	}
	return false
}

func combineDot(words []string, index int) []string {
	if index <= 0 || index >= len(words) {
		return words
	}
	if len(words[index]) == 0 {
		return words
	}
	words[index-1] += words[index]
	return append(words[:index], words[index+1:]...)
}

func mergeSpecialPunctuation(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		// Проверяем пары "?" и "!"
		if (words[i] == "?" && words[i+1] == "!") || (words[i] == "!" && words[i+1] == "?") {
			words[i] = words[i] + words[i+1]
			words = removeIndex(words, i+1)
			i-- // после удаления следующий элемент смещается
			continue
		}
		// Проверяем три точки подряд
		if i+2 < len(words) && words[i] == "." && words[i+1] == "." && words[i+2] == "." {
			words[i] = "..."
			words = removeIndex(words, i+2) // удаляем третью точку
			words = removeIndex(words, i+1) // удаляем вторую точку
			i--                             // после двух удалений остаёмся на том же i (сместилось)
		}
	}
	return words
}

func detectQuotes(words []string) []string {
	result := make([]string, 0, len(words))
	i := 0
	for i < len(words) {
		if words[i] == "'" {
			j := i + 1
			for j < len(words) && words[j] != "'" {
				j++
			}
			if j == len(words) {
				result = append(result, words[i])
				i++
				continue
			}
			if j == i+1 {
				i = j + 1
				continue
			}
			words[i+1] = "'" + words[i+1]
			words[j-1] = words[j-1] + "'"
			for k := i + 1; k < j; k++ {
				result = append(result, words[k])
			}
			i = j + 1
		} else {
			result = append(result, words[i])
			i++
		}
	}
	return result
}

func detectDQuotes(words []string) []string {
	result := make([]string, 0, len(words))
	i := 0
	for i < len(words) {
		if words[i] == "\"" {
			j := i + 1
			for j < len(words) && words[j] != "\"" {
				j++
			}
			if j == len(words) {
				result = append(result, words[i])
				i++
				continue
			}
			if j == i+1 {
				i = j + 1
				continue
			}
			words[i+1] = "\"" + words[i+1]
			words[j-1] = words[j-1] + "\""
			for k := i + 1; k < j; k++ {
				result = append(result, words[k])
			}
			i = j + 1
		} else {
			result = append(result, words[i])
			i++
		}
	}
	return result
}

func splitPunct(words []string) []string {
	var res []string
	for _, w := range words {
		if (strings.HasPrefix(w, "(") && strings.HasSuffix(w, ")")) || w == "'" || w == "\"" {
			res = append(res, w)
			continue
		}

		runes := []rune(w)
		i := 0
		for i < len(runes) {
			if isPunct(runes[i]) {
				j := i
				for j < len(runes) && isPunct(runes[j]) {
					j++
				}
				res = append(res, string(runes[i:j]))
				i = j
			} else {
				j := i
				for j < len(runes) && !isPunct(runes[j]) {
					j++
				}
				res = append(res, string(runes[i:j]))
				i = j
			}
		}
	}
	return res
}

func isPunct(r rune) bool {
	return r == '.' || r == ',' || r == '!' || r == '?' || r == ':' || r == ';'
}
