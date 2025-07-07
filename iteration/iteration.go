package iteration

import "strings"


func Iteration(character string, count int) string{
	var repeated strings.Builder
	for i := 0; i < count; i++{
		//repeated += character ---> repeated = repeated + chracter (Cocatenation)
		repeated.WriteString(character)
	}
	return repeated.String()
}