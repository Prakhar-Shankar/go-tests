package iteration

import "strings"

const repeatCount = 5

func Iteration(character string) string{
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++{
		//repeated += character ---> repeated = repeated + chracter (Cocatenation)
		repeated.WriteString(character)
	}
	return repeated.String()
}