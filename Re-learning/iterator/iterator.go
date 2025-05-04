package iterator

import "strings"

func loop (char string) string{
	var samp strings.Builder

	for i:=0; i<=5; i++ {
		samp.WriteString(char)
	}
	return samp.String();
}
