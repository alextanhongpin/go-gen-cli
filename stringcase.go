package gen

import (
	"regexp"
	"strings"
)

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(`(?i)(^[a-z]|[^a-z0-9]+[a-z0-9])`)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func CamelCase(str string) string {
	str = PascalCase(str)
	i := min(1, len(str))
	str = strings.ToLower(str[:i]) + str[i:]
	return str
}

func PascalCase(str string) string {
	return re.ReplaceAllStringFunc(str, func(s string) string {
		return strings.ToUpper(s[len(s)-1:])
	})
}
