package utils

import "strings"

func SnakeToCamel(snake string) string {
	var b strings.Builder
	shouldUpper := false
	diff := 'a' - 'A'
	for i := 0; i < len(snake); i++ {
		if i == 0 {
			b.WriteRune(rune(snake[i]) - diff)
			continue
		}

		if snake[i] == '_' {
			shouldUpper = true
			continue
		}

		if shouldUpper {
			shouldUpper = false
			b.WriteRune(rune(snake[i]) - diff)
			continue
		}

		b.WriteRune(rune(snake[i]))
	}

	return b.String()
}

func CamelToSnake(camel string) (snake string) {
	var b strings.Builder
	diff := 'a' - 'A'
	l := len(camel)
	for i, v := range camel {
		// A is 65, a is 97
		if v >= 'a' {
			b.WriteRune(v)
			continue
		}
		// v is capital letter here
		// irregard first letter
		// add underscore if last letter is capital letter
		// add underscore when previous letter is lowercase
		// add underscore when next letter is lowercase
		if (i != 0 || i == l-1) && (          // head and tail
		(i > 0 && rune(camel[i-1]) >= 'a') || // pre
			(i < l-1 && rune(camel[i+1]) >= 'a')) { //next
			b.WriteRune('_')
		}
		b.WriteRune(v + diff)
	}
	return b.String()
}
