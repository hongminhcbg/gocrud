package utils

import (
	"log"
	"strings"

	"github.com/hongminhcbg/gocrud/models"
)

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

func ParseSqlHint(in string) *models.SqlHint {
	ans := new(models.SqlHint)
	args := strings.Split(in, ",")
	for _, kv := range args {
		keyVal := strings.Split(kv, "=")
		if len(keyVal) != 2 {
			log.Printf("invalid sql hint format %s, ignore", kv)
			continue
		}

		switch strings.ToLower(keyVal[0]) {
		case "type":
			ans.DataType = strings.TrimSpace(keyVal[1])
		case "default":
			ans.DefaultVal = strings.TrimSpace(keyVal[1])
		case "is_not_null":
			{
				if keyVal[1] == "true" {
					ans.IsNotNull = true
				}
			}
		case "key":
			{
				switch keyVal[1] {
				case "primary":
					ans.KeyType = models.SqlKeyType_PRIMARY
				case "candidate":
					ans.KeyType = models.SqlKeyType_CANDIDATE
				case "unique":
					ans.KeyType = models.SqlKeyType_UNIQUE
				default:
					log.Println("SQL hint key type must in list ['primary', 'candidate', 'unique', '']")
				}
			}
		}

	}
	return ans
}
