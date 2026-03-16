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

	// Allow extra outer parentheses like "((up))" but reject inner parentheses like "(((u)p))".
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

func applyAllCommands(words []string) []string {
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

	for j := 1; j <= count && index-j >= 0; j++ {
		switch cmd.Name {
		case "cap":
			capWord(words, index-j)
		case "low":
			words[index-j] = strings.ToLower(words[index-j])
		case "up":
			words[index-j] = strings.ToUpper(words[index-j])
		case "hex":
			hex(words, index-j)
		case "bin":
			bin(words, index-j)
		}
	}
}
