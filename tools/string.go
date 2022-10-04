package tools

import "bytes"

func PingString(s []string) string {
	var buffer bytes.Buffer
	for _, i := range s {
		buffer.WriteString(i)
	}
	return buffer.String()
}
