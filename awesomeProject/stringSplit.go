package main

func isSep(ch byte) bool {
	return ch == '\v' || ch == '\f'
}

func spacesToken(n int) string {
	if n <= 0 {
		return ""
	}
	b := make([]byte, n)
	for i := range b {
		b[i] = ' '
	}
	return string(b)
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
		if ch == ' ' {
			addWord(&words, &buf)

			start := i
			for i < len(data) && data[i] == ' ' {
				i++
			}
			n := i - start
			if n >= 2 {
				canAssumeBase := len(words) > 0 && !isWS(words[len(words)-1])
				if canAssumeBase {
					words = append(words, spacesToken(n-1))
				} else {
					words = append(words, spacesToken(n))
				}
			}
			continue
		}

		if ch == '\t' {
			addWord(&words, &buf)
			words = append(words, "\t")
			i++
			continue
		}
		if ch == '\r' {
			addWord(&words, &buf)
			if i+1 < len(data) && data[i+1] == '\n' {
				i += 2
			} else {
				i++
			}
			words = append(words, "\n")
			continue
		}
		if ch == '\n' {
			addWord(&words, &buf)
			words = append(words, "\n")
			i++
			continue
		}

		if isSep(ch) {
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
