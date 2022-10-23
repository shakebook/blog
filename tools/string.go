package tools

import (
	"strings"
)

func PingString(s []string) string {
	var buffer strings.Builder
	for _, i := range s {
		buffer.WriteString(i)
	}
	return buffer.String()
}
