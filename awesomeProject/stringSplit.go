package main

func isSpace(ch byte) bool {
	return ch == ' ' || ch == '\n' || ch == '\r' || ch == '\t' || ch == '\v' || ch == '\f'
}

func addWord(words *[]string, buf *[]byte) {
	if len(*buf) == 0 {
		return
	}
	*words = append(*words, string(*buf))
	*buf = nil
}

func razbiv(data []byte) []string {
	var words []string
	var buf []byte

	i := 0
	for i < len(data) {
		ch := data[i]

		if isSpace(ch) {
			addWord(&words, &buf)
			i++
			continue
		}

		if ch == '(' {
			j := i + 1
			for j < len(data) && data[j] != ')' {
				j++
			}
			if j < len(data) && data[j] == ')' {
				s := string(data[i : j+1])
				if _, ok := parseCommand(s); ok {
					addWord(&words, &buf)
					words = append(words, s)
					i = j + 1
					continue
				}
			}
		}

		buf = append(buf, ch)
		i++
	}

	addWord(&words, &buf)
	return words
}
