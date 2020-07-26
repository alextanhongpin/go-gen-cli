package gen

import (
	"regexp"
	"strings"
)

var camelCaseRe, pascalCaseRe *regexp.Regexp

func init() {
	camelCaseRe = regexp.MustCompile(`(?i)[^a-z]+[a-z]`)
	pascalCaseRe = regexp.MustCompile(`(?i)(^[a-z]|[^a-z]+[a-z])`)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func CamelCase(str string) string {
	i := min(1, len(str))
	str = strings.ToLower(str[:i]) + str[i:]
	return camelCaseRe.ReplaceAllStringFunc(str, func(s string) string {
		return strings.ToUpper(s[len(s)-1:])
	})
}

func PascalCase(str string) string {
	return pascalCaseRe.ReplaceAllStringFunc(str, func(s string) string {
		return strings.ToUpper(s[len(s)-1:])
	})
}
