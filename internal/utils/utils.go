package utils

import "strings"

func Concat(a, b string) string {
	var sb strings.Builder
	sb.WriteString(a)
	sb.WriteString(b)
	return sb.String()
}
