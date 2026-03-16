package main

func razbiv(data []byte) []string {
	var slice []string
	var word []byte
	level := 0

	for _, ch := range data {
		switch {
		case ch == '(':
			// If a command is attached to the end of a word (e.g. "23(up)"),
			// flush the word first so the command becomes its own token "(up)".
			if level == 0 && len(word) > 0 {
				slice = append(slice, string(word))
				word = nil
			}
			level++
			word = append(word, ch)
		case ch == ')':
			word = append(word, ch)
			level--
			if level == 0 {
				slice = append(slice, string(word))
				word = nil
			}
		case level > 0:
			word = append(word, ch)
		case ch == ' ':
			if len(word) > 0 {
				slice = append(slice, string(word))
				word = nil
			}
		default:
			word = append(word, ch)
		}
	}
	if len(word) > 0 {
		slice = append(slice, string(word))
	}
	return slice
}
