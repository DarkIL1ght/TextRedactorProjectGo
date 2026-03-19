package main

import (
	"strconv"
	"strings"
)

type Command struct {
	Name  string
	Count int
	HasN  bool
}

func isSpaces(tok string) bool {
	if tok == "" {
		return false
	}
	for i := 0; i < len(tok); i++ {
		if tok[i] != ' ' {
			return false
		}
	}
	return true
}

func isWS(tok string) bool {
	return isSpaces(tok) || tok == "\n" || tok == "\t"
}

func isPunctTok(tok string) bool {
	if tok == "" {
		return false
	}
	for _, r := range tok {
		if !isPunct(r) {
			return false
		}
	}
	return true
}

func skipForCmd(tok string) bool {
	return isWS(tok) || tok == "'" || tok == "\"" || isPunctTok(tok)
}

func isKnownCommand(cmd string) bool {
	switch cmd {
	case "cap", "low", "up", "hex", "bin":
		return true
	default:
		return false
	}
}

func parseCommand(word string) (Command, bool) {
	if !strings.HasPrefix(word, "(") || !strings.HasSuffix(word, ")") {
		return Command{}, false
	}

	content := word
	for strings.HasPrefix(content, "(") && strings.HasSuffix(content, ")") && len(content) >= 2 {
		content = content[1 : len(content)-1]
	}
	if strings.ContainsAny(content, "()") {
		return Command{}, false
	}
	content = strings.TrimSpace(content)

	parts := strings.Split(content, ",")
	cmd := strings.TrimSpace(parts[0])
	cmd = strings.ToLower(cmd)
	if !isKnownCommand(cmd) {
		return Command{}, false
	}
	if len(parts) == 1 {
		return Command{Name: cmd}, true
	}
	if len(parts) == 2 {
		n, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return Command{}, false
		}
		return Command{
			Name:  cmd,
			Count: n,
			HasN:  true,
		}, true
	}

	return Command{}, false

}

func clearity(word string) string {
	var newcommand []rune
	for _, r := range word {
		if r == '(' || r == ')' {
			continue
		} else {
			newcommand = append(newcommand, r)
		}
	}
	return string(newcommand)
}

func applyCmds(words []string) []string {
	for i := 0; i < len(words); i++ {
		cmd, ok := parseCommand(words[i])
		if !ok {
			continue
		}
		applyCommand(words, i, cmd)
		words = removeIndex(words, i)
		i--
	}
	return words
}

func applyCommand(words []string, index int, cmd Command) {
	count := 1
	if cmd.HasN {
		count = cmd.Count
	}

	pos := index - 1
	applied := 0
	for applied < count && pos >= 0 {
		for pos >= 0 && skipForCmd(words[pos]) {
			pos--
		}
		if pos < 0 {
			break
		}
		switch cmd.Name {
		case "cap":
			capWord(words, pos)
		case "low":
			words[pos] = strings.ToLower(words[pos])
		case "up":
			words[pos] = strings.ToUpper(words[pos])
		case "hex":
			hex(words, pos)
		case "bin":
			bin(words, pos)
		}
		applied++
		pos--
	}
}
