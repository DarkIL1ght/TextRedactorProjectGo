package main

func razbiv(data []byte) []string {
	var slice []string
	var word []byte
	level := 0

	for _, ch := range data {
		switch {
		case ch == '(':
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
