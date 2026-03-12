package main

func razbiv(data []byte) []string {
	var slice []string
	var word []byte
	insideParen := false

	for _, f := range data {
		switch {
		case f == '(':
			insideParen = true
			word = append(word, f)
		case f == ')' && insideParen:
			word = append(word, f)
			slice = append(slice, string(word))
			word = nil
			insideParen = false
		case insideParen:
			word = append(word, f)
		case f == ' ':
			if len(word) > 0 {
				slice = append(slice, string(word))
				word = nil
			}
		default:
			word = append(word, f)
		}
	}
	if len(word) > 0 {
		slice = append(slice, string(word))
	}
	return slice
}
